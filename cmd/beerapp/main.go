package main

import (
	"log"
	"os"

	"github.com/katzien/beerapp/cache"
	"github.com/katzien/beerapp/http"
	"github.com/katzien/beerapp/mysql"
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
