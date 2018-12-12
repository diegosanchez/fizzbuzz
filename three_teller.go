package main

type ThreeTeller struct {
	number int
}

func NewThreeTeller(number int) Teller {
	result := new(ThreeTeller)
	result.number = number
	return result
}

func (t *ThreeTeller) isDivisible() bool {
	return t.number%3 == 0
}

func (t *ThreeTeller) oust(another Teller) Teller {
	if t.isDivisible() {
		return t
	}

	return another

}

func (t *ThreeTeller) say() string {
	return "Fizz"
}
