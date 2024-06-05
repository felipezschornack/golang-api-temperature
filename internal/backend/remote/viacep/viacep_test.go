package viacep

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCep(t *testing.T) {

	data, err := BuscaCep("91360000")
	assert.Nil(t, err)
	assert.Equal(t, "Porto Alegre", data.Localidade)

	data, err = BuscaCep("")
	assert.Nil(t, data)
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Message, "invalid zipcode")

	data, err = BuscaCep("sadasdas")
	assert.Nil(t, data)
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Message, "invalid zipcode")

	data, err = BuscaCep("01153001")
	assert.Nil(t, data)
	assert.Equal(t, err.Status, http.StatusNotFound)
	assert.Equal(t, err.Message, "can not find zipcode")
}
