package handler

import (
	"net/http"
	"os"

	middleware "github.com/otie173/webpage/api/middlewares"
)

func RegisterHandlers() {
	http.Handle("GET /", middleware.LoggingMiddleware(http.HandlerFunc(sayHello)))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	htmlTemplate, err := os.ReadFile("static/templates/index.html")
	if err != nil {
		http.Error(w, "Error loading index.html template", http.StatusInternalServerError)
	}

	w.Header().Set("Context-Type", "text-html")
	w.Write(htmlTemplate)
}
