package data

import (
	"sync"

	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
	"bitbucket.org/marcoboschetti/visionmonk/src/utils"
)

// TODO: Implement a real cache for users
var (
	usersCache       = map[string]entities.User{}
	orderedUsersList = []string{}
	mutex            = &sync.RWMutex{}
)

// InsertNewUser Adds a user in the repository
func InsertNewUser(user *entities.User) error {
	err := pgConnection.Insert(user)
	return err
}

// UpdateUser updates a user in the repository
func UpdateUser(user *entities.User) error {
	err := pgConnection.Update(user)
	return err
}

// GetUserByCredentials retrieves the requeted user from DB, if exists
func GetUserByCredentials(email, password string) (entities.User, error) {
	passwordHash := utils.GetPasswordHash(password)
	var user entities.User

	err := pgConnection.Model(&user).
		Where("email = ? AND password_hash = ?", email, passwordHash).
		Select()

	return user, err
}

// GetUserByID retrieves the requeted user from DB, if exists
func GetUserByToken(userToken string) (entities.User, error) {

	mutex.RLock()
	if _, ok := usersCache[userToken]; ok {
		user := usersCache[userToken]
		mutex.RUnlock()
		return user, nil
	}

	mutex.RUnlock()
	mutex.Lock()
	defer mutex.Unlock()

	var user entities.User
	err := pgConnection.Model(&user).
		Where("secret_token = ?", userToken).
		Select()

	if err == nil {
		// Add to cache
		usersCache[userToken] = user
		orderedUsersList = append(orderedUsersList)

		// Evict first element from cache
		if len(orderedUsersList) > 100 {
			delete(usersCache, orderedUsersList[0])
			orderedUsersList = orderedUsersList[1:]
		}
	}

	return user, err
}

// ExistsUser returns if any user with the given email or username exists
func ExistsUser(email string) (bool, error) {
	var user *entities.User
	exists, err := pgConnection.Model(user).
		Where("email = ?", email).
		Exists()
	return exists, err
}
