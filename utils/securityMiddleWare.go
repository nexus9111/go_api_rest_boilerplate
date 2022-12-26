package utils

import (
	"net/http"
	"strings"
)

// Fonction pour vérifier si l'IP est blacklistée
func isIPBlacklisted(ip string) bool {
	// TODO: implémenter la logique de vérification de l'IP blacklistée ici
	return false
}

// Middleware pour vérifier que l'IP n'est pas blacklistée
func IpCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vérifier si l'IP est blacklistée
		ip := strings.Split(r.RemoteAddr, ":")[0]
		if isIPBlacklisted(ip) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
