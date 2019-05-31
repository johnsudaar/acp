package smartview

import "net/http"

func (a *SmartView) API() http.Handler {
	return http.NotFoundHandler()
}
