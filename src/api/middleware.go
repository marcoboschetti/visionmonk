package api

import (
	"net/http"

	"bitbucket.org/marcoboschetti/visionmonk/src/data"
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	"github.com/gin-gonic/gin"
)

type endpointFunc func(c *gin.Context, user *entities.User)

func RequireUser(apiEndpoint endpointFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load user
		userToken := c.Request.Header.Get("x-auth-token")
		if len(userToken) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Authentication required"})
			return
		}

		user, err := data.GetUserByToken(userToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid authentication error with token "})
			return
		}

		apiEndpoint(c, &user)
	}
}
