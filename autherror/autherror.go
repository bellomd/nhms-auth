// Package autherror is an error package that contains
// errors related to auth application.
package autherror

// AuthError is an error type that is returned
// when an entity is not found
type AuthError struct {
	s string
}

func (e *AuthError) Error() string {
	return e.s
}

// New create a new AuthError
func New(message string) error {
	return &AuthError{message}
}
