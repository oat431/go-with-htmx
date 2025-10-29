package main

import (
	"go_htmx/service"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error while reading .env")
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatal("Invalid type")
	}

	return value
}

func main() {
	port := viperEnvVariable("PORT")

	http.HandleFunc("/", service.RootHandler)
	http.HandleFunc("/increment", service.IncrementHandler)

	log.Printf("Go HTMX Server starting on http://localhost:%s", port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
