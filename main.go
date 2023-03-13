package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"bitbucket.org/marcoboschetti/visionmonk/src/api"
	"bitbucket.org/marcoboschetti/visionmonk/src/data"
)

func main() {

	gin.SetMode(gin.DebugMode)
	rand.Seed(time.Now().Unix())

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set. Defaulted to 8080")
		port = "8080"
	}

	data.SetDbConnection()

	// err := data.InsertNewShop(&entities.Shop{
	// 	ID:            utils.GetRandomString(3),
	// 	DateCreated:   time.Now(),
	// 	OwnerIDs:      []string{},
	// 	SecretToken:   utils.NewShopSecretToken(),
	// 	Email:         "test@shop.com",
	// 	Tier:          "1",
	// 	Status:        "active",
	// 	Name:          "Negocio de Prueba",
	// 	Addess:        "San Martin 3213",
	// 	Locality:      "CABA",
	// 	Neighborhood:  "CABA",
	// 	ZipCode:       "1234",
	// 	AditionalInfo: "",
	// 	UsersCount:    0,
	// 	ClientsCount:  0,
	// })
	// fmt.Println("ERR", err)

	// Start server
	r := gin.Default()
	// // *************** API **************
	public := r.Group("/api")
	public.POST("/user/new", api.PostNewUser)
	public.GET("/login", api.Login)

	// Private. Request an authenticated user header token
	private := public.Group("/p")
	private.GET("/user", api.RequireUser(api.GetUser))

	// Entities CRUD
	api.RegisterClientEndpoints(private)
	api.RegisterCatalogEndpoints(private)
	api.RegisterShopProductEndpoints(private)

	// TODO: ADMIN role
	public.GET("/user/:user_token", api.GetUserByToken)

	// *************** SITE **************
	// Public Static Resources
	r.GET("/", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./site/index.html") })
	r.GET("/favicon.ico", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./site/logo.ico") })
	publicSite := r.Group("/site")
	publicSite.Static("/", "./site/")

	r.Run(":" + port)
}
