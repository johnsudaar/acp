package hyperdeck

import "net/http"

func (h *Hyperdeck) API() http.Handler {
	return http.NotFoundHandler()
}
