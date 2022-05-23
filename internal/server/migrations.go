package server

import (
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/infrahq/infra/api"
	"github.com/infrahq/infra/uid"
)

// TODO: enforce that the slice of versions is sorted before building routes
func (a *API) populateVersionHandlers() {
	a.versions = map[routeKey][]routeVersion{
		routeKey{http.MethodGet, "/api/grants"}: {
			version("0.12.2", listGrantsV0_12_2),
		},
	}
}

func version[Req, Res any](v string, handler HandlerFunc[Req, Res]) routeVersion {
	return routeVersion{
		version: semver.MustParse(v),
		handler: wrapHandler(handler),
	}
}

type identityGrant struct {
	ID uid.ID `json:"id"`

	Created   api.Time `json:"created"`
	CreatedBy uid.ID   `json:"created_by"`
	Updated   api.Time `json:"updated"`

	Subject   uid.PolymorphicID `json:"subject,omitempty"`
	Privilege string            `json:"privilege"`
	Resource  string            `json:"resource"`
}

func migrateUserGrantToIdentity(grant api.Grant) identityGrant {
	var sub uid.PolymorphicID

	if grant.User != 0 {
		sub = uid.NewIdentityPolymorphicID(grant.User)
	} else {
		sub = uid.NewGroupPolymorphicID(grant.Group)
	}

	return identityGrant{
		ID:        grant.ID,
		Created:   grant.Created,
		CreatedBy: grant.CreatedBy,
		Updated:   grant.Updated,
		Subject:   sub,
		Privilege: grant.Privilege,
		Resource:  grant.Resource,
	}
}

type loginResponseV0_12_2 struct {
	PolymorphicID          uid.PolymorphicID `json:"polymorphicID"`
	Name                   string            `json:"name"`
	AccessKey              string            `json:"accessKey"`
	PasswordUpdateRequired bool              `json:"passwordUpdateRequired,omitempty"`
	Expires                api.Time          `json:"expires"`
}
