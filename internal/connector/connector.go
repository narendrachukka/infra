package connector

// Connector is a standard interface that every
// connector plugin needs to implement
type Connector interface {
	// Creates a short lived credential in the upstream destination for a user
	CreateToken(id string, user string) (string, error)

	// Revokes a short lived credential in the upstream destination
	DeleteToken(id string)
}
