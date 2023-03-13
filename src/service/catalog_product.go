package service

import (
	"errors"
	"fmt"
	"time"

	"bitbucket.org/marcoboschetti/visionmonk/src/data"
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	uuid "github.com/satori/go.uuid"
)

func GetCatalogProductByID(productID string) (entities.CatalogProduct, error) {
	product, err := data.GetByID[entities.CatalogProduct](productID)
	return product, err
}

func UpdateCatalogProduct(inputCatalogProduct *entities.CatalogProduct, user *entities.User, shopIDFromProduct *string) error {
	// Retrieve product
	dbCatalogProduct, err := data.GetByID[entities.CatalogProduct](inputCatalogProduct.ID)
	if err != nil {
		return err
	}
	if dbCatalogProduct.IsGlobalCatalog && !user.IsAdmin() {
		return errors.New("user can't edit global catalog item")
	}
	if !dbCatalogProduct.IsGlobalCatalog && (shopIDFromProduct == nil && *shopIDFromProduct != user.ShopID) {
		return fmt.Errorf("unable to delete product %s (err 1492)", dbCatalogProduct.ID)
	}

	inputCatalogProduct.DateCreated = dbCatalogProduct.DateCreated
	inputCatalogProduct.DateUpdated = time.Now()
	inputCatalogProduct.Status = entities.EntityStatusActive

	// TODO: Log update user history record

	err = data.Update(inputCatalogProduct)
	return err
}

func DeleteCatalogProduct(productID string, user *entities.User, shopIDFromProduct *string) error {

	// Retrieve product
	dbCatalogProduct, err := data.GetByID[entities.CatalogProduct](productID)
	if err != nil {
		return err
	}
	if dbCatalogProduct.IsGlobalCatalog && !user.IsAdmin() {
		return errors.New("user can't edit global catalog item")
	}
	if shopIDFromProduct == nil || *shopIDFromProduct != user.ShopID {
		return fmt.Errorf("unable to delete product %s (err 1492)", productID)
	}

	dbCatalogProduct.DateUpdated = time.Now()
	dbCatalogProduct.Status = entities.EntityStatusDeleted

	// TODO: Log update user history record
	err = data.Update(&dbCatalogProduct)
	return err
}

func GetAllShopCatalogProducts(user *entities.User) ([]entities.CatalogProduct, error) {
	products, err := data.GetAllByShopID[entities.CatalogProduct](user.ShopID)
	return products, err
}

// PostNewJob insets a new calendarEvent in the DB and assigns a worker for it
func PostNewCatalogProduct(product *entities.CatalogProduct, user *entities.User) (*entities.CatalogProduct, error) {
	if product.IsGlobalCatalog && !user.IsAdmin() {
		return nil, errors.New("user can't create global catalog item")
	}
	if !user.IsAdmin() {
		product.ShopID = user.ShopID
	}

	product.ID = uuid.NewV4().String()
	product.DateCreated = time.Now()
	product.DateUpdated = time.Now()

	product.Status = entities.EntityStatusActive

	// Check cache
	err := data.InsertNew(product)
	return product, err
}
