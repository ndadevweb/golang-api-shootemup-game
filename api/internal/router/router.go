package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	AuthSignup "gowebgame/features/auth/signup"
	AuthSignin "gowebgame/features/auth/signin"
)

func InitRoute() http.Handler {
	r := chi.NewRouter()

	r.Get("/", WelcomeHandler)

	r.Post("/auth/signup", AuthSignup.SignupHandler())

	r.Post("/auth/signin", AuthSignin.SigninHandler())


	// Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

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
