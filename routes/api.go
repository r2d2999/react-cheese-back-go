package routes

import (
	"net/http"
	"quesos/controllers"
	"quesos/utils"
	"github.com/rs/cors"

)

func SetupRoutes() http.Handler {  
    router := http.NewServeMux()

    // Rutas
    quesoController := controllers.NewQuesoController(utils.DB)
    router.HandleFunc("POST /quesos", quesoController.CreateQuesoCtlr)
    router.HandleFunc("GET /quesos/{id}", quesoController.GetQuesoByIDCtlr)
    router.HandleFunc("GET /quesos", quesoController.GetAllQuesos)

    //  CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:4200"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    return c.Handler(router) 
}