package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokeAPIHTTPClient_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test response"))
	}))
	defer ts.Close()

	c := &PokeAPIHTTPClient{}

	body, err := c.Get(ts.URL)

	assert.NoError(t, err)
	assert.Equal(t, "Test response", string(body))
}
