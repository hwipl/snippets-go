package main

import (
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

// onExit is executed when systray is exited
func onExit() {
	log.Println("Goodbye.")
}

// onReady is executed when systray is ready
func onReady() {
	log.Println("Hello.")

	// set icon, title and tooltip
	systray.SetIcon(icon.Data)
	systray.SetTitle("Hello")
	systray.SetTooltip("Hello, there.")

	// add hello entry
	mHi := systray.AddMenuItem("Hi", "Hello.")
	go func() {
		for {
			<-mHi.ClickedCh
			log.Println("Hello.")
		}
	}()

	// add separator
	systray.AddSeparator()

	// add goodbye entry
	mBye := systray.AddMenuItem("Bye", "Goodbye.")
	go func() {
		<-mBye.ClickedCh
		systray.Quit()
	}()
}

func main() {
	// run systray until quit
	systray.Run(onReady, onExit)
}
