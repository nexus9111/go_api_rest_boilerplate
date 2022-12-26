package controllers

import (
	"fmt"
	"net/http"
)

// Fonction appelée pour chaque route
func HelloWorlController(w http.ResponseWriter, r *http.Request) {
	// TODO: implémenter la logique de la route ici
	fmt.Fprintf(w, "Hello, World!")
}
