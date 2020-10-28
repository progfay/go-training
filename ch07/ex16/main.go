package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/progfay/go-training/ch07/ex15/eval"
)

var htmlTemplate *template.Template

type parsedResult struct {
	Expr eval.Expr
	Err  error
}

var parseMemo map[string]parsedResult
var varsMemo map[string][]string

func init() {
	parseMemo = make(map[string]parsedResult)
	varsMemo = make(map[string][]string)

	htmlTemplate = template.Must(template.New("escape").Parse(`
	<form action="/eval">
		<label for="expr">expr:</label>
		<input id="expr" name="expr">
		<button type="submit">eval</button>
		<hr>
		<div id="vars"></div>
	</form>
	<script>
		window.onload = () => {
			const exprInput = document.getElementById('expr')
			const varsDiv = document.getElementById('vars')

			const debounce = (func, wait = 0) => {
				let timeout = null
				return (...args) => {
					if (!timeout) clearTimeout(timeout)
					timeout = setTimeout(() => { func(args) }, wait)
				}
			}

			const inputHandler = debounce(() => {
				fetch('/getVars?expr=' + encodeURIComponent(exprInput.value))
					.then(res => res.json())
					.then(vars => {
						varsDiv.innerHTML = ''
						for (const v of vars) {
							const label = document.createElement('label')
							label.for = v
							label.appendChild(document.createTextNode(v + ':'))
							varsDiv.appendChild(label)

							const input = document.createElement('input')
							input.id = v
							input.name = v
							varsDiv.appendChild(input)
						}
					})
			}, 500)

			exprInput.addEventListener('input', inputHandler)
		}
	</script>
	`))
}

func parse(expr string) (eval.Expr, error) {
	if result, ok := parseMemo[expr]; ok {
		return result.Expr, result.Err
	}

	e, err := eval.Parse(expr)
	parseMemo[expr] = parsedResult{
		Expr: e,
		Err:  err,
	}
	return e, err
}

func getVars(expr string) []string {
	if vars, ok := varsMemo[expr]; ok {
		return vars
	}

	e, err := parse(expr)
	if err != nil {
		varsMemo[expr] = []string{}
		return varsMemo[expr]
	}

	varsMemo[expr] = eval.GetVars(e)
	return varsMemo[expr]
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/getVars", getVarsHandler)
	http.HandleFunc("/eval", evalHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := htmlTemplate.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error on generating html")
	}
}

func getVarsHandler(w http.ResponseWriter, req *http.Request) {
	expr := req.URL.Query().Get("expr")

	vars := getVars(expr)
	bytes, err := json.Marshal(vars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(bytes))
}

func evalHandler(w http.ResponseWriter, req *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}
	}()

	expr := req.URL.Query().Get("expr")

	e, err := parse(expr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	vars := getVars(expr)

	env := eval.Env{}
	for _, v := range vars {
		valueStr := req.URL.Query().Get(v)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid float64 value %q of %q: %v\n", valueStr, v, err)
			return
		}
		env[eval.Var(v)] = value
	}

	result := e.Eval(env)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, result)
}
