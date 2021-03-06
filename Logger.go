package main

import (
	"log"
	"net/http"
	"time"
)

//Logger does what a logger do....
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		IPAddress := r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = r.RemoteAddr
		}
		log.Printf(
			"METHOD: %s | PATH: %s | HANDLER: %s | REQ_TIME: %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
