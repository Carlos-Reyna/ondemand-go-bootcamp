package controller

import (
	"net/http"

	"github.com/Carlos-Reyna/go-api/domain"
	"github.com/Carlos-Reyna/go-api/service"
	"github.com/Carlos-Reyna/go-api/utils"
	"github.com/julienschmidt/httprouter"
)

func GetPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ch := make(chan []byte)
	go func() {
		param := ps.ByName("id")

		if param == "" {
			e := "Id cannot be empty"
			errorResponse := utils.ResponseWrapper(domain.Pokemon{}, e)
			ch <- errorResponse
		} else {
			response := service.GetPokemon(param)
			ch <- response
		}
	}()

	w.Write(<-ch)
}

func GetPokemons(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ch := make(chan []byte)
	go func() {
		query_type := r.URL.Query().Get("type")
		items := r.URL.Query().Get("items")
		item_worker := r.URL.Query().Get("items_per_workers")
		if query_type == "" || items == "" || item_worker == "" {
			e := "Query param cannot be empty"
			errorResponse := utils.ResponseWrapper(domain.Pokemon{}, e)
			ch <- errorResponse
		} else {
			data, err := service.GetPokemons(query_type, items, item_worker, "")
			response := utils.ArrayResponseWrapper(data, err)
			ch <- response
		}
	}()

	w.Write(<-ch)
}
