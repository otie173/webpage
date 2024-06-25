package main

import (
	"net/http"

	handler "github.com/otie173/webpage/api/handlers"
)

func main() {
	mux := http.NewServeMux()
	handler.RegisterHandlers(mux)

	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("static/", http.StripPrefix("static", fileServer))
	http.ListenAndServe("localhost:8080", mux)
}
