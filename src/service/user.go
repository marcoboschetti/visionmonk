package service

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"bitbucket.org/marcoboschetti/visionmonk/src/data"
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	"bitbucket.org/marcoboschetti/visionmonk/src/utils"
)

// PostNewUser insets an empty user in the DB
func PostNewUser(password, firstname, lastname, email, ip, shopToken string) (*entities.User, error) {

	// For now, check default user creation
	if len(password) == 0 || len(email) == 0 || len(firstname) == 0 || len(lastname) == 0 {
		return nil, errors.New("missing required fields")
	}

	// Check that email not exists
	userAlreadyExists, err := data.ExistsUser(email)
	if err != nil {
		return nil, err
	}
	if userAlreadyExists {
		return nil, errors.New("email info already exists")
	}

	// Validate shop token
	shop, err := data.GetShopByToken(shopToken)
	if err != nil {
		return nil, errors.New("invalid shop token")
	}

	passwordHash := utils.GetPasswordHash(password)
	u := entities.User{
		ID:           uuid.NewV4().String(),
		PasswordHash: passwordHash,
		DateCreated:  time.Now(),
		FirstName:    firstname,
		LastName:     lastname,
		Email:        email,
		IP:           ip,
		ShopID:       shop.ID,
		SecretToken:  uuid.NewV4().String(),
	}

	err = data.InsertNewUser(&u)
	return &u, err
}

// GetUserByToken returns an user if exists
func GetUserByToken(userToken string) (*entities.User, error) {
	user, err := data.GetUserByToken(userToken)
	if err != nil {
		return nil, err
	}

	shop, err := data.GetShopByID(user.ShopID)
	user.Shop = shop

	return &user, err
}

// GetUserByToken returns an user if exists
func GetUserByCredentials(email, password string) (*entities.User, error) {
	user, err := data.GetUserByCredentials(email, password)
	if err != nil {
		return nil, err
	}
	return &user, err
}
