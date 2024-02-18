package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnauthenticatedError(err error) error {
	return status.Errorf(codes.Unauthenticated, "unauthorized: %s", err)
}
