package pattern

import (
	"fmt"
	"reflect"
	"regexp"
)

var (
	validator = map[string]*regexp.Regexp{
		"post-code": regexp.MustCompile("^[0-9]{3}-[0-9]{4}$"),
		"email":     regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"),
	}
)

func Validate(target interface{}) error {
	value := reflect.ValueOf(target)
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Kind() != reflect.String {
			continue
		}

		fieldInfo := value.Type().Field(i)
		tag := fieldInfo.Tag
		pattern := tag.Get("pattern")
		if pattern == "" {
			continue
		}
		r, ok := validator[pattern]
		if !ok {
			return fmt.Errorf("unknown pattern name: %q", pattern)
		}

		s := value.Field(i).String()
		if !r.MatchString(s) {
			return fmt.Errorf("%q: %q is not %q format", fieldInfo.Name, s, pattern)
		}
	}

	return nil
}

func AddValidator(name string, r *regexp.Regexp) {
	validator[name] = r
}
