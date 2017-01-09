package beerapp

import "time"

type Beer struct {
	ID               int        `json:"id,omitempty"`
	Name             string     `json:"name,omitempty"`
	Brewery          string     `json:"brewery,omitempty"`
	Abv              float32    `json:"abv,omitempty"`
	ShortDescription string     `json:"short_description,omitempty"`
	Created          *time.Time `json:"created,omitempty"`
}

type Review struct {
	ID        int        `json:"id,omitempty"`
	BeerID    int        `json:"beer_id,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Score     int        `json:"score,omitempty"`
	Text      string     `json:"text,omitempty"`
	Created   *time.Time `json:"created,omitempty"`
}

type BeerService interface {
	Beers() ([]*Beer, error)
	Beer(id int) (*Beer, error)
	CreateBeer(b *Beer) (int, error)
	Reviews(id int) ([]*Review, error)
	CreateReview(r *Review) (int, error)
}
