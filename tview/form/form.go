package main

import (
	"log"
	"strings"

	"github.com/rivo/tview"
)

func main() {
	// create application
	app := tview.NewApplication()

	// create form values
	greeting := ""
	name := ""
	yell := false

	// create form
	form := tview.NewForm().
		AddDropDown("Greeting", []string{"Hi", "Hey", "Bye"}, 0,
			func(option string, optionIndex int) {
				greeting = option
			}).
		AddInputField("Name", "", 20, nil,
			func(text string) {
				name = text
			}).
		AddCheckbox("Yell", false,
			func(checked bool) {
				yell = checked
			}).
		AddButton("Greet", func() {
			app.Stop()
		}).
		AddButton("Cancel", func() {
			greeting = ""
			name = ""
			yell = false
			app.Stop()
		})

	// run form
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		log.Fatal(err)
	}

	// no name, no greeting :(
	if name == "" {
		return
	}

	if yell {
		// yell greeting
		log.Printf("%s, %s!",
			strings.ToUpper(greeting),
			strings.ToUpper(name))
		return
	}

	// regular greeting
	log.Printf("%s, %s.", greeting, name)
}
