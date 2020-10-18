package join

import "strings"

func Join(sep string, elms ...string) string {
	return strings.Join(elms, sep)
}
