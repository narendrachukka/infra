package cmd

import (
	"fmt"
	"net/http"

	"github.com/infrahq/infra/api"
)

// UserFacingError wraps an error to provide better error messages to users.
// Any non-UserFacingError will be printed as "unhandled errors".
type UserFacingError struct {
	Underlying error
	Message    string
}

func (u UserFacingError) Error() string {
	if apiError, ok := u.Underlying.(api.Error); ok {
		return formatAPIError(apiError, u.Message)
	}

	// TODO: format the error better
	return fmt.Sprintf("%v: %v", u.Message, u.Underlying)
}

func formatAPIError(apiError api.Error, message string) string {
	switch apiError.Code {
	case http.StatusBadRequest:
		return fmt.Sprintf("%v: bad request: %v", message, apiError.Message)
	case http.StatusBadGateway:
		// this error should be displayed to the user so they can see its an external problem
		return fmt.Sprintf("%v: bad gateway: %v", message, apiError.Message)
	case http.StatusInternalServerError:
		return fmt.Sprintf("%v: internal error: %v", message, apiError.Message)
	case http.StatusGone:
		return fmt.Sprintf("%v: endpoint no longer exists, upgrade the CLI: %v", message, apiError.Message)
	default:
		return fmt.Sprintf("%v: request failed: %v", message, apiError.Message)
	}
}

func (u UserFacingError) Unwrap() error {
	return u.Underlying
}
