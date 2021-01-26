package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "true")
		} else {
			fmt.Fprint(buf, "false")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, `"#C(%f %f)"`, real(c), imag(c))

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(",")
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct:
		buf.WriteByte('{')
		isFirst := true
		for i := 0; i < v.NumField(); i++ {
			zero := reflect.Zero(v.Field(i).Type()).Interface()
			if reflect.DeepEqual(v.Field(i).Interface(), zero) {
				continue
			}
			if !isFirst {
				buf.WriteByte(',')
			} else {
				isFirst = false
			}
			fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map:
		buf.WriteByte('{')
		isFirst := true
		for _, key := range v.MapKeys() {
			zero := reflect.Zero(v.MapIndex(key).Type())
			if reflect.DeepEqual(v.MapIndex(key).Interface(), zero) {
				continue
			}
			if !isFirst {
				buf.WriteByte(',')
			} else {
				isFirst = false
			}
			if key.Kind() != reflect.String {
				return fmt.Errorf("key of map must be string")
			}
			fmt.Fprintf(buf, "%q:", key.String())
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
