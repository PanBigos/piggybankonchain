package types

import (
	"crypto/ecdsa"
	"fmt"
	"regexp"
	"strings"

	"time"

	"github.com/Exca-DK/pegism/core/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

const (
	registerMsg = "I grant permission to %v for %v."
	header      = "header:"
)

type msgMatcher struct {
	template *regexp.Regexp
}

func NewMessageMatcher(msg string) msgMatcher {
	return msgMatcher{template: regexp.MustCompile(msg)}
}

func (r msgMatcher) matches(msg string) bool {
	return !r.template.Match([]byte(msg))
}

type RegistrationMatcher struct {
	msg     string
	matcher msgMatcher
}

func NewRegistrationMatcher(entity string, action string) RegistrationMatcher {
	msg := fmt.Sprintf(registerMsg, entity, action)
	return RegistrationMatcher{matcher: NewMessageMatcher(msg), msg: msg}
}

func (r RegistrationMatcher) Matches(msg string) bool {
	diff := len(msg) - len(r.msg)
	if diff != 0 {
		return false
	} else {
		return !r.matcher.matches(msg)
	}
}

type AuthMessage string

func (msg AuthMessage) Sign(key *ecdsa.PrivateKey) (hash common.Hash, sig []byte, err error) {
	hash, sig, err = crypto.EthersSign([]byte(msg), key)
	return
}

type RegistrationMessageTemplate struct {
	raw     string
	matcher RegistrationMatcher
}

func (r RegistrationMessageTemplate) NewMessage(
	address common.Address,
	deadline time.Time,
) AuthMessage {
	return AuthMessage(
		r.raw + fmt.Sprintf("\t%v/%v/%v", header, address.Hex(), deadline.Format(time.RFC3339)),
	)
}

func (r RegistrationMessageTemplate) recover(
	msg AuthMessage,
) (string, time.Time, common.Address, bool) {
	elems := strings.Split(string(msg), header)
	if len(elems) != 2 {
		return "", time.Time{}, common.Address{}, false
	}

	registerMessage := strings.TrimSpace(elems[0])
	if !r.IsKnown(registerMessage) {
		return "", time.Time{}, common.Address{}, false
	}

	metadata := strings.Split(elems[1], "/")

	if len(metadata) != 3 {
		return "", time.Time{}, common.Address{}, false
	}
	metadata = metadata[1:]

	deadline, ok := r.getDeadline(metadata[1])
	if !ok {
		return "", time.Time{}, common.Address{}, false
	}

	addr, ok := r.getAcquirer(metadata[0])
	if !ok {
		return "", time.Time{}, common.Address{}, false
	}

	return registerMessage, deadline, addr, true
}

type AuthHeader struct {
	Acquirer common.Address
	Deadline time.Time
}

func (r RegistrationMessageTemplate) Recover(msg AuthMessage) (string, AuthHeader, bool) {
	message, deadline, addr, ok := r.recover(msg)
	if !ok {
		return "", AuthHeader{}, false
	}

	return message, AuthHeader{Acquirer: addr, Deadline: deadline}, true
}

func (r RegistrationMessageTemplate) getDeadline(msg string) (time.Time, bool) {
	time, err := time.Parse(time.RFC3339, msg)
	if err != nil {
		return time, false
	}
	return time, true
}

func (r RegistrationMessageTemplate) getAcquirer(msg string) (common.Address, bool) {
	if !common.IsHexAddress(msg) {
		return common.Address{}, false
	}
	return common.HexToAddress(msg), true
}

func (r RegistrationMessageTemplate) IsKnown(msg string) bool {
	return r.matcher.Matches(msg)
}

type RegistrationMessageFactory struct{}

func (RegistrationMessageFactory) GenerateNewTemplate(
	entity string,
	action string,
) *RegistrationMessageTemplate {
	return &RegistrationMessageTemplate{
		raw:     fmt.Sprintf(registerMsg, entity, action),
		matcher: NewRegistrationMatcher(entity, action),
	}
}

func RandomAuth(
	deadline time.Time,
	template *RegistrationMessageTemplate,
) (string, common.Address) {
	sig, addr := SignRandom(func(a common.Address) AuthMessage {
		temp := template.NewMessage(a, deadline)
		return temp
	})
	return hexutil.Encode(sig), addr
}

func SignRandom(
	genMsg func(common.Address) AuthMessage,
) (sig []byte, address common.Address) {
	privateKey, err := ethCrypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic(err)
	}

	address = ethCrypto.PubkeyToAddress(*publicKeyECDSA)
	generated := genMsg(address)
	_, sig, err = generated.Sign(privateKey)
	if err != nil {
		panic(fmt.Sprintf("signging failure. error: %v", err))
	}
	return
}

func SignAuth(
	privateKey *ecdsa.PrivateKey,
	msg AuthMessage,
) []byte {
	_, sig, err := msg.Sign(privateKey)
	if err != nil {
		panic(fmt.Sprintf("signging failure. error: %v", err))
	}
	return sig
}

func RandomKey() (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := ethCrypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic(err)
	}
	return privateKey, ethCrypto.PubkeyToAddress(*publicKeyECDSA)
}

func PrivateKeyToPublic(privateKey *ecdsa.PrivateKey) common.Address {
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("shouldnt happen")
	}
	return ethCrypto.PubkeyToAddress(*publicKeyECDSA)
}
