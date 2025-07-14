package db

import "net/http"

func CORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        // Разрешаем CORS
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        // Если это preflight-запрос (OPTIONS) — просто завершаем
        if r.Method == "OPTIONS" {
            return
        }

        handler(w, r)
	}
}