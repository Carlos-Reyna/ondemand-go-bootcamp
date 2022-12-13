package main

import (
	"net/http"

	"github.com/Carlos-Reyna/go-api/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	router.GET("/read", controller.GetPokemons)
	router.GET("/read/:id", controller.GetPokemon)

	http.ListenAndServe(":8080", router)
}
