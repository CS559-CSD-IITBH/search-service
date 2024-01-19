package routes

import (
	"github.com/CS559-CSD-IITBH/search-service/controllers"
	"github.com/CS559-CSD-IITBH/search-service/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SetupRouter(store *sessions.FilesystemStore) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	auth := middlewares.SessionAuth(store)
	r.Use(auth)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/search/global/:keyword", controllers.GlobalSearch)
		v1.GET("/search/store/:storeID/:keyword", controllers.StoreSearch)
	}

	return r
}
