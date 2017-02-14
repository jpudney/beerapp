package http_test

import (
	"testing"

	"log"
	"os"

	"net/http/httptest"

	"strings"

	"github.com/katzien/beerapp"
	"github.com/katzien/beerapp/http"
	"github.com/katzien/beerapp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetBeer(t *testing.T) {

	s := new(mocks.BeerStore)

	s.On("Beer", 1).Return(&beerapp.Beer{
		ID:   1,
		Name: "TestBeer",
	}, nil)

	h := http.NewHandler(s, log.New(os.Stdout, "logger: ", log.Lshortfile))

	r := httptest.NewRequest("GET", "http://example.com/beers/1", nil)
	w := httptest.NewRecorder()

	// need to go through the router so that the params
	// are picked up by chi
	h.ServeHTTP(w, r)

	if got := w.Code; !assert.Equal(t, 200, got) {
		t.Errorf("testgetbeer failed, unexpected status code: %d", got)
	}

	if got := strings.TrimSpace(w.Body.String()); !assert.Equal(t, `{"beer":{"id":1,"name":"TestBeer"}}`, got) {
		t.Errorf("testgetbeer failed, unexpected response body: %s", got)
	}

}
