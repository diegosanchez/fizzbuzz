package main

type FiveTeller struct {
	number int
}

func NewFiveTeller(number int) Teller {
	result := new(FiveTeller)
	result.number = number
	return result
}

func (t *FiveTeller) isDivisible() bool {
	return t.number%5 == 0
}

func (t *FiveTeller) oust(another Teller) Teller {
	if t.isDivisible() {
		return t
	}

	return another
}

func (t *FiveTeller) say() string {
	if t.isDivisible() {
		return "Buzz"
	}

	return NewDefaultTeller(t.number).say()
}
