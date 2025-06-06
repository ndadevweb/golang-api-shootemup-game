package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoute() http.Handler {
	r := chi.NewRouter()

	// @Summary Ping
	// @Description Route de test
	// @Success 200 {object} map[string]string
	// @Router /ping [get]
	r.Get("/", WelcomeHandler)

	// Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	// r.Get("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs"))).ServeHTTP)

	return r
}

// @Summary Accueil
// @Description Affiche un message de bienvenue
// @Tags Accueil
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Bienvenue sur l'API Gowebgame"}`))
}
