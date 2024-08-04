package mongodb

import (
	"context"

	"github.com/freightcms/webservice-template/db"
	"github.com/freightcms/webservice-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type appDbContext struct {
	session        mongo.Session
	databaseName   string
	collectionName string
}

// CreateEntity implements db.DbContext.
func (a *appDbContext) CreateEntity(ctx context.Context, entity *models.EntityModel) (interface{}, error) {
	if err := a.session.StartTransaction(); err != nil {
		return nil, err
	}
	insert, err := a.collection().InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	if err := a.session.CommitTransaction(ctx); err != nil {
		return nil, err
	}

	return insert.InsertedID, nil
}

// DeleteEntity implements db.DbContext.
func (a *appDbContext) DeleteEntity(ctx context.Context, id interface{}) error {
	if err := a.session.StartTransaction(); err != nil {
		return err
	}
	if _, err := a.collection().DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return err
	}
	if err := a.session.CommitTransaction(ctx); err != nil {
		return err
	}
	return nil
}

// FindEntity implements db.DbContext.
func (a *appDbContext) FindEntity(ctx context.Context, id interface{}) (*models.EntityModel, error) {
	entity := &models.EntityModel{}
	record := a.collection().FindOne(ctx, bson.M{"_id": id})
	return entity, record.Decode(&entity)
}

// UpdateEntity implements db.DbContext.
func (a *appDbContext) UpdateEntity(ctx context.Context, id interface{}, entity *models.EntityModel) error {
	if err := a.session.StartTransaction(); err != nil {
		return err
	}

	a.collection().FindOneAndUpdate(ctx, &bson.M{"_id": id}, &entity)
	if err := a.session.CommitTransaction(ctx); err != nil {
		return err
	}

	return nil
}

func (a *appDbContext) collection() *mongo.Collection {
	return a.session.Client().Database(a.databaseName).Collection(a.collectionName)
}

// CreateDbContext creates a new database object to create, update, fetch, delete, query
// entities within a database. the database parameter should be the name of the mongodb
// database name to connect to when calling `Client().Database()`. Collectin parameter should
// be the collection in the mongodb database that is used to store information. Used when
// calling Client().Database(dbname).Collection(collectionName)
func CreateDbContext(database, collection string, session mongo.Session) db.DbContext {
	return &appDbContext{
		databaseName:   database,
		collectionName: collection,
		session:        session,
	}
}
