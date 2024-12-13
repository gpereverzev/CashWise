package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gpereverzev/CashWise/src/api/requests"
	"github.com/gpereverzev/CashWise/src/data"
	"net/http"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUserID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.UserID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	masterDB := r.Context().Value(MasterDBContextKey).(data.MasterDB)
	user, err := masterDB.Users().Get(req.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve user information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
