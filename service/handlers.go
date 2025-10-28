package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
	tmpl    *template.Template
)

func init() {
	var err error
	tmpl, err = template.ParseFiles("index.htmx")
	if err != nil {
		log.Fatalf("Error loading template 'index.htmx': %v", err)
	}
	log.Println("Template loaded successfully.")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Count int
	}{
		Count: counter,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
	}
}

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	newCount := counter
	mu.Unlock()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fragment := fmt.Sprintf(`<div class="text-6xl font-black text-gray-800 mb-8 p-4 border-b-4 border-indigo-200" id="count-display">%d</div>`, newCount)

	_, err := w.Write([]byte(fragment))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
