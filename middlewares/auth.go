package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// AuthMiddleware checks if the user is authenticated as a user
func SessionAuth(store *sessions.FilesystemStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")
		_, ok := session.Values["user_type"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
