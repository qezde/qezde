package router

import (
	"gateway/pkg/server/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() (r *gin.Engine) {
	r = gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(response.MethodNotAllowedMiddleware())

	return
}