// Package all binds all the routes into the specified app.
package all

import (
	"github.com/razvan286/software-design-go-and-kubernetes/api/http/api/mux"
	"github.com/razvan286/software-design-go-and-kubernetes/api/http/domain/authapi"
	"github.com/razvan286/software-design-go-and-kubernetes/api/http/domain/checkapi"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	authapi.Routes(app, authapi.Config{
		Auth: cfg.Auth,
	})
}
