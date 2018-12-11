package main

type Tellers struct {
}

func NewTellers() *Tellers {
	return new(Tellers)
}

func (t *Tellers) say(number int) string {
	if number == 3 {
		return "Fizz"
	}

	return string(number)
}
