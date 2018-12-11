package main

type Teller interface {
	oust(another Teller) Teller
	say() string
}
