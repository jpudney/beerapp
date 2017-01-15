package cache

import "github.com/jpudney/beerapp"

type BeerCache struct {
	beers   map[int]*beerapp.Beer
	reviews map[int][]*beerapp.Review

	store beerapp.BeerStore
}

func NewBeerCache(store beerapp.BeerStore) (beerapp.BeerStore, error) {

	beers, err := store.Beers()

	if err != nil {
		return nil, err
	}

	cache := &BeerCache{
		beers:   make(map[int]*beerapp.Beer),
		reviews: make(map[int][]*beerapp.Review),
		store:   store,
	}

	for _, b := range beers {
		cache.beers[b.ID] = b
	}

	return cache, nil
}

// Beers returns all the beers from the cache. It should probably check
// the underlying store if the map of beers is empty
func (c *BeerCache) Beers() ([]*beerapp.Beer, error) {

	beers := make([]*beerapp.Beer, 0)

	for _, b := range c.beers {
		beers = append(beers, b)
	}

	return beers, nil

}

// Beer looksup a beer in the cache. If nothing exists, we hit the store
// and store any results in the cache.
func (c *BeerCache) Beer(id int) (*beerapp.Beer, error) {

	if b, ok := c.beers[id]; ok {
		return b, nil
	}

	b, err := c.store.Beer(id)

	if err != nil {
		return nil, err
	}

	c.beers[b.ID] = b

	return b, nil

}

// CreateBeer creates a beer in the underlying store and then stores the
// resulting beer in the cache store. The result is then resturned.
func (c *BeerCache) CreateBeer(b *beerapp.Beer) (*beerapp.Beer, error) {

	beer, err := c.store.CreateBeer(b)

	if err != nil {
		return nil, err
	}

	c.beers[beer.ID] = beer

	return beer, nil

}

// Reviews returns a review for a beer ID. It checks for cached reviews before
// hitting the store.
func (c *BeerCache) Reviews(id int) ([]*beerapp.Review, error) {

	if reviews, ok := c.reviews[id]; ok {
		return reviews, nil
	}

	reviews, err := c.store.Reviews(id)

	if err != nil {
		return nil, err
	}

	c.reviews[id] = reviews

	return reviews, nil

}

// CreateReview creates a review in the underlying store and caches the result.
func (c *BeerCache) CreateReview(r *beerapp.Review) (*beerapp.Review, error) {

	review, err := c.store.CreateReview(r)

	if err != nil {
		return nil, err
	}

	for k, v := range c.reviews[r.BeerID] {
		if v.ID == review.ID {
			c.reviews[r.BeerID][k] = review
		}
	}

	return review, nil

}
