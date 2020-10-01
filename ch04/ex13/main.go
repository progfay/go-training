package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/progfay/go-training/ch04/ex13/github"
)

const templ = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>GitHub Issue Viewer</title>
		<style>
			.issue {
				margin: 10px 0;
				padding: 5px;
				width: 100vw;
				border: solid 1px black;
			}
		</style>
</head>

<body>
    <form action="/" method="GET">
        <label for="query">Search Term</label> 
        <input type="text" id="query" name="query">
        <input type="submit">
    </form>

		{{if ne .Query ""}}
			<hr>

			<article>
					<h2>Search Result of <code>{.Query}</code>: found {{.TotalCount}}</h2>

					{{range .Items}}
							<section class="issue">
									<h3><a href="{{.HTMLURL}}">#{{.Number}} {{.Title}}</a></h3>
									<p>Created by <a href="{{.User.HTMLURL}}">@{{.User.Login}}</a></p>
									<h4>Description</h4>
									<pre>
										<code>{{.Body | trimSpace}}</code>
									</pre>
									{{if .Milestone}}
										<h4>Milestone</h4>
										<details>
												<summary>
														<a href="{{.Milestone.HTMLURL}}">#{{.Milestone.Number}} {{.Milestone.Title}}</a>
												</summary>
												<p>
														{{.Milestone.Description | trimSpace}}
												</p>
										</details>
									{{end}}
							</section>
					{{end}}
			</article>
		{{end}}

    <script>
        window.onload = () => {
            document.getElementById('query').value = new URLSearchParams(location.search).get('query')
        }
    </script>
</body>

</html>
`

var issueTemplate *template.Template

func init() {
	issueTemplate = template.Must(template.New("escape").Funcs(template.FuncMap{"trimSpace": strings.TrimSpace}).Parse(templ))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	query := ""
	if len(r.Form["query"]) > 0 && r.Form["query"][0] != "" {
		query = r.Form["query"][0]
	}

	result, err := github.SearchIssues(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := issueTemplate.Execute(w, result); err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listen on http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
