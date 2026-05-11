package routes

import (
	"net/http"

	"github.com/desirorie/job-tracker-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	healthRoutes(v1)
	applicationRoutes(v1)
}

func healthRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "successful start",
		})
	})
}
func applicationRoutes(rg *gin.RouterGroup) {

	rg.POST("/applications", handlers.CreateApplication)
	rg.GET("/applications", handlers.GetApplications)
	rg.GET("/applications/:id", handlers.GetApplicationById)
	rg.DELETE("/applications/:id", handlers.DeleteApplication)

}
