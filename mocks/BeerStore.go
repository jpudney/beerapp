package mocks

import "github.com/jpudney/beerapp"
import "github.com/stretchr/testify/mock"

type BeerStore struct {
	mock.Mock
}

func (_m *BeerStore) Beers() ([]*beerapp.Beer, error) {
	ret := _m.Called()

	var r0 []*beerapp.Beer
	if rf, ok := ret.Get(0).(func() []*beerapp.Beer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*beerapp.Beer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *BeerStore) Beer(id int) (*beerapp.Beer, error) {
	ret := _m.Called(id)

	var r0 *beerapp.Beer
	if rf, ok := ret.Get(0).(func(int) *beerapp.Beer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*beerapp.Beer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *BeerStore) CreateBeer(b *beerapp.Beer) (*beerapp.Beer, error) {
	ret := _m.Called(b)

	var r0 *beerapp.Beer
	if rf, ok := ret.Get(0).(func(*beerapp.Beer) *beerapp.Beer); ok {
		r0 = rf(b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*beerapp.Beer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*beerapp.Beer) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *BeerStore) Reviews(id int) ([]*beerapp.Review, error) {
	ret := _m.Called(id)

	var r0 []*beerapp.Review
	if rf, ok := ret.Get(0).(func(int) []*beerapp.Review); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*beerapp.Review)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *BeerStore) CreateReview(r *beerapp.Review) (*beerapp.Review, error) {
	ret := _m.Called(r)

	var r0 *beerapp.Review
	if rf, ok := ret.Get(0).(func(*beerapp.Review) *beerapp.Review); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*beerapp.Review)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*beerapp.Review) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
