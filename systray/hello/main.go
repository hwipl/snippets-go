package main

import (
	"log"
	"time"

	"github.com/getlantern/systray"
)

// onExit is executed when systray is exited
func onExit() {
	log.Println("Goodbye.")
}

// onReady is executed when systray is ready
func onReady() {
	log.Println("Hello.")

	// set icon, title and tooltip
	systray.SetIcon(icon1)
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

	// add periodic icon changing
	go func() {
		visible1 := true
		for {
			<-time.NewTimer(time.Second).C
			if visible1 {
				systray.SetIcon(icon2)
				visible1 = false
			} else {
				systray.SetIcon(icon1)
				visible1 = true
			}
		}
	}()
}

func main() {
	// run systray until quit
	systray.Run(onReady, onExit)
}
