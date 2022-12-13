package service

import (
	"testing"

	"github.com/Carlos-Reyna/go-api/domain"
	"github.com/stretchr/testify/assert"
)

func TestSearchCSV(t *testing.T) {
	id := "25"

	expectedPokemon := domain.Pokemon{
		Id:   25,
		Name: "pikachu",
	}

	pokemon, e := SearchCSV(id, 6, "")

	if e != "" {
		t.Errorf("SearchCSV returned an error: %s", e)
	}
	if pokemon != expectedPokemon {
		t.Errorf("SearchCSV(%s) = %v, want %v", id, pokemon, expectedPokemon)
	}
}

func TestGetPokemons(t *testing.T) {

	pokemons, err := GetPokemons("odd", "5", "3", "")
	assert.Len(t, pokemons, 5, "Expected 5 odd pokemons")
	assert.Equal(t, err, "", "Expected no error")

	pokemonsEven, err := GetPokemons("even", "5", "3", "")
	assert.Len(t, pokemonsEven, 5, "Expected 5 even pokemons")
	assert.Equal(t, err, "", "Expected no error")

	//Items  = 0
	var emptyPokemons []domain.Pokemon
	emptyPokemons, err = GetPokemons("even", "0", "10", "")
	assert.Len(t, emptyPokemons, 0, "Expected no pokemons")
	assert.Equal(t, err, "User didn't request data", "User didn't request data")

	//Workers = 0
	emptyPokemons, err = GetPokemons("even", "1", "0", "")
	assert.Len(t, emptyPokemons, 0, "Expected no pokemons")
	assert.Equal(t, err, "Worker value cannot be 0 or lower than 0", "Expected no error")

}
