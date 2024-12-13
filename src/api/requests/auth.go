package requests

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuth(r *http.Request) (*Auth, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var authRequest Auth
	err = json.Unmarshal(body, &authRequest)
	if err != nil {
		return nil, err
	}

	return &authRequest, nil
}
