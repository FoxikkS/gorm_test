package userApi

import (
	"encoding/json"
	"gorm_test_proj/models"
	"log/slog"
	"net/http"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		slog.Error(err.Error())
	}

	err = models.Validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error(err.Error())
	}

	w.WriteHeader(http.StatusCreated)

}
