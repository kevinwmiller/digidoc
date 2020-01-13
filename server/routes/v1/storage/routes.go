package storage

import (
	"github.com/kevinwmiller/digidoc/server/routes"
)

// Router manages routes related to authentication of user objects
type Router struct{}

// List returns a list of authentication route handlers
func (r Router) List() routes.Routes {
	return routes.Routes{
		"/v1/storage/createbucket": routes.Methods{
			routes.Post: createBucket,
		},
		"/v1/storage/{bucketName}": routes.Methods{
			routes.Delete: deleteBucket,
			routes.Get:    listObjects,
		},
		"/v1/storage/{bucketName}/objects": routes.Methods{
			routes.Post:   createObject,
			routes.Get:    downloadObject,
			routes.Delete: deleteObject,
		},
	}
}
