package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	Middleware "gowebgame/internal/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	AuthSignup "gowebgame/features/auth/signup"
	AuthSignin "gowebgame/features/auth/signin"
)

func InitRoute() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(api chi.Router) {
		api.Get("/", WelcomeHandler)
		api.With(Middleware.JWTMiddleware).Get("/private", PrivateHandler)
	
		api.Post("/auth/signup", AuthSignup.SignupHandler())
		api.Post("/auth/signin", AuthSignin.SigninHandler())
	})
	

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
// @Router /api [get]
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Bienvenue sur l'API Gowebgame"}`))
}

// @Summary Page privee
// @Description Affiche un message indiquant que l'on est sur une page privee
// @Tags Private
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/private [get]
func PrivateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Bienvenue sur une page privée"}`))
}
