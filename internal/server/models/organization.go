package models

import (
	"regexp"

	"github.com/infrahq/infra/api"
	"github.com/infrahq/infra/internal/generate"
	"github.com/infrahq/infra/uid"
)

const DefaultOrganizationName = "Default"

type Organization struct {
	Model

	Name      string `gorm:"uniqueIndex:idx_organizations_name,where:deleted_at is NULL"`
	Domain    string `gorm:"uniqueIndex:idx_org_domain,where:deleted_at is NULL"`
	CreatedBy uid.ID
}

func (o *Organization) ToAPI() *api.Organization {
	return &api.Organization{
		ID:     o.ID,
		Name:   o.Name,
		Domain: o.Domain,
	}
}

var domainNameReplacer = regexp.MustCompile(`[^\da-zA-Z-]`)

// TODO: is the model the right place for this? Should we do this in CreateOrg ?
func (o *Organization) SetDefaultDomain() {
	if len(o.Domain) > 0 {
		return
	}
	o.Domain = domainNameReplacer.ReplaceAllStringFunc(o.Name, func(s string) string {
		if s == " " {
			return "-"
		}
		return ""
	}) + "-" + generate.MathRandom(5, generate.CharsetAlphaNumeric)
}

type OrganizationMember struct {
	// OrganizationID of the organization this entity belongs to.
	OrganizationID uid.ID
}

func (OrganizationMember) IsOrganizationMember() {}

func (o *OrganizationMember) SetOrganizationID(id uid.ID) {
	if o.OrganizationID == 0 {
		o.OrganizationID = id
	}
}
