package utils

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware pour gérer les problèmes de CORS
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Afficher la route, la date et la méthode dans la console
		fmt.Printf("Route: %s, Date: %s, Method: %s\n",
			r.URL.Path,
			time.Now().Format("02/01/2006 15:04:05"),
			r.Method,
		)
		next.ServeHTTP(w, r)
	})
}
