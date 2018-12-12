package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestNumberTellerSayOne(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(1), teller.say(1))
}

func TestNumberTellerSayTwo(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(2), teller.say(2))
}

func TestNumberTellerSayTres(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, "Fizz", teller.say(3))
}

func TestNumberTellerSayCuatro(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(4), teller.say(4))
}
