package main

type BothTeller struct {
	tellers []Teller
}

func NewBothTeller(five Teller, three Teller) Teller {
	result := new(BothTeller)

	result.tellers = append(result.tellers, three)
	result.tellers = append(result.tellers, five)

	return result
}

func (t *BothTeller) isDivisible() bool {
	return t.tellers[0].(*ThreeTeller).isDivisible() &&
		t.tellers[1].(*FiveTeller).isDivisible()
}

func (t *BothTeller) oust(another Teller) Teller {
	if t.isDivisible() {
		return t
	}

	return another

}

func (t *BothTeller) say() string {
	result := ""

	for _, t := range t.tellers {
		result = result + t.say()
	}

	return result
}
