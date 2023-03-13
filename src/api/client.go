package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	"bitbucket.org/marcoboschetti/visionmonk/src/service"
)

func RegisterClientEndpoints(private *gin.RouterGroup) {
	private.POST("/client", RequireUser(postNewClient))
	private.PUT("/client/:client_id", RequireUser(updateClient))
	private.DELETE("/client/:client_id", RequireUser(deleteClient))
	private.GET("/client/:client_id", RequireUser(getClient))
	private.GET("/clients", RequireUser(getAllShopClients))
}

func postNewClient(c *gin.Context, user *entities.User) {
	client := entities.Client{}
	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = service.PostNewClient(&client, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"new_client_id": client.ID})
}

func getClient(c *gin.Context, user *entities.User) {
	clientID := c.Query("client_id")
	client, err := service.GetClientByID(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"client": client})
}

func updateClient(c *gin.Context, user *entities.User) {
	clientID := c.Param("client_id")

	client := entities.Client{}
	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	client.ID = clientID
	err = service.UpdateClient(&client, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"client_id": client.ID})
}

func deleteClient(c *gin.Context, user *entities.User) {
	clientID := c.Param("client_id")

	err := service.DeleteClient(clientID, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": clientID})
}

func getAllShopClients(c *gin.Context, user *entities.User) {
	clients, err := service.GetAllShopClients(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"records": clients})
}
