package main

type ThreeTeller struct {
	number int
}

func NewThreeTeller(number int) Teller {
	result := new(ThreeTeller)
	result.number = number
	return result
}

func (t *ThreeTeller) oust(another Teller) Teller {
	if t.number == 3 {
		return t
	}

	return another

}

func (t *ThreeTeller) say() string {
	if t.number == 3 {
		return "Fizz"
	}

	return NewDefaultTeller(t.number).say()
}
