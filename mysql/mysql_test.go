package mysql_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/katzien/beerapp"
	"github.com/katzien/beerapp/mysql"
	"github.com/stretchr/testify/assert"
)

func TestBeerService_Beers_AllTheBeers(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("unexpected error '%s' stubbing db", err)
	}

	defer db.Close()

	columns := []string{"id", "name", "brewery", "abv", "short_description", "created"}

	now := time.Now()

	mock.ExpectQuery("SELECT id, name, brewery, abv, short_description, created FROM beers WHERE id = \\?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Best Bitter", "T&R Theakston Ltd", 3.8, "Theakston Best Bitter is the leading session ale within the Theakston portfolio and has been for time immemorial. It is quite possible that when Robert Theakston founded the brewery in 1827 the range of ales would have been limited to just two or three of which almost certainly, one would have been a bitter beer. Consequently it would be reasonable to argue that Theakston Best Bitter is one of the longest established session ales in Yorkshire.", now))

	s := &mysql.BeerStore{
		DB: db,
	}

	b, err := s.Beer(1)

	if err != nil {
		t.Fatalf("unexpected error retrieving beer: %s", err)
	}

	beer := &beerapp.Beer{
		ID:               1,
		Name:             "Best Bitter",
		Brewery:          "T&R Theakston Ltd",
		Abv:              3.8,
		ShortDescription: "Theakston Best Bitter is the leading session ale within the Theakston portfolio and has been for time immemorial. It is quite possible that when Robert Theakston founded the brewery in 1827 the range of ales would have been limited to just two or three of which almost certainly, one would have been a bitter beer. Consequently it would be reasonable to argue that Theakston Best Bitter is one of the longest established session ales in Yorkshire.",
		Created:          &now,
	}

	if !assert.Equal(t, beer, b) {
		t.Fatalf("incorrect beer returned, got: %v, want: %v", b, beer)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("db mock expectations not met: %s", err)
	}
}
