package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateTest(t *testing.T) {
	expected := Contact{ID: 1, FirstName: "test", LastName: "test", Phone: "test", Email: "test", Position: "test"}
	actual := createContact(Contact{})
	assert.Equal(t, expected, actual, "contact creation")
}
