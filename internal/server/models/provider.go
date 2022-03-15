package models

import (
	"time"

	"github.com/infrahq/infra/api"
	"github.com/infrahq/infra/uid"
)

type Provider struct {
	Model

	Name         string `gorm:"uniqueIndex:,where:deleted_at is NULL" validate:"required"`
	URL          string `validate:"required"`
	ClientID     string
	ClientSecret EncryptedAtRest

	Users []User
}

func (p *Provider) ToAPI() *api.Provider {
	return &api.Provider{
		Name:    p.Name,
		ID:      p.ID,
		Created: p.CreatedAt.Unix(),
		Updated: p.UpdatedAt.Unix(),

		URL:      p.URL,
		ClientID: p.ClientID,
	}
}

// ProviderToken tracks the access and refresh tokens from an identity provider associated with a user
type ProviderToken struct {
	Model

	UserID      uid.ID
	ProviderID  uid.ID
	RedirectURL string `validate:"required"` // needs to match the redirect URL specified when the token was issued for refreshing

	AccessToken  EncryptedAtRest
	RefreshToken EncryptedAtRest
	ExpiresAt    time.Time
}
