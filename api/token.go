package api

import "github.com/infrahq/infra/uid"

type CreateTokenRequest struct {
	Destination uid.ID `json:"destination" validate:"required"`
}

type CreateTokenResponse struct {
	Expires Time   `json:"expires"`
	Token   string `json:"token"`
}
