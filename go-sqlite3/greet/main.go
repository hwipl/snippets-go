package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// addGreetings adds greetings to db with IDs starting from 0.
func addGreetings(db *sql.DB, greetings []string) {
	// fill db
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into greetings(id, greeting) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = stmt.Close() }()

	for i, g := range greetings {
		_, err = stmt.Exec(i+1, g)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

// addGreeting adds a greeting with an ID to db.
func addGreeting(db *sql.DB, id int, greeting string) {
	_, err := db.Exec(fmt.Sprintf(
		"insert into greetings(id, greeting) values(%d, '%s')",
		id, greeting))
	if err != nil {
		log.Fatal(err)
	}
}

// list lists all entries in db.
func list(db *sql.DB) {
	rows, err := db.Query("select id, greeting from greetings")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var id int
		var greeting string
		err = rows.Scan(&id, &greeting)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, greeting)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// getID prints ID of greeting.
func getID(db *sql.DB, greeting string) {
	stmt, err := db.Prepare("select id from greetings where greeting = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = stmt.Close() }()
	var id int
	err = stmt.QueryRow(greeting).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

// getGreeting prints greeting with ID.
func getGreeting(db *sql.DB, id int) {
	stmt, err := db.Prepare("select greeting from greetings where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = stmt.Close() }()
	var greeting string
	err = stmt.QueryRow(id).Scan(&greeting)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)
}

// deleteAll removes all entries from db.
func deleteAll(db *sql.DB) {
	_, err := db.Exec("delete from greetings")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// create db
	db, err := sql.Open("sqlite3", "file:greet.db?mode=memory")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	sqlStmt := `
	create table greetings (id integer not null primary key, greeting text);
		delete from greetings;
			`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// fill db
	addGreetings(db, []string{
		"hello",
		"hi",
		"good day",
		"greetings",
	})

	// command line arguments
	r := flag.Bool("run", false, "run some commands for testing")
	l := flag.Bool("list", false, "list greetings")
	i := flag.String("id", "", "get ID of greeting")
	g := flag.Int("greeting", 0, "get greeting with `id`")
	flag.Parse()

	if *r {
		list(db)
		getID(db, "good day")
		getGreeting(db, 2)
		deleteAll(db)
		list(db)
		addGreeting(db, 23, "bye")
		list(db)
	}

	if *l {
		list(db)
		return
	}
	if *i != "" {
		getID(db, *i)
		return
	}
	if *g != 0 {
		getGreeting(db, *g)
		return
	}
}
