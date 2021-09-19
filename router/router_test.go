package router

import (
	"bytes"
	"fmt"
	"minesweeper-api/factories"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	DomainLayersFactory := factories.NewDomainLayersFactory()
	engine := New(DomainLayersFactory)
	s := httptest.NewServer(engine)

	response, _ := http.Get(fmt.Sprintf("%s/ping", s.URL))

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(response.Body)
	responseBodyString := buf.String()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "{\"message\":\"pong\"}", responseBodyString)

	s.Close()
}
