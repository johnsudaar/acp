package tally

import "net/http"

func (a *Tally) API() http.Handler {
	return http.NotFoundHandler()
}
