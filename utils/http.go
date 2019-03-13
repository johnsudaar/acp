package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
)

func JSON(ctx context.Context, resp http.ResponseWriter, data interface{}) {
	log := logger.Get(ctx)
	resp.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(resp).Encode(data)
	if err != nil {
		log.WithError(err).Error("fail to encode json")
	}
}

func Err(ctx context.Context, resp http.ResponseWriter, code int, message string) {
	resp.WriteHeader(code)
	JSON(ctx, resp, map[string]string{"error": message})
}
