package main

type Tellers struct {
}

func NewTellers() *Tellers {
	result := new(Tellers)

	return result
}

func (t *Tellers) say(number int) string {
	teller := NewDefaultTeller(number)

	teller = teller.oust(NewThreeTeller(number))
	teller = teller.oust(NewFiveTeller(number))

	return teller.say()
}
