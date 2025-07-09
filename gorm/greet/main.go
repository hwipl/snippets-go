package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// dbFile is the db file.
const dbFile = "./greetings.db"

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
	_ = os.Remove(dbFile)

	// open db
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	// migrate schema
	_ = db.AutoMigrate(&Greeting{})

	// create
	db.Create(&Greeting{Greeting: "hello"})

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
