package requests

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type UserID struct {
	UserID int `json:"user_id"`
}

func NewUserID(r *http.Request) (*UserID, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var user UserID
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
