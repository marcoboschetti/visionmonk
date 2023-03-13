package entities

import (
	"time"

	"bitbucket.org/marcoboschetti/visionmonk/src/utils"
)

type User struct {
	ID          string       `json:"-"`
	DateCreated time.Time    `json:"date_created"`
	DateUpdated time.Time    `json:"date_updated"`
	ShopID      string       `json:"shop_id"`
	Status      EntityStatus `json:"status"`

	Email        string `json:"email"`
	PasswordHash string `json:"-"`

	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	IP          string   `json:"-"`
	Roles       []string `json:"-"`
	Validated   bool     `json:"-"`
	SecretToken string   `json:"-"`

	// Not in DB
	Shop *Shop `json:"shop" sql:"-"`
}

func (user *User) IsAdmin() bool {
	isAdmin, _ := utils.Contains(user.Roles, UserRoleGlobalAdmin)
	return isAdmin
}

const (
	UserRoleGlobalAdmin string = "global_admin"
)
