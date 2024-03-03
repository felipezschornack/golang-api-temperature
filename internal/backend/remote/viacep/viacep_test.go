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

	data, err = BuscaCep("99999999")
	assert.Nil(t, data)
	assert.Equal(t, err.Status, http.StatusUnprocessableEntity)
	assert.Equal(t, err.Message, "invalid zipcode")

	data, err = BuscaCep("98870-999")
	assert.Nil(t, data)
	assert.Equal(t, err.Status, http.StatusUnprocessableEntity)
	assert.Equal(t, err.Message, "invalid zipcode")
}
