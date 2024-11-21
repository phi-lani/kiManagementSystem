package config

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/phi-lani/kimanagementsystem/bindings"
)

var (
	client   *ethclient.Client
	auth     *bind.TransactOpts
	instance *bindings.Bindings // Your contract bindings
)

// InitBlockchain initializes the Ethereum client and contract instance
func InitBlockchain(contractAddress string) {
	var err error

	// Connect to the Ethereum client
	client, err = ethclient.Dial("http://localhost:8545") // Use your Ethereum node's URL
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Load the private key from environment variables
	privateKeyHex := "df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e"

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create an authenticated transactor with the chain ID
	chainID := big.NewInt(31337) // Update with the correct chain ID
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set the gas limit and gas price
	auth.GasLimit = uint64(300000)
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// Initialize the contract instance
	address := common.HexToAddress(contractAddress)
	instance, err = bindings.NewBindings(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	log.Println("Blockchain client and contract instance initialized successfully")
}

// GetClient returns the Ethereum client
func GetClient() *ethclient.Client {
	return client
}

// GetAuth returns the transaction options
func GetAuth() *bind.TransactOpts {
	return auth
}

// GetContractInstance returns the contract instance
func GetContractInstance() *bindings.Bindings {
	return instance
}
