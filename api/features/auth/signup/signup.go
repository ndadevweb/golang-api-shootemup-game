package signup

import (
	"encoding/json"
	"gowebgame/internal/models"
	"gowebgame/internal/db"
	"gowebgame/internal/repository"
	"gowebgame/internal/utils"
	"net/http"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignupHandler godoc
// @Summary Inscription utilisateur
// @Description Crée un nouveau compte utilisateur avec un identifiant unique et un mot de passe
// @Tags Authentification
// @Accept json
// @Produce json
// @Param request body signup.SignupRequest true "Informations de l'utilisateur"
// @Success 201 {object} map[string]string "Message de confirmation"
// @Failure 400 {object} map[string]string "Erreur de validation"
// @Failure 500 {object} map[string]string "Erreur serveur ou utilisateur déjà existant"
// @Router /auth/signup [post]
func SignupHandler() http.HandlerFunc {
	sqlite := db.GetDB()
	userRepo := repository.NewUserRepository(sqlite)

	return func(w http.ResponseWriter, r *http.Request) {
		var req SignupRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if len(req.Username) < 3 || len(req.Password) < 6 {
			http.Error(w, "Invalid username or password too short", http.StatusUnprocessableEntity)
			return
		}

		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		user := &models.User{
			Username: req.Username,
			Password: hashedPassword,
		}

		err = userRepo.CreateUser(user)
		if err != nil {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User successfully created",
		})
	}
}
