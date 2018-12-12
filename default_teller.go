package main

type DefaultTeller struct {
	number int
}

func NewDefaultTeller(number int) Teller {
	result := new(DefaultTeller)
	result.number = number
	return result
}

func (d *DefaultTeller) isDivisible() bool {
	return true
}

func (d *DefaultTeller) oust(another Teller) Teller {
	return another
}

func (d *DefaultTeller) say() string {
	return string(d.number)
}
