package main

import (
	"fmt"
	"go-rest-api-boiletplate/app/controllers"
	"go-rest-api-boiletplate/app/utils"
	"net/http"
	"os"
)

func main() {
	utils.LoadEnvVariables()
	__PORT__ := os.Getenv("PORT")

	// Créer un serveur HTTP
	httpServer := http.Server{
		Addr: ":" + __PORT__,
	}

	ipCheck := utils.IpCheckMiddleware
	cors := utils.CorsMiddleware
	logger := utils.LoggerMiddleware

	// Créer une route qui utilise les deux middlewares et la fonction handleRoute
	http.Handle("/", ipCheck(cors(logger(http.HandlerFunc(controllers.HelloWorlController)))))

	// Démarrer le serveur HTTP
	fmt.Println("Server started on port " + __PORT__)
	httpServer.ListenAndServe()
}
