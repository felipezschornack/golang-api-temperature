package weatherapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherApi(t *testing.T) {

	data, err := GetWeather("Porto Alegre", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Status, 500)
	assert.Nil(t, data)

	data, err = GetWeather("Porto Alegre", "8b3b137bec164888a27140948240303")
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Greater(t, data.Celsius, float32(0))
	assert.Greater(t, data.Fahrenheit, float32(0))
	assert.Greater(t, data.Kelvin, float32(0))

	data, err = GetWeather("Giruá", "8b3b137bec164888a27140948240303")
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Celsius, float32(0))
	assert.Equal(t, data.Fahrenheit, float32(0))
	assert.Greater(t, data.Kelvin, float32(0))

}
