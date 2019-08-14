package main

import (
	"strings"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	//	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/common"
	"gitlab.com/bienestar-mx/api/pkg/blockchain_ethereum/smartContract"

)

/*
func main(){

	deploySmartContract("Imodollar", "ImD", big.NewInt(100000), common.HexToAddress("0x4da93B7EeCe99748afac5FDf3B99138177710ff3"), big.NewInt(1))

	}
	*/



/* 
---------------  THIS FUNCTION DEPLOYS THE SMART CONTRACT ON AN ETHEREUM NODE  --------------------
Must be deployed at the opening of the app
*/

func deploySmartContract(tokenName string, tokenSymbol string, tokenInitialBalance *big.Int, tokenDevAddress common.Address, setSellPrice *big.Int)  (common.Address){
//func main()  {
	
	/*
	This is the private key of the palo's account 
	The acount used will be the only one to be able to modify the smart contract
	In the case of a public blockchain this account will have to pay a few transactions
	*/
	key := `{"address":"4da93b7eece99748afac5fdf3b99138177710ff3","crypto":{"cipher":"aes-128-ctr","ciphertext":"53492bc230964b57761e9aebb288f3e3e0c9311605dc8d74a2417b8d38bda12a","cipherparams":{"iv":"0fa80fd44a285f6871d1b234c1311ff8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"9af823c6e8059c8de80611277125fc050fed31b42686b499df2b44289d3694a9"},"mac":"6e8451f4fe4054a978444bc551af3826e43716f97ed40718e6bdefb38c61db27"},"id":"95daa908-efc1-433a-b402-d5e2e0dc4d19","version":3}`


	//Use the key and my account's password to bind to my account
	auth, err := bind.NewTransactor(strings.NewReader(key), "Ornithorynque1")
	if err != nil {
		fmt.Printf("Cannot bind to account: %v : %v", err, auth)
	}
	//Gas limit. In some case it shall be exactly equal to the gas transaction cost!
	auth.GasLimit = 3000000

	/*
	Link to the blockchain
	Here I connect to an existing node on the test plateform infura but once my contract is online 
	I shall put the address of the contract in the ethclient.Dial() function
	*/
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/44029fa59f854e649e0a4e9f3691de56")
	if err != nil {
		fmt.Printf("Cannot bind to blockchain: %v : %v", err, blockchain)

	}


	//Deploy contract
	//tokenInitialBalance : How much tokens the contract will dispose
	//tokenDevAddress : The ethereum account in charge of the contract that will be able to modify it (usually the one that is used to deploy)  
	address, tx, _, err := contracts.DeployPayments(auth, blockchain, tokenName, tokenSymbol, tokenInitialBalance, tokenDevAddress, setSellPrice)
	if err != nil {
		fmt.Printf("Failed to deploy new token contract: %v", err)
	}


	if err != nil {
		fmt.Printf("Failed to deploy the condo contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		fmt.Printf("Expected a valid deployment address. Received empty address byte array instead")
	}

	/*
	adress = address of the smart contract. example : 0xfb7974b7616a0c0dbd08c7fee7f1291548045e33
	we need this value for each function call
	have to copie paste this value in the function :
	contract, err := contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
	*/
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	return (address)
}

