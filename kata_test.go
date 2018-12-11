package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestNumberTellerSayOne(t *testing.T) {
	teller := NewTeller()

	assert.Equal(t, string(1), teller.say(1))
}
