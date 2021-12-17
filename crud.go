package crud

import (
	"net/http"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"webimizer.dev/webimizer"
)

/* Global notAllowHandler to use in all CRUD requests if request is not accepted */
var GlobalNotAllowHandler webimizer.HttpNotAllowHandler

/* CRUDInterface interface for use in CRUD operations by calling Create, ReadOne, ReadAll, Update or Delete func */
type CRUDInterface interface {
	/* Create new data object in mongoDb collection */
	CreateOne(obj *interface{}, collection *mongo.Collection) error
	/* Find data row from mongoDb by uuid read and load to obj */
	ReadOne(uuid uuid.UUID, obj *interface{}, collection *mongo.Collection) error
	/* Fliter results using filter bson and load into objs array */
	ReadAll(filter interface{}, objs *[]interface{}, collection *mongo.Collection) error
	/* Update obj data to mongoDb by uuid */
	UpdateOne(uuid uuid.UUID, obj *interface{}, collection *mongo.Collection) error
	/* Update all data by uuids from given objs array */
	UpdateAll(uuid []uuid.UUID, objs *[]interface{}, collection *mongo.Collection) error
	/* Delete data from mongoDb by uuid */
	DeleteOne(uuid uuid.UUID, collection *mongo.Collection) error
	/* Delete all items from mongoDb by given uuids */
	DeleteAll(uuid []uuid.UUID, collection *mongo.Collection) error
}

/* Add CRUD operations handlers to http.ServeMux */
func AddCRUDHandlers(mux *http.ServeMux, permisions_map [][]string, crudInterface CRUDInterface) {

}
