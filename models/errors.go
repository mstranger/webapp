package models

import "strings"

const (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound modelError = "models: resource not found"
	// ErrInvalidPassword is returned when an invalid password
	// is used when attemping to authenticate a user.
	ErrInvalidPassword modelError = "models: invalid password provided"
	// ErrEmailRequired is returned when an email address is
	// not provided when creating a user.
	ErrEmailRequired modelError = "models: email address is required"
	// ErrEmailInvalid is returned when an email address provided
	// does not match any of our requirements.
	ErrEmailInvalid modelError = "models: email addres is not valid"
	// ErrEmailTaken is returned when an update or create is
	// attempted with an email address that is already in use.
	ErrEmailTaken modelError = "models: email address is already taken"
	// ErrPasswordRequired is returned when a create is attempted
	// without a user password provider.
	ErrPasswordRequired modelError = "models: password is required"
	// ErrPasswordTooShort is returned when an update or create
	// attempted with a user password that is less than 8 chars.
	ErrPasswordTooShort modelError = "models: password must be at least 8 chars long"
	ErrTitleRequired    modelError = "models: title is required"

	ErrTokenInvalid modelError = "models: token provided is not valid"

	// ErrInvalidID is returned when an invalid ID is provided
	// to a method like Delete.
	ErrInvalidID privateError = "models: ID provided was invalid"
	// ErrRememberTooShort is returned when a remember token is
	// not at least 23 bytes.
	ErrRememberTooShort privateError = "models: remember token must be at least 32 bytes"
	// ErrRememberRequired is returned when a create or update
	// is attempted without a user remember token hash.
	ErrRememberRequired privateError = "models: remember token is required"
	ErrUserIDRequired   privateError = "models: user ID is required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}
