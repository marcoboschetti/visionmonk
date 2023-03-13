package data

import (
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
)

// InsertNewShop Adds a shop in the repository
func InsertNewShop(shop *entities.Shop) error {
	err := pgConnection.Insert(shop)
	return err
}

// UpdateShop updates a shop in the repository
func UpdateShop(shop *entities.Shop) error {
	err := pgConnection.Update(shop)
	return err
}

// GetShop retrieves the requeted shop from DB, if exists
func GetAllShops() ([]entities.Shop, error) {
	var shops []entities.Shop

	err := pgConnection.Model(&shops).Select()

	return shops, err
}

// GetShopByID retrieves the requeted shop from DB, if exists
func GetShopByID(shopID string) (*entities.Shop, error) {
	var shop entities.Shop
	err := pgConnection.Model(&shop).
		Where("id = ?", shopID).
		Select()
	if err != nil {
		return nil, err
	}

	usersCount, err := pgConnection.Model(&[]entities.User{}).Where("shop_id = ?", shopID).SelectAndCount()
	if err != nil {
		return nil, err
	}

	clientsCount, err := pgConnection.Model(&[]entities.Client{}).Where("shop_id = ? AND status != ?", shopID, entities.EntityStatusDeleted).SelectAndCount()
	if err != nil {
		return nil, err
	}

	shop.UsersCount = usersCount
	shop.ClientsCount = clientsCount

	return &shop, err
}

// GetShopByID retrieves the requeted shop from DB, if exists
func GetShopByToken(token string) (entities.Shop, error) {
	var shop entities.Shop
	err := pgConnection.Model(&shop).
		Where("secret_token = ?", token).
		Select()
	return shop, err
}

// ExistsShop returns if any shop with the given email or shopname exists
func ExistsShop(email string) (bool, error) {
	var shop *entities.Shop
	exists, err := pgConnection.Model(shop).
		Where("email = ?", email).
		Exists()
	return exists, err
}
