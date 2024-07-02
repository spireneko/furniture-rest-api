package app

import (
	"log"
	"net/http"

	"github.com/spireneko/furniture-rest-api/internal/app/service"
)

func Run(port string) error {
	handler := new(service.Service)

	router := http.NewServeMux()

	router.HandleFunc("POST /furniture", handler.Create)
	router.HandleFunc("GET /furniture", handler.GetAll)
	router.HandleFunc("GET /furniture/{id}", handler.Get)
	router.HandleFunc("PUT /furniture/{id}", handler.Update)
	router.HandleFunc("PATCH /furniture/{id}", handler.Patch)
	router.HandleFunc("DELETE /furniture/{id}", handler.Delete)

	server := http.Server{Addr: port, Handler: router}

	log.Printf("Run server on localhost%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Println("Error while starting server listening")
	}

	return nil
}
