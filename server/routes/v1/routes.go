package v1

import (
	"github.com/kevinwmiller/digidoc/server/routes"
	"github.com/kevinwmiller/digidoc/server/routes/v1/auth"
	"github.com/kevinwmiller/digidoc/server/routes/v1/storage"
)

// Router is the router for v1 of the Digidoc API
type Router struct{}

// List returns a list of route handlers
func (r Router) List() routes.Routes {
	routeHandlers := auth.Router{}.List()
	routeHandlers = append(routeHandler, storage.Router{}.List())
	return routeHandlers
}
