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

func TestNumberTellerSayCinco(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, "Buzz", teller.say(5))
}

func TestNumberTellerSaySeis(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, "Fizz", teller.say(6))
}

func TestNumberTellerSaySiete(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(7), teller.say(7))
}

func TestNumberTellerSayOcho(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(8), teller.say(8))
}

func TestNumberTellerSayNueve(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, "Fizz", teller.say(9))
}

func TestNumberTellerSayDiez(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, "Buzz", teller.say(10))
}

func TestNumberTellerSayOnce(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(11), teller.say(11))
}

func TestNumberTellerSayDoce(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, "Fizz", teller.say(12))
}

func TestNumberTellerSayTrece(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(13), teller.say(13))
}

func TestNumberTellerSayCatorce(t *testing.T) {
	teller := NewTellers()

	assert.Equal(t, string(14), teller.say(14))
}
