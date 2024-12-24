package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HomeInfo struct {
	Porcupine   string
	Clustering  bool
	Nodes       []string
	DefaultPort string
}

func Serve() {
	http.HandleFunc("/configure", ConfigHandler)
	http.HandleFunc("/plot", plotHandler)
	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func plotHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderTemplate(w, "./assets/line.html", nil); err != nil {
		log.Printf("Failed to render template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	info := HomeInfo{"HashMap", false, []string{}, "8080"}
	if err := renderTemplate(w, "./assets/index.html", info); err != nil {
		log.Printf("Failed to render template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func renderTemplate(w http.ResponseWriter, filepath string, data interface{}) error {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	if err := tmpl.Execute(w, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
