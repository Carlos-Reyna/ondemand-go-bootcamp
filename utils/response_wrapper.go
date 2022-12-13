package utils

import (
	"encoding/json"
	"log"

	"github.com/Carlos-Reyna/go-api/domain"
)

func ResponseWrapper(r domain.Pokemon, e string) []byte {
	var response domain.BaseResponse
	response.Data = r
	response.SetErrorMessage(e)
	b, err := json.Marshal(response)

	if err != nil {
		log.Fatal("Could not parse response")
	}

	return b
}

func ArrayResponseWrapper(r []domain.Pokemon, e string) []byte {
	var response domain.BaseResponse
	response.DataArray = r
	response.SetErrorMessage(e)
	b, err := json.Marshal(response)

	if err != nil {
		log.Fatal("Could not parse response")
	}

	return b
}

func ResponseUnWrapper(r []byte) (domain.Pokemon, string) {
	var pokemon domain.Pokemon
	err := json.Unmarshal(r, &pokemon)
	var e string = ""
	if err != nil {
		e = "Poke Request did not work"

		return domain.Pokemon{}, e
	}
	return pokemon, e
}
