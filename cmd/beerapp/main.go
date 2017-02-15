package main

import (
	"log"
	"os"

	"github.com/jpudney/beerapp/cache"
	"github.com/jpudney/beerapp/http"
	"github.com/jpudney/beerapp/mysql"
)

func main() {
	db, err := mysql.Open(os.Getenv("DB_DSN") + "?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	cache, err := cache.NewBeerCache(&mysql.BeerStore{DB: db})

	if err != nil {
		log.Fatal(err)
	}

	h := http.NewHandler(cache, log.New(os.Stdout, "BeerApp: ", log.LstdFlags))

	s := http.NewServer(":3000", h)
	defer s.Close()

	log.Fatal(s.Open())
}
