package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/progfay/go-training/ch07/ex15/eval"
)

func main() {
	r := bufio.NewScanner(os.Stdin)

	fmt.Fprint(os.Stdout, "expr: ")
	if !r.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan expr")
		os.Exit(1)
	}
	input := r.Text()
	e, err := eval.Parse(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	vars := eval.GetVars(e)
	env := eval.Env{}
	for _, v := range vars {
		for {
			fmt.Fprintf(os.Stdout, "%s: ", v)
			if !r.Scan() {
				fmt.Fprintf(os.Stderr, "failed to scan value of var %q\n", v)
				os.Exit(1)
			}
			input = r.Text()
			value, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "float 64 value %q: %v\n", input, err)
				continue
			}
			env[eval.Var(v)] = value
			break
		}
	}

	fmt.Fprintf(os.Stdout, "result: %f\n", e.Eval(env))
}
