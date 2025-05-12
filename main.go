package main

import (
	"log"
	"net/http"
	"quesos/routes"
	"quesos/utils"
)

func main(){
	//Se inicia la DB
	utils.InitDB()

	//Se configuran las rutas
	router := routes.SetupRoutes()

	//Iniciar el server
	log.Println("Servidiro inicuado en :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}