package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const md = `
# {{.Title}}
{{with .Abstract}}
## Abstract

{{.}}
{{end}}
{{with .Topics}}## Topics
{{range .}}
- {{.}}{{end}}{{end}}
`

	type doc struct {
		Title    string
		Abstract string
		Topics   []string
	}

	docs := []doc{
		doc{
			Title:    "test document",
			Abstract: "this is a test document",
			Topics:   []string{"test", "document"},
		},
		doc{
			Title:  "Another test document",
			Topics: []string{"another", "test", "document"},
		},
		doc{
			Title:    "Yet another test document",
			Abstract: "this is yet another test document",
		},
	}
	t := template.Must(template.New("md").Parse(md))
	for _, d := range docs {
		err := t.Execute(os.Stdout, d)
		if err != nil {
			log.Println(err)
		}
	}
}
