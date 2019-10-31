package main

import (
	"log"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("logger", err)
	}

	logger.Info("starting")

	mux := http.NewServeMux()

	mux.HandleFunc("/-/health", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "OK", http.StatusOK)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := *r.URL
		u.Scheme = "https"
		u.Host = r.Host
		http.Redirect(w, r, u.String(), 301)
		logger.Info(
			"redirect",
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("request_uri", r.RequestURI),
			zap.String("method", r.Method),
			zap.String("user_agent", r.UserAgent()),
			zap.String("trace_context", r.Header.Get("x-cloud-trace-context")),
		)
	})

	logger.Info("listening on :8080")
	logger.Fatal("serve", zap.Error(http.ListenAndServe(":8080", mux)))
}
