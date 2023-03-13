package service

import (
	"bitbucket.org/marcoboschetti/visionmonk/src/data"
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
)

// GetShopByID returns a shop if exists
func GetShopByID(shopID string) (*entities.Shop, error) {
	shop, err := data.GetShopByID(shopID)
	if err != nil {
		return nil, err
	}

	return shop, err
}
