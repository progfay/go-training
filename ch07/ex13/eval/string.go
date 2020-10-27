package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%v)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%v %c %v)", b.x, b.op, b.y)
}

func (c call) String() string {
	args := make([]string, len(c.args))
	for i, arg := range c.args {
		args[i] = arg.String()
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}
