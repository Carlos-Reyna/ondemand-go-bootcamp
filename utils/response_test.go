package utils

import (
	"encoding/json"
	"testing"

	"github.com/Carlos-Reyna/go-api/domain"
	"github.com/stretchr/testify/assert"
)

func TestResponseWrapper(t *testing.T) {
	pokemon := domain.Pokemon{
		Name: "Charmander",
		Id:   4,
	}
	errorMessage := "This is an error"

	b := ResponseWrapper(pokemon, errorMessage)
	var response domain.BaseResponse
	err := json.Unmarshal(b, &response)
	assert.Nil(t, err)
	assert.Equal(t, response.Data, pokemon)

	b = ResponseWrapper(pokemon, "")
	err = json.Unmarshal(b, &response)
	assert.Nil(t, err)
	assert.Equal(t, response.Data, pokemon)
}

func TestResponseArrayWrapper(t *testing.T) {
	pokemon := domain.Pokemon{
		Name: "Charmeleon",
		Id:   25,
	}
	var pokeSlice []domain.Pokemon
	pokeSlice = append(pokeSlice, pokemon)
	errorMessage := "This is an error"

	b := ArrayResponseWrapper(pokeSlice, errorMessage)
	var response domain.BaseResponse
	err := json.Unmarshal(b, &response)
	assert.Nil(t, err)
	assert.Equal(t, response.DataArray, pokeSlice)

	b = ArrayResponseWrapper(([]domain.Pokemon{}), errorMessage)
	err = json.Unmarshal(b, &response)
	assert.Nil(t, err)
	assert.Equal(t, response.DataArray, pokeSlice)
}

func TestResponseUnWrapper(t *testing.T) {
	pokemon := domain.Pokemon{
		Name: "Charizard",
		Id:   7,
	}

	// Test the function with a valid Pokemon
	b, _ := json.Marshal(pokemon)
	p, e := ResponseUnWrapper(b)
	assert.Equal(t, p, pokemon)
	assert.Equal(t, e, "")

	// Test the function with an invalid Pokemon
	b = []byte("invalid Pokemon")
	p, e = ResponseUnWrapper(b)
	assert.Equal(t, p, domain.Pokemon{})
	assert.Equal(t, e, "Poke Request did not work")
}
