package data

import "bitbucket.org/marcoboschetti/visionmonk/src/entities"

// InsertNew Adds a entity in the repository
func InsertNew[T any](entity *T) error {
	err := pgConnection.Insert(entity)
	return err
}

// Update updates a entity in the repository
func Update[T any](entity *T) error {
	err := pgConnection.Update(entity)
	return err
}

// Get retrieves the requeted entity from DB, if exists
func GetAll[T any]() ([]T, error) {
	var entities []T
	err := pgConnection.Model(&entities).Select()
	return entities, err
}

func GetAllByShopID[T any](shopID string) ([]T, error) {
	var entitiesArr []T
	err := pgConnection.Model(&entitiesArr).
		Where("shop_id = ? AND status != ?", shopID, entities.EntityStatusDeleted).
		Select()

	return entitiesArr, err
}

// GetByID retrieves the requeted entity from DB, if exists
func GetByID[T any](entityID string) (T, error) {
	var entity T
	err := pgConnection.Model(&entity).
		Where("id = ?", entityID).
		Select()
	return entity, err
}
