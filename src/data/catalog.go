package data

import (
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
)

func GetAllProductsByShopID(shopID string) ([]entities.ShopProduct, error) {
	var entitiesArr []entities.ShopProduct
	err := pgConnection.Model(&entitiesArr).
		Column("shop_product.*").
		Relation("CatalogProduct").
		Where(`shop_product.shop_id = ? AND shop_product.status != ?`, shopID, entities.EntityStatusDeleted).
		Select()

	return entitiesArr, err
}

// GetShop retrieves the requeted shop from DB, if exists
func GetAllCatalogsFromProducts(products []entities.ShopProduct) ([]entities.CatalogProduct, error) {
	var catalogs []entities.CatalogProduct
	return catalogs, nil

	// ids := []string{}
	// for _, p := range products {
	// 	catalogID := p.CatalogProductID
	// 	ids = append(ids, "'"+catalogID+"'")
	// }
	// fmt.Println("IDS", strings.Join(ids, ","))

	// err := pgConnection.Model(&catalogs).Where("id in (?)", strings.Join(ids, ",")).Select()
	// return catalogs, err
}
