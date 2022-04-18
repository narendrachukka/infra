package models

import (
	"github.com/infrahq/infra/api"
)

type Destination struct {
	Model

	Name string `gorm:"uniqueIndex:,where:deleted_at is NULL" validate:"required"`

	CA    string
	URL   string
	Token EncryptedAtRest
}

func (d *Destination) ToAPI() *api.Destination {
	return &api.Destination{
		ID:      d.ID,
		Created: api.Time(d.CreatedAt),
		Updated: api.Time(d.UpdatedAt),
		Name:    d.Name,
		Connection: api.DestinationConnection{
			URL: d.URL,
			CA:  d.CA,
		},
	}
}
