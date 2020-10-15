package realtime

import (
	"context"
	"net/http"
	"strings"

	"github.com/centrifugal/centrifuge"
)

func (r *RealtimeServer) Websocket(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Access-Control-Allow-Origin", strings.TrimSuffix(req.Referer(), "/"))
	ctx := context.Background()
	cred := &centrifuge.Credentials{
		UserID: "",
	}
	newCtx := centrifuge.SetCredentials(ctx, cred)
	req = req.WithContext(newCtx)
	r.websocketHandler.ServeHTTP(resp, req)
}

func (r *RealtimeServer) SockJS(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Access-Control-Allow-Origin", strings.TrimSuffix(req.Referer(), "/"))
	ctx := context.Background()
	cred := &centrifuge.Credentials{
		UserID: "",
	}
	newCtx := centrifuge.SetCredentials(ctx, cred)
	req = req.WithContext(newCtx)

	r.sockjsHandler.ServeHTTP(resp, req)
}
