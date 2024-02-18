package e2e_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/Exca-DK/pegism/contracts/erc20"
	"github.com/Exca-DK/pegism/contracts/piggy/factory"
	"github.com/Exca-DK/pegism/contracts/piggy/router"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/Exca-DK/pegism/core/log"
	"github.com/Exca-DK/pegism/core/types"
	pegismv1 "github.com/Exca-DK/pegism/gen/go/proto/v1"
	"github.com/Exca-DK/pegism/service/app"
	"github.com/Exca-DK/pegism/service/backend"
	"github.com/Exca-DK/pegism/service/rpc"
	"github.com/Exca-DK/pegism/service/rpc/auth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

var (
	// testKey is a private key to use for funding a tester account. it is first private key of foundry anvil node
	testFirstKey, _ = crypto.HexToECDSA(
		"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	)
	firstAddress = crypto.PubkeyToAddress(testFirstKey.PublicKey)

	// testKey is a private key to use for funding a tester account. it is second private key of foundry anvil node
	testSecondKey, _ = crypto.HexToECDSA(
		"59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	)
	secondAddress = crypto.PubkeyToAddress(testSecondKey.PublicKey)

	testUsdAddress = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	testDaiAddress = common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
	factoryAddress = common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")
	routerAddress  = common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")
)

func TestService(t *testing.T) {
	log.ChangeLvl(log.LvlTrace)
	nodeEndpoint := "http://127.0.0.1:8545"
	application := app.New()
	require.NoError(t, application.Setup(app.Config{
		Clock: clock.NewClock(),
		Db: struct {
			Username string
			Password string
			Db       string
			Endpoint string
		}{Username: "foo", Password: "bar", Db: "baz", Endpoint: "127.0.0.1:5432"},
		IssueSecret:     "very-very-secret-password-0123456789",
		AccessDuration:  10 * time.Minute,
		RefreshDuration: 24 * time.Hour * 31,
		Rpc: struct {
			Host           string
			Port           int
			GatewayPort    int
			MaxMsgSizeInMb uint64
		}{
			Host:        "127.0.0.1",
			Port:        5433,
			GatewayPort: 5434,
		},
		FactoryAddress: factoryAddress,
		RouterAddress:  routerAddress,
		NodeEndpoint:   nodeEndpoint,
	}))
	var (
		rpcService     *rpc.RpcService
		backendService *backend.Service
	)
	require.NoError(t, application.GetService(&rpcService))
	require.NoError(t, application.GetService(&backendService))

	authServer := rpcService.GetAuthServer()
	profileServer := rpcService.GetProfileServer()

	firstAuthMsg, err := getAuth(testFirstKey, authServer)
	require.NoError(t, err)

	secondAuthMsg, err := getAuth(testSecondKey, authServer)
	require.NoError(t, err)

	registedResponse, err := profileServer.IsRegistered(
		context.Background(),
		&pegismv1.IsRegisteredRequest{
			Address: firstAddress.Hex(),
		},
	)
	require.NoError(t, err)
	require.False(t, registedResponse.GetRegistered())

	firstAuthCtx := prepareAuthCtx(context.Background(), firstAuthMsg.AccessToken)
	secondAuthCtx := prepareAuthCtx(context.Background(), secondAuthMsg.AccessToken)

	_, err = profileServer.Register(firstAuthCtx, nil)
	require.NoError(t, err)

	_, err = profileServer.Register(secondAuthCtx, nil)
	require.NoError(t, err)

	registedResponse, err = profileServer.IsRegistered(
		context.Background(),
		&pegismv1.IsRegisteredRequest{
			Address: firstAddress.Hex(),
		},
	)
	require.NoError(t, err)
	require.True(t, registedResponse.GetRegistered())

	registedResponse, err = profileServer.IsRegistered(
		context.Background(),
		&pegismv1.IsRegisteredRequest{
			Address: secondAddress.Hex(),
		},
	)
	require.NoError(t, err)
	require.True(t, registedResponse.GetRegistered())

	profileResponse, err := profileServer.GetProfile(
		context.Background(),
		&pegismv1.GetProfileRequest{
			Address: firstAddress.Hex(),
		},
	)
	require.NoError(t, err)
	require.Len(t, profileResponse.Profile.Piggies, 0)
	require.Equal(t, common.HexToAddress(profileResponse.Profile.Address), firstAddress)

	client, err := ethclient.Dial(nodeEndpoint)
	require.NoError(t, err)

	chainId, err := client.ChainID(context.Background())
	require.NoError(t, err)

	daiClient, err := erc20.NewToken(testDaiAddress, client)
	require.NoError(t, err)

	transactorFirst, err := bind.NewKeyedTransactorWithChainID(testFirstKey, chainId)
	require.NoError(t, err)
	factoryContract, err := factory.NewFactory(factoryAddress, client)
	require.NoError(t, err)
	routerContract, err := router.NewRouter(routerAddress, client)
	require.NoError(t, err)

	require.NoError(t, err)
	tx, err := factoryContract.CreatePiggyBank(
		transactorFirst,
		secondAddress,
		big.NewInt(time.Now().Add(1*time.Hour).Unix()),
	)

	require.NoError(t, err)
	time.Sleep(1 * time.Second)
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	// equal succesfull
	require.Equal(t, uint64(1), receipt.Status)
	require.NoError(t, backendService.Api().NotifyNewTransaction(tx.Hash()))

	// approve for spending
	tx, err = daiClient.Approve(transactorFirst, routerAddress, big.NewInt(50000*4))
	require.NoError(t, err)
	time.Sleep(1 * time.Second)
	receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	tx, err = send_message(
		"bar",
		"foo",
		routerContract,
		transactorFirst,
		secondAddress,
		true,
		client,
	)
	require.NoError(t, err)
	time.Sleep(3 * time.Second)
	tx, _, err = client.TransactionByHash(context.Background(), tx.Hash())
	require.NoError(t, err)
	receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)
	require.NoError(t, backendService.Api().NotifyNewTransaction(tx.Hash()))

	profileResponse, err = profileServer.GetProfile(
		context.Background(),
		&pegismv1.GetProfileRequest{
			Address: secondAddress.Hex(),
		},
	)
	require.NoError(t, err)
	require.Len(t, profileResponse.Profile.Piggies, 1)
	require.Equal(t, common.HexToAddress(profileResponse.Profile.Address), secondAddress)
	require.Len(t, profileResponse.Profile.Piggies[0].Messages, 1)
	require.Equal(t, profileResponse.Profile.Piggies[0].Messages[0].Nick, "foo")
	require.Equal(t, profileResponse.Profile.Piggies[0].Messages[0].Content, "bar")
}

func getAuth(
	key *ecdsa.PrivateKey,
	server *auth.AuthServer,
) (*pegismv1.AuthResponse_AuthorizationDataContainer, error) {
	resigsterResponse, err := server.GetAuthMessage(
		context.Background(),
		&pegismv1.AuthMessageRequest{
			Address: types.PrivateKeyToPublic(key).Hex(),
		},
	)
	if err != nil {
		return nil, err
	}
	signed := hexutil.Encode(
		types.SignAuth(key, types.AuthMessage(resigsterResponse.Message.Content)),
	)
	loginResponse, err := server.Authorize(context.Background(), &pegismv1.AuthRequest{
		Args: &pegismv1.AuthArgs{
			Sig: signed,
			Msg: resigsterResponse.Message.Content,
		},
	})
	if err != nil {
		return nil, err
	}
	return loginResponse.Data, nil
}

func prepareAuthCtx(
	ctx context.Context,
	token string,
) context.Context {
	return metadata.NewIncomingContext(
		ctx,
		metadata.Pairs(
			auth.AuthorizationHeader,
			token,
		),
	)
}

func send_message(
	msg string,
	nickname string,
	router *router.Router,
	opts *bind.TransactOpts,
	to common.Address,
	erc20 bool,
	client *ethclient.Client,
) (*ethTypes.Transaction, error) {
	return router.TransferTokenWithFee(opts, testDaiAddress, to, big.NewInt(50000))
}
