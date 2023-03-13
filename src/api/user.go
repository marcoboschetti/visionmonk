package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	"bitbucket.org/marcoboschetti/visionmonk/src/service"
	"github.com/gin-gonic/gin"
)

func PostNewUser(c *gin.Context) {
	input := struct {
		Password  string `json:"password"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		ShopToken string `json:"shop_token"`
	}{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	newUser, err := service.PostNewUser(input.Password, input.Firstname, input.Lastname, input.Email, c.ClientIP(), input.ShopToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": newUser})
}

func GetUserByToken(c *gin.Context) {
	userToken := c.Param("user_token")

	user, err := service.GetUserByToken(userToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUser(c *gin.Context, user *entities.User) {
	shop, err := service.GetShopByID(user.ShopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Shop = shop
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {
	tokenBase64 := c.Request.Header.Get("x-auth-token-bearer")

	rawDecoded, err := base64.StdEncoding.DecodeString(tokenBase64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	rawDecodedStr := fmt.Sprintf("%s", rawDecoded)
	s := strings.Split(string(rawDecodedStr), "&!&")
	user, err := service.GetUserByCredentials(s[0], s[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"secret_token": user.SecretToken})
}
