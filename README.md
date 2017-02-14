# Beer Server

RESTful service for reviewing beers.

## Notes

Build to benefit the guidance of https://github.com/golang-bristol.

This has been heavily influenced by https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.xnb4d49aa, hat tip to @benbjohnson.

## Get Started

I've provided a db dump. Import that if you want to see data.

Note the DSN in below example. Modify it as necessary.

```bash
$ make build
$ export DB_DSN="root@tcp(127.0.0.1:3306)/beers"
$ ./beepapp &
$ curl http://localhost:3000/beers/1
{"beer":{"id":1,"name":"Best Bitter","brewery":"T\u0026R Theakston Ltd","abv":3.8,"short_description":"Theakston Best Bitter is the leading session ale within the Theakston portfolio and has been for time immemorial. It is quite possible that when Robert Theakston founded the brewery in 1827 the range of ales would have been limited to just two or three of which almost certainly, one would have been a bitter beer. Consequently it would be reasonable to argue that Theakston Best Bitter is one of the longest established session ales in Yorkshire.","created":"2017-01-09T18:55:43Z"}}
$ curl http://localhost:3000/beers/2/reviews
{"reviews":[{"id":1,"beer_id":2,"first_name":"Bob","last_name":"Thornton","score":4,"text":"Incredible beer, copper in colour.","created":"2017-01-09T19:00:59Z"},{"id":2,"beer_id":2,"first_name":"Ted","last_name":"Newton","score":1,"text":"Not the nicest beer.","created":"2017-01-09T12:30:12Z"}]}
```

