package main

import "weightApi/internal/api/router"

func main() {
	router := router.RegisterRoutes()
	router.Run(":8080") // Start on localhost:8080
}
