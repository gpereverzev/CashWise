package handlers

import (
	"encoding/json"
	"github.com/gpereverzev/CashWise/src/api/requests"
	"github.com/gpereverzev/CashWise/src/data"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAuth(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	masterDB := r.Context().Value(MasterDBContextKey).(data.MasterDB)
	user, err := masterDB.Users().FindByEmail(req.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User authenticated successfully",
		"userID":  string(user.UserID),
	})
}
