package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(200, 200))

	greetings := []string{
		"Hi",
		"Hello",
		"Good Day",
		"Bye",
		"Goodbye",
		"Farewell",
	}
	msg := widget.NewLabel("")
	sel := widget.NewSelect(greetings, nil)
	sel.Selected = greetings[0]
	but := widget.NewButton("greet", func() {
		msg.SetText(sel.Selected + "!")
	})
	quit := widget.NewButton("quit", func() {
		a.Quit()
	})

	w.SetContent(container.NewVBox(
		container.NewHBox(sel, but, quit),
		msg,
	))

	w.ShowAndRun()
}
