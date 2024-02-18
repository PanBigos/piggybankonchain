package database

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type DatabaseError uint8

var (
	ErrDbNotUnique DatabaseError = 1
	ErrDbNotFound  DatabaseError = 2
)

func (err DatabaseError) Error() string { return err.String() }

func (err DatabaseError) String() string {
	switch err {
	case ErrDbNotUnique:
		return "ErrDbNotUnique"
	case ErrDbNotFound:
		return "ErrNotFound"
	default:
		return "unknown"
	}
}

type StoreErrorInterceptor struct{}

func NewPostgresErrorInterceptor() StoreErrorInterceptor {
	return StoreErrorInterceptor{}
}

func (s *StoreErrorInterceptor) InterceptError(err error) error {
	if err == nil {
		return nil
	}
	switch v := err.(type) {
	case *pq.Error:
		return s.handlePqError(v)
	default:
		if errors.Is(err, sql.ErrNoRows) {
			return ErrDbNotFound
		}
		return err
	}
}

func (s *StoreErrorInterceptor) handlePqError(err *pq.Error) error {
	switch err.Code.Name() {
	case "forgeign_key_violation", "unique_violation":
		return ErrDbNotUnique
	default:
		return err
	}
}
