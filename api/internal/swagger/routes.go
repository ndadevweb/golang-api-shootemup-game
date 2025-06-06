//go:build dev
// +build dev

package swagger

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Add Swagger, only using with dev mode
func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
