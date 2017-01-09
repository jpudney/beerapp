package main

import (
	"log"
	"os"

	"github.com/jpudney/beerapp/http"
	"github.com/jpudney/beerapp/mysql"
)

func main() {
	db, err := mysql.Open(os.Getenv("DB_DSN") + "?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	h := http.NewHandler(&mysql.BeerService{DB: db}, log.New(os.Stdout, "BeerApp: ", log.LstdFlags))

	s := http.NewServer(":3000", h)
	defer s.Close()

	log.Fatal(s.Open())
}
