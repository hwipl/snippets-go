package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	left := tview.NewBox().
		SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[::b] Hello Left ").
		SetTitleColor(tcell.ColorBlue)
	right_top := tview.NewBox().
		SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[::b] Hello Right Top ").
		SetTitleColor(tcell.ColorBlue)
	right_bottom := tview.NewBox().
		SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[::b] Hello Right Bottom ").
		SetTitleColor(tcell.ColorBlue)
	flex := tview.NewFlex().
		AddItem(left, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(right_top, 0, 3, false).
			AddItem(right_bottom, 0, 1, false), 0, 2, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		log.Fatal(err)
	}
}
