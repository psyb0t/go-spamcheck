package spamcheck

import "errors"

var (
	// ErrEmptyEmail is returned when the raw email passed to Check is empty.
	ErrEmptyEmail = errors.New("raw email is empty")
	// ErrUnexpectedStatus is returned when the API responds with a non-200 status.
	ErrUnexpectedStatus = errors.New("unexpected response status")
	// ErrDecodeResponse is returned when the API response cannot be decoded.
	ErrDecodeResponse = errors.New("failed to decode response")
	// ErrAPIFailure is returned when the API responds with success=false.
	ErrAPIFailure = errors.New("api reported failure")
)
