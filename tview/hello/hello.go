package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	left := tview.NewList().
		AddItem("Hello 1", "Hello", 'a', nil).
		AddItem("Hello 2", "Hello", 'b', nil).
		AddItem("Hello 3", "Hello", 'c', nil).
		AddItem("Hello 4", "Hello", 'd', nil).
		AddItem("Bye", "Bye", 'q', func() {
			app.Stop()
		})
	left.SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[::b] Hello Left ").
		SetTitleColor(tcell.ColorBlue)
	right_top := tview.NewTextView().
		SetText("Hello Hello Hello Hello Hello\n" +
			"Hello Hello Hello Hello Hello\n" +
			"Hello Hello Hello Hello Hello\n" +
			"Hello Hello Hello Hello Hello\n" +
			"Hello Hello Hello Hello Hello\n")
	right_top.SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[::b] Hello Right Top ").
		SetTitleColor(tcell.ColorBlue)
	right_bottom := tview.NewTextView().
		SetChangedFunc(func() {
			app.Draw()
		})
	go func() {
		for {
			for _, letter := range []string{"H", "e", "l", "l", "o"} {
				fmt.Fprintf(right_bottom, "%s", letter)
				time.Sleep(time.Second)
			}
			right_bottom.SetText("")
		}
	}()
	right_bottom.SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[::b] Hello Right Bottom ").
		SetTitleColor(tcell.ColorBlue)
	flex := tview.NewFlex().
		AddItem(left, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(right_top, 0, 3, false).
			AddItem(right_bottom, 0, 1, false), 0, 2, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		log.Fatal(err)
	}
}
