package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// opContains checks if op contains the specific file operation typ, i.e,
// fsnotify.Create, fsnotify.Write, fsnotify.Remove, fsnotify.Rename, or
// fsnotify.Chmod
func opContains(op, typ fsnotify.Op) bool {
	if op&typ == typ {
		return true
	}
	return false
}

// handleEvent handles an fsnotify event
func handleEvent(event fsnotify.Event) {
	if opContains(event.Op, fsnotify.Create) {
		log.Println("File created:", event.Name)
	}
	if opContains(event.Op, fsnotify.Write) {
		log.Println("File written:", event.Name)
	}
	if opContains(event.Op, fsnotify.Remove) {
		log.Println("File removed:", event.Name)
	}
	if opContains(event.Op, fsnotify.Rename) {
		log.Println("File renamed:", event.Name)
	}
	if opContains(event.Op, fsnotify.Chmod) {
		log.Println("File chmod:", event.Name)
	}
}

func main() {
	// create watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// watch tmp dir
	watcher.Add("/tmp")

	// handle file events and errors
	for {
		select {
		case event, more := <-watcher.Events:
			if !more {
				return
			}
			handleEvent(event)
		case err, more := <-watcher.Errors:
			if !more {
				return
			}
			log.Println(err)
		}
	}
}
