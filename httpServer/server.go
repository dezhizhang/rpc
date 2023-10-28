package httpServer

import "net/http"

type Server interface {
	http.Handler
}
