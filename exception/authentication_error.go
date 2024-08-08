package exception

type AuthenticationError struct {
	Error string
}

func NewAuthenticationError(error string) AuthenticationError {
	return AuthenticationError{Error: error}
}
