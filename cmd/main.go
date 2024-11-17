package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/phi-lani/kimanagementsystem/bindings"
)

// Contract ABI and address
// const contractABI = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"kiAddress","type":"address"},{"indexed":false,"internalType":"uint256","name":"documentIndex","type":"uint256"},{"indexed":false,"internalType":"string","name":"documentHash","type":"string"}],"name":"DocumentUploaded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"kiAddress","type":"address"},{"indexed":false,"internalType":"uint256","name":"documentIndex","type":"uint256"}],"name":"DocumentVerified","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"kiAddress","type":"address"},{"indexed":false,"internalType":"string","name":"name","type":"string"}],"name":"KIRegistered","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"kiAddress","type":"address"}],"name":"KIVerified","type":"event"},{"inputs":[],"name":"admin","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_kiAddress","type":"address"},{"internalType":"uint256","name":"_documentIndex","type":"uint256"}],"name":"getDocumentDetails","outputs":[{"internalType":"string","name":"hash","type":"string"},{"internalType":"bool","name":"isVerified","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_kiAddress","type":"address"}],"name":"getKIDetails","outputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"qualification","type":"string"},{"internalType":"string","name":"licenseType","type":"string"},{"internalType":"string","name":"experience","type":"string"},{"internalType":"uint256","name":"documentCount","type":"uint256"},{"internalType":"bool","name":"isVerified","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"keyIndividuals","outputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"qualification","type":"string"},{"internalType":"string","name":"licenseType","type":"string"},{"internalType":"string","name":"experience","type":"string"},{"internalType":"uint256","name":"documentCount","type":"uint256"},{"internalType":"bool","name":"isVerified","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"_name","type":"string"},{"internalType":"string","name":"_qualification","type":"string"},{"internalType":"string","name":"_licenseType","type":"string"},{"internalType":"string","name":"_experience","type":"string"}],"name":"registerKI","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"_documentHash","type":"string"}],"name":"uploadDocument","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_kiAddress","type":"address"},{"internalType":"uint256","name":"_documentIndex","type":"uint256"}],"name":"verifyDocument","outputs":[],"stateMutability":"nonpayable","type":"function"}]` // Replace with your contract's ABI
const contractAddress = " 0x5FbDB2315678afecb367f032d93F642f64180aa3" // Replace with your deployed contract address

func main() {
	// Connect to the Ethereum client
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	// Load the private key of the account you want to use
	privateKey, err := crypto.HexToECDSA("df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e") // Replace with your private key
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create an authenticated transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337)) // Use the appropriate chain ID
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set up transaction options
	auth.GasLimit = uint64(300000) // Set an appropriate gas limit
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// Load the contract
	address := common.HexToAddress(contractAddress)
	instance, err := bindings.NewBindings(address, client) // Replace with your contract's Go bindings
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}

	// Example 1: Register a new Key Individual
	tx, err := instance.RegisterKI(auth, "John Doe", "BSc Electrical Engineering", "Type A License", "5 years experience")
	if err != nil {
		log.Fatalf("Failed to register Key Individual: %v", err)
	}
	fmt.Printf("Key Individual registered successfully, tx hash: %s\n", tx.Hash().Hex())

	// Example 2: Upload a document
	// tx, err = instance.UploadDocument(auth, "abc123documenthash") // Replace with your document hash
	// if err != nil {
	// 	log.Fatalf("Failed to upload document: %v", err)
	// }
	// fmt.Printf("Document uploaded successfully, tx hash: %s\n", tx.Hash().Hex())

	// // Example 3: Get Key Individual details
	// kiAddress := common.HexToAddress("YOUR_KEY_INDIVIDUAL_ADDRESS") // Replace with the Key Individual's address
	// name, qualification, licenseType, experience, documentCount, isVerified, err := instance.GetKIDetails(nil, kiAddress)
	// if err != nil {
	// 	log.Fatalf("Failed to get Key Individual details: %v", err)
	// }
	// fmt.Printf("Name: %s, Qualification: %s, License Type: %s, Experience: %s, Document Count: %d, Is Verified: %t\n",
	// 	name, qualification, licenseType, experience, documentCount, isVerified)

	// // Example 4: Admin verifying a document
	// adminAuth := auth                                                      // Use the same or different auth for admin
	// tx, err = instance.VerifyDocument(adminAuth, kiAddress, big.NewInt(0)) // Verify the first document
	// if err != nil {
	// 	log.Fatalf("Failed to verify document: %v", err)
	// }
	// fmt.Printf("Document verified successfully, tx hash: %s\n", tx.Hash().Hex())

	// Example 5: Get details of a specific document
	// hash, isVerified, err := instance.GetDocumentDetails(nil, kiAddress, big.NewInt(0)) // Get the first document
	// if err != nil {
	// 	log.Fatalf("Failed to get document details: %v", err)
	// }
	// fmt.Printf("Document Hash: %s, Is Verified: %t\n", hash, isVerified)
}
