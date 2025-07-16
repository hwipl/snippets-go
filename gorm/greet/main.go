package main

import (
	"flag"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Greeting is a greeting
type Greeting struct {
	gorm.Model
	Greeting string
}

// list prints all greetings.
func list(db *gorm.DB) {
	var greetings []Greeting
	db.Find(&greetings)
	fmt.Println("list:")
	for _, g := range greetings {
		fmt.Println("  ", g.ID, g.Greeting)
	}
}

// getID prints greeting with ID.
func getID(db *gorm.DB, id int) {
	var g Greeting
	db.First(&g, id)
	fmt.Println("getID:", id, g.Greeting)
}

// getGreeting prints greeting.
func getGreeting(db *gorm.DB, greeting string) {
	var g Greeting
	db.First(&g, "Greeting = ?", greeting)
	fmt.Println("getGreeting:", greeting, g.ID)
}

// updateGreeting updates the greeting with ID.
func updateGreeting(db *gorm.DB, id int, greeting string) {
	var g Greeting
	db.First(&g, id)
	db.Model(&g).Update("Greeting", greeting)
}

// deleteID deletes greeting with ID.
func deleteID(db *gorm.DB, id int) {
	var g Greeting
	db.Delete(&g, id)
}

func main() {
	// open db
	db, err := gorm.Open(sqlite.Open("file:greet.db?mode=memory"), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	// migrate schema
	_ = db.AutoMigrate(&Greeting{})

	// create entries
	for _, g := range []string{"hello", "hi", "good day", "greetings"} {
		db.Create(&Greeting{Greeting: g})
	}

	// command line arguments
	r := flag.Bool("run", false, "run some commands for testing")
	l := flag.Bool("list", false, "list greetings")
	i := flag.Int("id", 0, "get greeting by `id`")
	g := flag.String("greeting", "", "get greeting")
	flag.Parse()

	if *r {
		// list
		list(db)

		// read
		getID(db, 1)
		getGreeting(db, "hello")

		// update
		updateGreeting(db, 1, "hi")
		getID(db, 1)

		// delete
		deleteID(db, 1)
		list(db)
	}

	if *l {
		list(db)
		return
	}
	if *i != 0 {
		getID(db, *i)
		return
	}
	if *g != "" {
		getGreeting(db, *g)
		return
	}
}
