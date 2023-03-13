package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	"bitbucket.org/marcoboschetti/visionmonk/src/service"
)

func RegisterCatalogEndpoints(private *gin.RouterGroup) {
	private.POST("/global_catalog/product", RequireUser(postNewCatalogProduct))
	private.PUT("/global_catalog/product/:product_id", RequireUser(updateCatalogProduct))
	private.DELETE("/global_catalog/product/:product_id", RequireUser(deleteCatalogProduct))
	private.GET("/global_catalog/product/:product_id", RequireUser(getCatalogProduct))
	private.GET("/global_catalog/products", RequireUser(getAllShopCatalogProducts))
}

func postNewCatalogProduct(c *gin.Context, user *entities.User) {
	product := entities.CatalogProduct{}

	// TODO validate image size <= 1000000 (1MB)
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newProduct, err := service.PostNewCatalogProduct(&product, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"new_product_id": newProduct.ID})
}

func getCatalogProduct(c *gin.Context, user *entities.User) {
	productID := c.Query("product_id")
	product, err := service.GetCatalogProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func updateCatalogProduct(c *gin.Context, user *entities.User) {
	productID := c.Param("product_id")

	product := entities.CatalogProduct{}
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	product.ID = productID
	err = service.UpdateCatalogProduct(&product, user, &user.ShopID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product_id": product.ID})
}

func deleteCatalogProduct(c *gin.Context, user *entities.User) {
	productID := c.Param("product_id")

	err := service.DeleteCatalogProduct(productID, user, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": productID})
}

func getAllShopCatalogProducts(c *gin.Context, user *entities.User) {
	products, err := service.GetAllShopCatalogProducts(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"records": products})
}
