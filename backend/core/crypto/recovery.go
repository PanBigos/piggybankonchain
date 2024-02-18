package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Recover address recovers address with ethers-metamask compability.
// Important! Signature must be a valid hexstring.
func EthersRecoverAddress(message string, signature string) (common.Address, error) {
	// Hash the unsigned message
	var hash = EthersHash([]byte(message))
	// Get the bytes of the signed message
	decodedSig, err := hexutil.Decode(signature)
	if err != nil {
		return common.Address{}, err
	}

	// see https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L452-L459
	if len(decodedSig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if decodedSig[64] == 27 || decodedSig[64] == 28 {
		decodedSig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	test, err := crypto.SigToPub(hash.Bytes(), decodedSig)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*test), nil
}
