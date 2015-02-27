package main

import(
	"fmt"
	"time"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

type Page struct {
	Id        int64     `db:"pk"`
	Title     string    `db:"unique" json:"title"`
	Body      string    `json:"body"`
	URL       string    `db:"-"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main(){
	db, err := genmai.New(&genmai.SQLite3Dialect{}, "./sample.db")
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error: %s\n",err)
		return
	}
	if err := db.CreateTableIfNotExists(&Page{}); err != nil {
		fmt.Fprintf(os.Stderr,"Error: %s\n",err)
		return
	}
	records := []Page{
		{Title: "Top Page",Body: "Top Contents"},
		{Title: "Next Page",Body: "Next Contents"},
	}
	if _, err := db.Insert(records); err != nil {
		panic(err)
	}

	var pages []Page
	if err := db.Select(&pages); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", pages)
}
	
