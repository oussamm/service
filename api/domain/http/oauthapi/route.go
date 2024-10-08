package oauthapi

import (
	"net/http"

	"github.com/ardanlabs/service/app/sdk/auth"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// Config contains all the configuration for the auth app.
type Config struct {
	Auth            *auth.Auth
	Log             *logger.Logger
	TokenKey        string
	GoogleKey       string
	GoogleSecret    string
	Callback        string
	StoreKey        string
	UIAdminRedirect string
	UILoginRedirect string
}

// Routes adds the routes for the auth app.
func Routes(app *web.App, cfg Config) {
	api := newAPI(cfg)

	app.RawHandlerFunc(http.MethodGet, "", "/api/auth/{provider}", api.authenticate)
	app.RawHandlerFunc(http.MethodGet, "", "/api/logout/{provider}", api.logout)
	app.RawHandlerFunc(http.MethodGet, "", "/api/auth/{provider}/callback", api.authCallback)
}
