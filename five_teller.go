package main

type FiveTeller struct {
	number int
}

func NewFiveTeller(number int) Teller {
	result := new(FiveTeller)
	result.number = number
	return result
}

func (t *FiveTeller) oust(another Teller) Teller {
	if t.number == 5 {
		return t
	}

	return another
}

func (t *FiveTeller) say() string {
	if t.number == 5 {
		return "Buzz"
	}

	return NewDefaultTeller(t.number).say()
}
