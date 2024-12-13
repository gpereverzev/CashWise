package handlers

import (
	"encoding/json"
	"github.com/gpereverzev/CashWise/src/api/requests"
	"github.com/gpereverzev/CashWise/src/data"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	log.Println("Received registration request")

	req, err := requests.NewRegistration(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := data.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Role:         "user",
		PasswordHash: string(hashedPassword),
	}

	masterDB := r.Context().Value(MasterDBContextKey).(data.MasterDB)
	err = masterDB.Users().Insert(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			w.WriteHeader(http.StatusConflict)
			return
		}

		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
		"userID":  string(user.UserID),
	})
}
