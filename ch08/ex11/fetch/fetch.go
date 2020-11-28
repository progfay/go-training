package fetch

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func fetch(url string, cancel <-chan struct{}) (*http.Response, error) {
	var buf bytes.Buffer
	req, err := http.NewRequest("GET", url, &buf)
	if err != nil {
		return nil, err
	}

	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	return resp, nil
}

type resultType struct {
	resp *http.Response
	err  error
}

// Race request to all url of first arguments, and return fastest response
// if fastest response occur error (e.g.: response status is not http.OK), return nil, error
func Race(urls []string) (*http.Response, error) {
	var wg sync.WaitGroup
	cancel := make(chan struct{})
	result := make(chan resultType)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			c := make(chan struct{})
			resp, err := fetch(url, c)
			select {
			case <-cancel:
				close(c)
				return

			default:
				result <- resultType{
					resp: resp,
					err:  err,
				}
				close(cancel)
			}
		}(url)
	}

	r := <-result
	wg.Wait()
	return r.resp, r.err
}

// AggregateError represents wrap multi errors to single error
type AggregateError struct {
	errs []error
}

func (e *AggregateError) Error() string {
	var buf bytes.Buffer
	for _, err := range e.errs {
		buf.WriteString(err.Error())
		buf.WriteRune('\n')
	}

	return buf.String()
}

// Any request to all url of first arguments, and return fastest success response
// if all request is failed, return nil, error
func Any(urls []string) (*http.Response, error) {
	var wg sync.WaitGroup
	cancel := make(chan struct{})
	result := make(chan resultType)
	errsChan := make(chan error, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			c := make(chan struct{})
			resp, err := fetch(url, c)
			select {
			case <-cancel:
				close(c)
				return

			default:
				if err != nil {
					errsChan <- fmt.Errorf("%q: %w", url, err)
					return
				}
			}
			result <- resultType{
				resp: resp,
				err:  err,
			}
			close(cancel)
		}(url)
	}

	go func() {
		wg.Wait()
		errs := make([]error, 0)

	loop:
		for {
			select {
			case err := <-errsChan:
				errs = append(errs, err)

			default:
				close(errsChan)
				break loop
			}
		}

		result <- resultType{
			resp: nil,
			err: &AggregateError{
				errs: errs,
			},
		}
		close(cancel)
	}()

	r := <-result
	wg.Wait()
	return r.resp, r.err
}
