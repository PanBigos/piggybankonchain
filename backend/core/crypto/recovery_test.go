package crypto

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestEthersSignatureCompability(t *testing.T) {
	var (
		msg        = "test"
		sig        = "0xf755d9a72d5b7386765e7f0e833af68795b739a267122dae933f41b781b5aed0626ce3263308ebd4c37bed84319b66da2794368771046825bd89b98ba68c4e871b"
		testKey, _ = crypto.HexToECDSA(
			"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		)
		testaddr = crypto.PubkeyToAddress(testKey.PublicKey)
	)

	recovered, err := EthersRecoverAddress(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, recovered, testaddr, "recovered address should match reference")
}

func TestGolangSignatureCompability(t *testing.T) {
	var (
		msg        = "test"
		testKey, _ = crypto.HexToECDSA(
			"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		)
		testaddr = crypto.PubkeyToAddress(testKey.PublicKey)
	)
	_, sig, err := EthersSign([]byte(msg), testKey)
	if err != nil {
		t.Fatal(err)
	}
	recovered, err := EthersRecoverAddress(msg, hexutil.Encode(sig))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, recovered, testaddr, "recovered address should match reference")
}
