package entities

type (
	EntityStatus string
)

const (
	EntityStatusActive  EntityStatus = "active"
	EntityStatusDeleted EntityStatus = "deleted"
)
