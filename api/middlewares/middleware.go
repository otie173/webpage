package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		log.Printf("Входящий запрос на страницу: %s\n", r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Обработка запроса заняла: %f секунд\n\n", time.Since(startTime).Seconds())
	})
}
