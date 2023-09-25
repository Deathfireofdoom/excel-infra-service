// api/api.go
package api

import (
	"github.com/Deathfireofdoom/excel-infra-service/api/health"
	"github.com/Deathfireofdoom/excel-infra-service/api/workbook"
	"github.com/go-chi/chi"
)

// NewRouter initializes all the api routes and returns the router.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/health", health.NewRouter())
	r.Mount("/workbook", workbook.NewRouter())
	return r
}
