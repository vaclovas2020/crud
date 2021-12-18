package crud

import (
	"net/http"

	"webimizer.dev/webimizer"
)

/* CRUD action constant: Create new item */
const CrudCreateOne string = "CreateOne"

/* CRUD action constant: Create new item */
const CrudCreateAll string = "CreateAll"

/* CRUD action constant:  Find and return one item */
const CrudReadOne string = "ReadOne"

/* CRUD action constant: Filter and return all items */
const CrudReadAll string = "ReadAll"

/* CRUD action constant: Update one item */
const CrudUpdateOne string = "UpdateOne"

/*CRUD action constant: Update all items */
const CrudUpdateAll string = "UpdateAll"

/* CRUD action constant: Delete one item */
const CrudDeleteOne string = "DeleteOne"

/*CRUD action constant: Delete all items */
const CrudDeleteAll string = "DeleteAll"

/* CRUD UserRole Admin */
const UserRoleAdmin UserRole = "Admin"

/* CRUD UserRole User */
const UserRoleUser UserRole = "User"

/* Global notAllowHandler to use in all CRUD requests if request is not accepted (required) */
var GlobalNotAllowHandler webimizer.HttpNotAllowHandler

/* Global allowOrigins to use in all CRUD request (optional) */
var GlobalAllowedOrigins []string

/* UserRole string type, using to give role to user with specific permissions.
There is two constants you can use UserRoleAdmin and UserRoleUser. You can create your own UserRole too. */
type UserRole string

/* Get UserRole name (string) */
func (ur UserRole) String() string {
	return string(ur)
}

/* Check if user has specific permission by given permisions_map */
func (ur UserRole) UserCan(permissions_map map[string][]string, permission string) bool {
	for role, permissions := range permissions_map {
		if role == ur.String() {
			for _, perm := range permissions {
				if perm == permission {
					return true
				}
			}
		}
	}
	return false
}

/* UserUUID string type, to define autheticated User UUID string */
type UserUUID string

/* Get UserUUID string value */
func (uuid UserUUID) String() string {
	return string(uuid)
}

/* HttpHandler for user authentication processing and return UserRole if success, otherwise return error */
type AuthHandler func(rw http.ResponseWriter, r *http.Request) (UserRole, UserUUID, error)

/* ErrorHandler to handle error messages */
type ErrorHandler func(err error)

/* CrudInterface interface for use in CRUD operations by calling Create, ReadOne, ReadAll, Update or Delete func */
type CrudInterface interface {
	/* Create new item */
	CreateOne(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
	/* Create new items */
	CreateAll(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
	/* Find and return one item */
	ReadOne(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
	/* Filter and return all items */
	ReadAll(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
	/* Update one item */
	UpdateOne(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
	/* Update all items */
	UpdateAll(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
	/* Delete one item */
	DeleteOne(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
	/* Delete all items */
	DeleteAll(rw http.ResponseWriter, r *http.Request, ur *UserRole, uuid *UserUUID) error
}

/* Add CRUD operations handlers to mux *http.ServeMux */
func AddCrudHandlers(mux *http.ServeMux, one_slug string, all_slug string,
	permissions_map map[string][]string, crudInterface CrudInterface,
	authHandler AuthHandler, errorHandler ErrorHandler) {
	mux.HandleFunc(one_slug, webimizer.HttpHandlerStruct{
		NotAllowHandler: GlobalNotAllowHandler,
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			userRole, userUUID, err := authHandler(rw, r)
			if err != nil {
				errorHandler(err)
				return
			}
			webimizer.Post(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudCreateOne) {
					if crudInterface.CreateOne(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
			webimizer.Get(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudReadOne) {
					if crudInterface.ReadOne(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
			webimizer.Put(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudUpdateOne) {
					if crudInterface.UpdateOne(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
			webimizer.Delete(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudDeleteOne) {
					if crudInterface.DeleteOne(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
		}),
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}.Build())
	mux.HandleFunc(all_slug, webimizer.HttpHandlerStruct{
		NotAllowHandler: GlobalNotAllowHandler,
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			userRole, userUUID, err := authHandler(rw, r)
			if err != nil {
				errorHandler(err)
				return
			}
			webimizer.Post(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudCreateAll) {
					if crudInterface.CreateAll(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
			webimizer.Get(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudReadAll) {
					if crudInterface.ReadAll(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
			webimizer.Put(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudUpdateAll) {
					if crudInterface.UpdateAll(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
			webimizer.Delete(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				if userRole.UserCan(permissions_map, CrudDeleteAll) {
					if crudInterface.DeleteAll(rw, r, &userRole, &userUUID) != nil {
						errorHandler(err)
					}
				}
			})
		}),
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}.Build())
}
