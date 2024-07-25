package db

import "github.com/freightcms/api-template/models"

type DbContext interface {
	// CreateEntity creates a new database record with the given obect and returns
	// the identifier of the record that was created. If there is/was an error within the applicatoin
	// creating the record from validation to database drivers the error is returned.
	CreateEntity(entity *models.EntityModel) (interface{}, error)
	// UpdateEntity in the database and returns the record.
	UpdateEntity(id interface{}, entity *models.EntityModel) error
	// DeleteEntity removes a record from the database. If there is an error deleting the record
	// it is returned. If the record does not exist, no error is thrown.
	DeleteEntity(id interface{}) error
	// FindEntity returns a single record from the database based on the identifier that
	// is used on the entity.
	FindEntity(id interface{}) (*models.EntityModel, error)
}
