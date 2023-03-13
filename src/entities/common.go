package entities

import "time"

type Shop struct {
	ID          string    `json:"id"`
	DateCreated time.Time `json:"date_created"`
	OwnerIDs    []string  `json:"owner_user_ids"`
	SecretToken string    `json:"secret_token"`

	Email  string `json:"email"`
	Tier   string `json:"tier"`
	Status string `json:"status"`

	Name          string `json:"name"`
	Addess        string `json:"addess"`
	Locality      string `json:"locality"`
	Neighborhood  string `json:"neighborhood"`
	ZipCode       string `json:"zipcode"`
	AditionalInfo string `json:"aditionalinfo"`

	// Not in DB
	// Counts
	UsersCount   int `json:"users_count" sql:"-"`
	ClientsCount int `json:"clients_count" sql:"-"`
}

type entityMetadata struct {
	ID          string       `json:"id"`
	Status      EntityStatus `json:"status"`
	DateCreated time.Time    `json:"date_created"`
	DateUpdated time.Time    `json:"date_updated"`
}

type shopEntityMetadata struct {
	entityMetadata
	ShopID            string `json:"shop_id"`
	SubmittedByUserID string `json:"submitted_by_user_id"`
}
