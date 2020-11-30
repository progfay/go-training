package memo_test

import (
	"testing"
	"time"

	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/progfay/go-training/ch09/ex03/memo"
)

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func TestMemo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
	}))
	defer ts.Close()

	m := memo.New(httpGetBody)
	defer m.Close()

	done := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()
	value, err := m.Get(ts.URL, done)
	if value != nil {
		t.Errorf("value: want %v, got %v", nil, value)
	}

	if err == nil {
		t.Error("err should be occurred when func is cancelled")
	}
}
