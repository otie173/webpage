package main

import (
	"net/http"

	handler "github.com/otie173/webpage/api/handlers"
)

func main() {
	handler.RegisterHandlers()

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("static/", http.StripPrefix("static", fileServer))
	http.ListenAndServe("localhost:8080", nil)
}
