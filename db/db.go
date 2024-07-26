package db

import (
	"context"

	"github.com/freightcms/api-template/models"
)

// DbContext provides an API for interacting with Entities. Implement this interface in a way
// for a Web API, Database, SOAP, WCF, etc.
type DbContext interface {
	// CreateEntity creates a new database record with the given obect and returns
	// the identifier of the record that was created. If there is/was an error within the applicatoin
	// creating the record from validation to database drivers the error is returned.
	CreateEntity(ctx context.Context, entity *models.EntityModel) (interface{}, error)
	// UpdateEntity in the database and returns the record.
	UpdateEntity(ctx context.Context, id interface{}, entity *models.EntityModel) error
	// DeleteEntity removes a record from the database. If there is an error deleting the record
	// it is returned. If the record does not exist, no error is thrown.
	DeleteEntity(ctx context.Context, id interface{}) error
	// FindEntity returns a single record from the database based on the identifier that
	// is used on the entity.
	FindEntity(ctx context.Context, id interface{}) (*models.EntityModel, error)
}
