package atem

import "net/http"

func (a *ATEM) API() http.Handler {
	return http.NotFoundHandler()
}
