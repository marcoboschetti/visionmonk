package service

import (
	"errors"
	"fmt"
	"time"

	"bitbucket.org/marcoboschetti/visionmonk/src/data"
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	uuid "github.com/satori/go.uuid"
)

func GetShopProductByID(productID string) (*entities.ShopProduct, error) {
	product, err := data.GetByID[entities.ShopProduct](productID)
	if err != nil {
		return nil, err
	}
	catalogProduct, err := data.GetByID[entities.CatalogProduct](product.CatalogProductID)
	if err != nil {
		return nil, err
	}

	product.CatalogProduct = &catalogProduct
	return &product, err
}

func UpdateShopProduct(inputShopProduct *entities.ShopProduct, user *entities.User) error {

	dbShopProduct, err := data.GetByID[entities.ShopProduct](inputShopProduct.ID)
	if err != nil {
		return err
	}
	if dbShopProduct.ShopID != user.ShopID {
		return errors.New("unauthorized by product shop_id")
	}
	inputShopProduct.ShopID = dbShopProduct.ID

	if inputShopProduct.CatalogProductID != "" {
		return errors.New("catalog change for existing product not allowed")
	}
	inputShopProduct.CatalogProductID = dbShopProduct.CatalogProductID

	if inputShopProduct.CatalogProduct == nil {
		return errors.New("needs to provide full catalog to update product")
	}

	// Update catalog
	existingCatalog, err := GetCatalogProductByID(dbShopProduct.CatalogProductID)
	if err != nil {
		return err
	}
	inputShopProduct.CatalogProduct.IsGlobalCatalog = existingCatalog.IsGlobalCatalog
	inputShopProduct.CatalogProduct.ShopID = existingCatalog.ShopID
	inputShopProduct.CatalogProduct.ID = existingCatalog.ID

	// Check if the catalog was updated, and if enabled do so
	if !existingCatalog.Equals(*inputShopProduct.CatalogProduct) {
		if existingCatalog.IsGlobalCatalog && !user.IsAdmin() {
			return errors.New("user can't edit global catalog item")
		}
		err = UpdateCatalogProduct(inputShopProduct.CatalogProduct, user, &inputShopProduct.ShopID)
		if err != nil {
			return err
		}
	}

	inputShopProduct.ID = dbShopProduct.ID
	inputShopProduct.DateCreated = dbShopProduct.DateCreated
	inputShopProduct.ShopID = dbShopProduct.ShopID
	inputShopProduct.SubmittedByUserID = dbShopProduct.SubmittedByUserID
	inputShopProduct.Status = dbShopProduct.Status
	inputShopProduct.DateUpdated = time.Now()

	// TODO: Log update user history record
	err = data.Update(inputShopProduct)
	return err
}

func DeleteShopProduct(productID string, user *entities.User) error {
	// Retrieve product
	dbShopProduct, err := data.GetByID[entities.ShopProduct](productID)
	if err != nil {
		return err
	}
	if dbShopProduct.ShopID != user.ShopID {
		return fmt.Errorf("unable to delete product %s (err 1492)", productID)
	}

	existingCatalog, err := GetCatalogProductByID(dbShopProduct.CatalogProductID)
	if err != nil {
		return err
	}
	// If authorized, soft delete catalog too
	if existingCatalog.ShopID == user.ShopID {
		err = DeleteCatalogProduct(dbShopProduct.CatalogProductID, user, &dbShopProduct.ShopID)
		if err != nil {
			return err
		}
	}

	dbShopProduct.DateUpdated = time.Now()
	dbShopProduct.Status = entities.EntityStatusDeleted

	// TODO: Log update user history record
	err = data.Update(&dbShopProduct)
	if err == nil {
		err = data.Update(&existingCatalog)
	}

	return err
}

func GetAllShopShopProducts(user *entities.User) ([]entities.ShopProduct, error) {
	products, err := data.GetAllProductsByShopID(user.ShopID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// PostNewJob insets a new calendarEvent in the DB and assigns a worker for it
func PostNewShopProduct(product entities.ShopProduct, user *entities.User) error {
	var catalog entities.CatalogProduct

	if product.CatalogProduct != nil {
		// Trying to create new catalog item
		newCatalog, err := PostNewCatalogProduct(product.CatalogProduct, user)
		if err != nil {
			return err
		}

		catalog = *newCatalog
	} else {
		existingCatalog, err := GetCatalogProductByID(product.CatalogProductID)
		if err != nil {
			return err
		}
		if !existingCatalog.IsGlobalCatalog && existingCatalog.ShopID != user.ShopID {
			return fmt.Errorf("unauthorized to link non-global catalog %s", product.CatalogProductID)
		}
	}

	product.ID = uuid.NewV4().String()
	product.CatalogProductID = catalog.ID
	product.DateCreated = time.Now()
	product.DateUpdated = time.Now()

	product.ShopID = user.ShopID
	product.SubmittedByUserID = user.ID
	product.Status = entities.EntityStatusActive

	// Check cache
	err := data.InsertNew(&product)
	return err
}
