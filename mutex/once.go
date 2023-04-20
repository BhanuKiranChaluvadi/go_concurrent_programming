package main

import (
	_ "github.com/mattn/go-sqlite3"
)

/*
func main() {
	Run()
}

var db *sql.DB
var o sync.Once

func Run() {
	o.Do(func() {
		log.Println("opening database connection")
		var err error
		db, err = sql.Open("mysql", "./mydb.db")
		if err != nil {
			log.Fatal(err)
		}
	})
}

*/
