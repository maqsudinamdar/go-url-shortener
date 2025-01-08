package domain

import "fmt"


const (

	// DbQueryFail represents DB query failures
	DbQueryFail = "DB_QUERY_FAIL"

	// DbNotSupported represents DB not supported operation
	DbNotSupported = "DB_NOT_SUPPORTED"

	// EntityNotExist represents error that entity doesn't exist in DB
	EntityNotExist = "ENTITY_NOT_EXIST"
)

// ErrDbQuery will be mapped to 409 conflict status
type ErrDbQuery struct {
	Err error
}

func (e ErrDbQuery) Error() string {
	return fmt.Sprintf("%s: %s", DbQueryFail, e.Err)
}

// ErrDbNotSupported will be mapped to 409 conflict status
type ErrDbNotSupported struct {
	Err error
}

func (e ErrDbNotSupported) Error() string {
	return fmt.Sprintf("%s: %s", DbNotSupported, e.Err)
}

// ErrEntityNotExist will be mapped to 409 conflict status
type ErrEntityNotExist struct {
	Err error
}

func (e ErrEntityNotExist) Error() string {
	return fmt.Sprintf("%s: %s", EntityNotExist, e.Err)
}