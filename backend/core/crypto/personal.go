package crypto

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var ethersMsg = "\x19Ethereum Signed Message:\n%d%s"

func EthersHash(data []byte) common.Hash {
	msg := fmt.Sprintf(ethersMsg, len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

func EthersSign(data []byte, key *ecdsa.PrivateKey) (hash common.Hash, sig []byte, err error) {
	hash = EthersHash(data)
	sig, err = crypto.Sign(hash.Bytes(), key)
	return
}
