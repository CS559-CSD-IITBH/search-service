package routes

import (
	"github.com/CS559-CSD-IITBH/search-service/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/search/:keyword", controllers.GlobalSearch)
		v1.GET("/search/:storeID/:keyword", controllers.StoreSearch)
	}

	return r
}
