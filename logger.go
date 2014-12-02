package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%12s\t%6s\t%12s\t%s",
			time.Since(start),
			r.Method,
			name,
			r.RequestURI,
		)
	})
}
