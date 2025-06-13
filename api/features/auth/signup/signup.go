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
	Username string `json:"login"`
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
// @Router /api/auth/signup [post]
func SignupHandler() http.HandlerFunc {
	sqlite := db.GetDB()
	userRepo := repository.NewUserRepository(sqlite)

	return func(w http.ResponseWriter, r *http.Request) {
		var req SignupRequest

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid request",
			})
			return
		}

		if len(req.Username) < 2  {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid username (2 characters or more required)",
			})
			return
		}

		if len(req.Password) < 8  {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid password (8 characters or more required)",
			})
			return
		}

		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		user := &models.User{
			Username: req.Username,
			Password: hashedPassword,
		}

		err = userRepo.CreateUser(user)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "This username is already used",
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User successfully created",
		})
	}
}
