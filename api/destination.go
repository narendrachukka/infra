package api

import (
	"github.com/infrahq/infra/uid"
)

type DestinationConnection struct {
	URL string `json:"url" validate:"required" example:"aa60eexample.us-west-2.elb.amazonaws.com"`
	CA  string `json:"ca" example:"-----BEGIN CERTIFICATE-----\nMIIDNTCCAh2gAwIBAgIRALRetnpcTo9O3V2fAK3ix+c\n-----END CERTIFICATE-----\n"`
}

type Destination struct {
	ID      uid.ID `json:"id"`
	Name    string `json:"name" form:"name"`
	Created Time   `json:"created"`
	Updated Time   `json:"updated"`

	Connection DestinationConnection `json:"connection"`
}

type ListDestinationsRequest struct {
	Name     string `form:"name"`
	UniqueID string `form:"unique_id"`
}

type CreateDestinationRequest struct {
	Name       string                `json:"name" validate:"required"`
	Connection DestinationConnection `json:"connection"`
	Token      string                `json:"token"`
}

type UpdateDestinationRequest struct {
	ID         uid.ID                `uri:"id" json:"-" validate:"required"`
	Name       string                `json:"name" validate:"required"`
	UniqueID   string                `json:"uniqueID"`
	Connection DestinationConnection `json:"connection"`
	Token      string                `json:"token"`
}
