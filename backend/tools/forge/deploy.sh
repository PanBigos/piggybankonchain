# Available Accounts
# ==================

# (0) "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" (10000 ETH)
# (1) "0x70997970C51812dc3A010C7d01b50e0d17dc79C8" (10000 ETH)
# (2) "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC" (10000 ETH)
# (3) "0x90F79bf6EB2c4f870365E785982E1f101E93b906" (10000 ETH)
# (4) "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65" (10000 ETH)
# (5) "0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc" (10000 ETH)
# (6) "0x976EA74026E726554dB657fA54763abd0C3a0aa9" (10000 ETH)
# (7) "0x14dC79964da2C08b23698B3D3cc7Ca32193d9955" (10000 ETH)
# (8) "0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f" (10000 ETH)
# (9) "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720" (10000 ETH)

# Private Keys
# ==================

# (0) 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
# (1) 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
# (2) 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
# (3) 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
# (4) 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
# (5) 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
# (6) 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e
# (7) 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
# (8) 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
# (9) 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

# Wallet
# ==================
# Mnemonic:          test test test test test test test test test test test junk
# Derivation path:   m/44'/60'/0'/0/

#bin/sh

echo "Waiting for anvil to start at 8545..."

while ! nc -z localhost 8545; do   
  sleep 0.1 # wait for 1/10 of the second before check again
done

echo "Anvil launched"

echo "Deploying ERC20 tokens"

forge init demo --no-git
cp /app/erc20/erc20.sol ./demo/src
cp -R /app/piggy_bank_contracts/src ./demo/src/piggy_bank_contracts

cd ./demo && git init -b main
forge install OpenZeppelin/openzeppelin-contracts --no-commit

# DEPLOY 1K FUSD
# Deployer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
# Deployed to: 0x5FbDB2315678afecb367f032d93F642f64180aa3
# Transaction hash: 0x3149f3f638abd283e4de5984a71e0b0021f004f3552be27bd6ce047e7e1f5a55
forge create --rpc-url localhost:8545 --constructor-args "ForgeUSD" "FUSD" --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 ./src/erc20.sol:Token 

# DEPLOY 1K DAY
# Deployer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
# Deployed to: 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
# Transaction hash: 0xd8178576e7683cec237825fac7a7b35cb698175c914d7356fad5e37708072d4c
forge create --rpc-url localhost:8545 --constructor-args "ForgeDai" "FDAI" --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 ./src/erc20.sol:Token 

echo "ERC20 tokens deployed!"

echo "Deploying PiggyBankRouter."
# Deployer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
# Deployed to: 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
# Transaction hash: 0x91f35c845ac7c7461ec0d641c1507fb2728c45725cdda6dd3c0a0e3e642d15be
# _feeRate, _maxFee
forge create --rpc-url localhost:8545 --constructor-args 50 500 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 ./src/piggy_bank_contracts/fuse/PiggyBankRouter.sol:PiggyBankRouter 

echo "Deploying TimeLockedPiggyBankFactory."
# Deployer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
# Deployed to: 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
# Transaction hash: 0x77dce5c64a4a5d1e51a9f5741707067c355911f42c1fe47b25ef1b396a24cbf7
forge create --rpc-url localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 ./src/piggy_bank_contracts/fuse/TimeLockedPiggyBankFactory.sol:TimeLockedPiggyBankFactory 

# This has to be at the end of the file!!
# Some tests are matching this value for consistent foundry setup.
echo "**Deployments finished**"