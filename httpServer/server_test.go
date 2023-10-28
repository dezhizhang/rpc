package httpServer

import (
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	var s Server
	http.ListenAndServe(":8080", s)
}
