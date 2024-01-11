package routes

import (
    "github.com/CS559-CSD-IITBH/search-service/controllers"
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} 
	r.Use(cors.New(config))

	v1 := r.Group("/api/v1")
	{
    	v1.GET("/search/:keyword", controllers.SearchHandler)
	}

	return r
}
