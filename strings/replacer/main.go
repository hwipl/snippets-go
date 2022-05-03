package main

import (
	"fmt"
	"strings"
)

func main() {
	const text = `
this is an example text. It conains $VAR and
{test} and also {{other}}.
`
	r := strings.NewReplacer("$VAR", "this")
	fmt.Println(r.Replace(text))

	r = strings.NewReplacer("$VAR", "this", "{test}", "that",
		"{{other}}", "other")
	fmt.Println(r.Replace(text))
}
