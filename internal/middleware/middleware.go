package middleware

import "net/http"

// CorsMiddleware handles CORS preflight requests and adds CORS headers
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from specific origin(s)
		w.Header().Set("Access-Control-Allow-Origin", "*") // or specify a specific origin like "http://localhost:3000"
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If it's a preflight request (OPTIONS method), just send the response and return
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed with the next handler if not OPTIONS
		next.ServeHTTP(w, r)
	})
}
