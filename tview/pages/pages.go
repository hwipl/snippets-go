package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// create application
	app := tview.NewApplication()

	// create boxes
	box1 := tview.NewBox().SetBorder(true).SetTitle("Box 1")
	box2 := tview.NewBox().SetBorder(true).SetTitle("Box 2")
	box3 := tview.NewBox().SetBorder(true).SetTitle("Box 3")

	// create pages and put boxes in it
	pages := tview.NewPages().
		AddPage("box 3", box3, true, true).
		AddPage("box 2", box2, true, true).
		AddPage("box 1", box1, true, true)

	// handle user input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// close program when user presses "q" or <Esc>
		if event.Rune() == 'q' || event.Key() == tcell.KeyCtrlLeftSq {
			app.Stop()
			return nil
		}

		// for all other keys, simply rotate boxes
		name, _ := pages.GetFrontPage()
		pages.SendToBack(name)

		return event
	})

	// run everything
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		log.Fatal(err)
	}
}
