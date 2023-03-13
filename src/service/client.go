package service

import (
	"fmt"
	"time"

	"bitbucket.org/marcoboschetti/visionmonk/src/data"
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	uuid "github.com/satori/go.uuid"
)

func GetClientByID(clientID string) (entities.Client, error) {
	client, err := data.GetByID[entities.Client](clientID)
	return client, err
}

func UpdateClient(inputClient *entities.Client, user *entities.User) error {
	// Retrieve client
	dbClient, err := data.GetByID[entities.Client](inputClient.ID)
	if err != nil {
		return err
	}
	if dbClient.ShopID != user.ShopID {
		return fmt.Errorf("unable to update client %s (err 1423)", inputClient.ID)
	}

	inputClient.DateCreated = dbClient.DateCreated
	inputClient.DateUpdated = time.Now()
	inputClient.ShopID = user.ShopID
	inputClient.Status = entities.EntityStatusActive

	// TODO: Log update user history record

	err = data.Update(inputClient)
	return err
}

func DeleteClient(clientID string, user *entities.User) error {

	// Retrieve client
	dbClient, err := data.GetByID[entities.Client](clientID)
	if err != nil {
		return err
	}
	if dbClient.ShopID != user.ShopID {
		return fmt.Errorf("unable to delete client %s (err 1492)", clientID)
	}

	dbClient.DateUpdated = time.Now()
	dbClient.Status = entities.EntityStatusDeleted

	// TODO: Log update user history record
	err = data.Update(&dbClient)
	return err
}

func GetAllShopClients(user *entities.User) ([]entities.Client, error) {
	clients, err := data.GetAllByShopID[entities.Client](user.ShopID)
	return clients, err
}

// PostNewJob insets a new calendarEvent in the DB and assigns a worker for it
func PostNewClient(client *entities.Client, user *entities.User) error {
	client.ID = uuid.NewV4().String()
	client.DateCreated = time.Now()
	client.DateUpdated = time.Now()

	client.ShopID = user.ShopID
	client.SubmittedByUserID = user.ID
	client.Status = entities.EntityStatusActive

	// Check cache
	err := data.InsertNew(client)
	return err
}
