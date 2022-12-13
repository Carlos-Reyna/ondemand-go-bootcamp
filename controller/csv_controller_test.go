package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemon(t *testing.T) {
	// Test with valid input
	req, err := http.NewRequest("GET", "/pokemon/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pokemon/:id", GetPokemon)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Test with empty id
	req, err = http.NewRequest("GET", "/pokemon/", nil)
	assert.NoError(t, err)

	rr = httptest.NewRecorder()
	router = httprouter.New()
	router.GET("/pokemon/:id", GetPokemon)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

}
