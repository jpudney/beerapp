package http

import (
	"log"
	"net/http"

	"encoding/json"

	"fmt"

	"strconv"

	"github.com/jpudney/beerapp"
	"github.com/pressly/chi"
)

type Handler struct {
	*chi.Mux

	BeerService beerapp.BeerStore
	Logger      *log.Logger
}

func NewHandler(s beerapp.BeerStore, l *log.Logger) *Handler {

	h := &Handler{
		Mux:         chi.NewRouter(),
		BeerService: s,
		Logger:      l,
	}

	h.Get("/beers", h.GetBeers)
	h.Post("/beers", h.CreateBeer)
	h.Get("/beers/:id", h.GetBeer)
	h.Get("/beers/:id/reviews", h.GetBeerReviews)
	h.Post("/beers/:id/review", h.CreateBeerReview)

	return h
}

type getBeersResponse struct {
	Beers []*beerapp.Beer `json:"beers"`
}

func (h *Handler) GetBeers(w http.ResponseWriter, r *http.Request) {
	beers, err := h.BeerService.Beers()

	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	h.encodeJSON(w, &getBeersResponse{beers})
}

type createResponse struct {
	ID int `json:"id"`
}

func (h *Handler) CreateBeer(w http.ResponseWriter, r *http.Request) {

	b := &beerapp.Beer{}

	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	beer, err := h.BeerService.CreateBeer(b)

	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	h.encodeJSON(w, &createResponse{ID: beer.ID})

}

type getBeerResponse struct {
	Beer *beerapp.Beer `json:"beer"`
}

func (h *Handler) GetBeer(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	b, err := h.BeerService.Beer(id)

	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	h.encodeJSON(w, &getBeerResponse{Beer: b})

}

type getBeerReviewsResponse struct {
	Reviews []*beerapp.Review `json:"reviews"`
}

func (h *Handler) GetBeerReviews(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	reviews, err := h.BeerService.Reviews(id)

	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	h.encodeJSON(w, &getBeerReviewsResponse{Reviews: reviews})

}

func (h *Handler) CreateBeerReview(w http.ResponseWriter, r *http.Request) {

	var review *beerapp.Review

	if err := json.NewDecoder(r.Body).Decode(review); err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	review, err := h.BeerService.CreateReview(review)

	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

	h.encodeJSON(w, &createResponse{ID: review.ID})

}

type errorResponse struct {
	Error error `json:"error"`
}

func (h *Handler) handleError(w http.ResponseWriter, err error, code int) {

	h.Logger.Printf("error occurred: %s", err)

	if code == http.StatusInternalServerError {
		err = fmt.Errorf(http.StatusText(code))
	}

	json.NewEncoder(w).Encode(&errorResponse{Error: err})
}

func (h *Handler) encodeJSON(w http.ResponseWriter, v interface{}) {

	// try and encode the json response, if it fails handle that too
	if err := json.NewEncoder(w).Encode(v); err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
	}

}
