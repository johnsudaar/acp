package hs50

import "net/http"

func (s *Switcher) API() http.Handler {
	return http.NotFoundHandler()
}
