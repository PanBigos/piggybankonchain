package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Exca-DK/pegism/core/crypto"
	"github.com/Exca-DK/pegism/core/types"
	"google.golang.org/grpc/metadata"
)

const (
	AuthorizationHeader = "Authorization"
)

var (
	ErrMetadataMissing            = errors.New("missing metadata header")
	ErrAuthorizationHeaderMissing = errors.New("missing authorization header")
	ErrAuthorizationHeaderInvalid = errors.New("invalid authorization header")
	ErrInvalidAuthMessage         = errors.New("invalid auth message")
	ErrDeadlineElapsed            = errors.New("deadline elapsed")
	ErrFaker                      = errors.New("provided message is not produced by you")
)

func GrpcRecoverAuthTokenHeader(
	ctx context.Context,
) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrMetadataMissing
	}

	auth := md.Get(AuthorizationHeader)
	if len(auth) == 0 {
		return "", ErrAuthorizationHeaderMissing
	} else if len(auth) != 1 {
		return "", ErrAuthorizationHeaderInvalid
	}

	return auth[0], nil
}

type IAuthMsg interface {
	GetSig() string
	GetMsg() string
}

func ValidateAuth(req IAuthMsg, header types.AuthHeader) error {
	if time.Until(header.Deadline) < 0 {
		return ErrDeadlineElapsed
	}

	address, err := crypto.EthersRecoverAddress(req.GetMsg(), req.GetSig())
	if err != nil {
		return err
	}

	if address != header.Acquirer {
		return ErrFaker
	}

	return nil
}
