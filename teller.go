package main

type Teller struct{}

func NewTeller() *Teller {
	return new(Teller)
}

func (t *Teller) say(number int) string {
	return string(number)
}
