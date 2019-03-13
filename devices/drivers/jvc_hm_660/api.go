package jvc

import "net/http"

func (a *JVCHM660) API() http.Handler {
	return http.NotFoundHandler()
}
