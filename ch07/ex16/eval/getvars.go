package eval

import "fmt"

func GetVars(e Expr) []string {
	varMap := make(map[string]bool, 0)

	var f func(Expr)
	f = func(e Expr) {
		switch e := e.(type) {
		case literal:

		case Var:
			varMap[string(e)] = true

		case unary:
			f(e.x)

		case binary:
			f(e.x)
			f(e.y)

		case call:
			for _, x := range e.args {
				f(x)
			}

		case abs:
			f(e.x)

		default:
			panic(fmt.Sprintf("unknown Expr: %T", e))
		}
	}

	f(e)

	vars := []string{}
	for v := range varMap {
		vars = append(vars, v)
	}

	return vars
}
