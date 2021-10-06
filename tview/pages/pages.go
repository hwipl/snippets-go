package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const numPages = 16

func main() {
	// create application
	app := tview.NewApplication()

	// create boxes
	boxes := []*tview.Box{}
	for i := 0; i < numPages; i++ {
		box := tview.NewBox().
			SetBorder(true).
			SetTitle(fmt.Sprintf(" Box %d ", i+1))
		boxes = append(boxes, box)
	}

	// create pages and put boxes in it
	pages := tview.NewPages()
	for i := numPages - 1; i >= 0; i-- {
		pages.AddPage(fmt.Sprintf("box %d", i+1), boxes[i], true, true)
	}

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
