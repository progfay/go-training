package params

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Pack(s interface{}) url.URL {
	u := url.URL{}
	q := u.Query()

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		q.Set(name, stringify(v.Field(i)))
	}

	u.RawQuery = q.Encode()
	return u
}

func stringify(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()

	case reflect.Int:
		return fmt.Sprintf("%d", v.Int())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Float())

	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())

	default:
		return v.String()
	}
}
