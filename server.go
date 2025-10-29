package main

import (
	"go_htmx/config"
	"go_htmx/service"
	"log"
	"net/http"
)

func main() {
	port := config.GetEnv("PORT")

	http.HandleFunc("/", service.RootHandler)
	http.HandleFunc("/increment", service.IncrementHandler)

	log.Printf("Go HTMX Server starting on http://localhost:%s", port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
