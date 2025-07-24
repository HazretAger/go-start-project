package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc{
	return func(c *gin.Context) {
        // Разрешаем CORS
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Если это preflight-запрос (OPTIONS) — просто завершаем
        if c.Request.Method == "OPTIONS" {
            return
        }
	}
}