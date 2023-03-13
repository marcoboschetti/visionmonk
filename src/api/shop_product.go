package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	"bitbucket.org/marcoboschetti/visionmonk/src/service"
)

func RegisterShopProductEndpoints(private *gin.RouterGroup) {
	private.POST("/shop_product", RequireUser(postNewShopProduct))
	private.PUT("/shop_product/:product_id", RequireUser(updateShopProduct))
	private.DELETE("/shop_product/:product_id", RequireUser(deleteShopProduct))
	private.GET("/shop_product/:product_id", RequireUser(getShopProduct))
	private.GET("/shop_products", RequireUser(getAllShopShopProducts))
}

func postNewShopProduct(c *gin.Context, user *entities.User) {
	productDto := entities.ShopProduct{}

	// TODO validate image size <= 1000000 (1MB) from catalog
	err := c.ShouldBindJSON(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if productDto.CatalogProduct == nil && productDto.CatalogProductID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Errorf("catalog product unassigned")})
		return
	}

	err = service.PostNewShopProduct(productDto, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func getShopProduct(c *gin.Context, user *entities.User) {
	productID := c.Query("product_id")
	product, err := service.GetShopProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func updateShopProduct(c *gin.Context, user *entities.User) {
	productID := c.Param("product_id")
	productDto := entities.ShopProduct{}

	// TODO validate image size <= 1000000 (1MB) from catalog
	err := c.ShouldBindJSON(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	productDto.ID = productID
	err = service.UpdateShopProduct(&productDto, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product_id": productDto.ID})
}

func deleteShopProduct(c *gin.Context, user *entities.User) {
	productID := c.Param("product_id")

	err := service.DeleteShopProduct(productID, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": productID})
}

func getAllShopShopProducts(c *gin.Context, user *entities.User) {
	products, err := service.GetAllShopShopProducts(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"records": products})
}
