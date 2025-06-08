package signin

import (
	"encoding/json"
	"gowebgame/internal/db"
	"gowebgame/internal/jwt"
	"gowebgame/internal/repository"
	"gowebgame/internal/utils"
	"net/http"
)

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Token string `json:"token"`
}

// SigninHandler godoc
// @Summary Connexion utilisateur
// @Description Authentifie un utilisateur existant et renvoie un token JWT
// @Tags Authentification
// @Accept json
// @Produce json
// @Param request body SigninRequest true "Identifiants de l'utilisateur"
// @Success 200 {object} SigninResponse "JWT généré"
// @Failure 400 {object} map[string]string "Requête invalide"
// @Failure 401 {object} map[string]string "Identifiants incorrects"
// @Router /auth/signin [post]
func SigninHandler() http.HandlerFunc {
	sqlite := db.GetDB()
	userRepo := repository.NewUserRepository(sqlite)

	return func(w http.ResponseWriter, r *http.Request) {
		var req SigninRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		user, err := userRepo.FindByUsername(req.Username)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		if !utils.CheckPasswordHash(req.Password, user.Password) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		token, err := jwt.GenerateJWT(user.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SigninResponse{Token: token})
	}
}
