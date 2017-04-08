package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/katzien/beerapp"
)

// BeerStore represents a MySQL implementation of beerapp.BeerStore
type BeerStore struct {
	DB *sql.DB
}

// Open a database connection
func Open(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}

// Beer returns a beer for the given ID
func (s *BeerStore) Beer(id int) (*beerapp.Beer, error) {
	b := new(beerapp.Beer)

	row := s.DB.QueryRow("SELECT id, name, brewery, abv, short_description, created FROM beers WHERE id = ?", id)

	if err := row.Scan(&b.ID, &b.Name, &b.Brewery, &b.Abv, &b.ShortDescription, &b.Created); err != nil {
		return nil, err
	}

	return b, nil
}

// Beers returns all beers
func (s *BeerStore) Beers() ([]*beerapp.Beer, error) {
	beers := make([]*beerapp.Beer, 0)

	rows, err := s.DB.Query("SELECT id, name, brewery, abv, short_description, created FROM beers")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		b := new(beerapp.Beer)

		if err := rows.Scan(&b.ID, &b.Name, &b.Brewery, &b.Abv, &b.ShortDescription, &b.Created); err != nil {
			return nil, err
		}

		beers = append(beers, b)
	}

	return beers, nil
}

// CreateBeer creates a beer
func (s *BeerStore) CreateBeer(b *beerapp.Beer) (*beerapp.Beer, error) {

	stmt, err := s.DB.Prepare("INSERT INTO beers (name, brewery, abv, short_description) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(b.Name, b.Brewery, b.Abv, b.ShortDescription)

	if err != nil {
		return nil, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	beer := &beerapp.Beer{}

	// select the inserted result
	row := s.DB.QueryRow("SELECT id, name, brewery, abv, short_description, created FROM beers WHERE id = ?", lastInsertedID)

	if err := row.Scan(&b.ID, &b.Name, &b.Brewery, &b.Abv, &b.ShortDescription, &b.Created); err != nil {
		// may want to check err == sql.ErrNoRows and return a better error
		return nil, err
	}

	return beer, nil
}

// Reviews returns reviews the given beer ID
func (s *BeerStore) Reviews(id int) ([]*beerapp.Review, error) {
	reviews := make([]*beerapp.Review, 0)

	rows, err := s.DB.Query("SELECT id, beer_id, first_name, last_name, score, text, created FROM reviews WHERE beer_id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return reviews, nil
		}
		return nil, err
	}

	for rows.Next() {
		r := new(beerapp.Review)

		if err := rows.Scan(&r.ID, &r.BeerID, &r.FirstName, &r.LastName, &r.Score, &r.Text, &r.Created); err != nil {
			return nil, err
		}

		reviews = append(reviews, r)
	}

	return reviews, nil
}

// CreateReview creates a review
func (s *BeerStore) CreateReview(r *beerapp.Review) (*beerapp.Review, error) {

	stmt, err := s.DB.Prepare("INSERT INTO reviews (beer_id, first_name, last_name, score, text) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(r.BeerID, r.FirstName, r.LastName, r.Score, r.Text)

	if err != nil {
		return nil, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	// select the inserted result
	review := &beerapp.Review{}

	row := s.DB.QueryRow("SELECT id, beer_id, first_name, last_name, score, text, created FROM reviews WHERE id = ?", lastInsertedID)

	if err := row.Scan(review); err != nil {
		// may want to check err == sql.ErrNoRows and return a better error
		return nil, err
	}

	return review, nil
}
