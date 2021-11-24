package main

import "github.com/gen2brain/beeep"

func main() {
	beeep.Notify("Hi", "Hi, there!", "")
	beeep.Alert("Hey!", "Hey, there!", "")
	beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	beeep.Notify("Bye", "Ok, bye :(", "")
}
