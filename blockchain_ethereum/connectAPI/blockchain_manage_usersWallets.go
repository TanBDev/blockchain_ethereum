package main

import (
	"fmt"
	"strings"

	"io/ioutil"
	"log"
	"math/big"
	"os"
	"bytes"
    	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	contracts "gitlab.com/bienestar-mx/api/pkg/blockchain_ethereum/smartContract"
)

/*
Doesn't work yet!!!
funtion to erase the firsts few caracters of a string
because when we import the wallet of a user the string gotten as the wallet's direction adds "keystore://" before the path
*/
var(
	prefixToTrim = []byte(`keystore://`)
	sampleJSON   = []byte(`{"keystore://C:\Projects\Go\src\gitlab.com\bienestar-mx\api\pkg\blockchain_ethereum\connectAPI\wallets\UTC--2019-07-31T21-32-07.078488600Z--cd913ebe3063cdd1b10902912fd4c0edb655f9b3"}`)
)
type IDField string
type Data struct { ID IDField `json:"id"` }
// UnmarshalJSON provides custom unmarshalling to trim `keystore://` prefix from IDField
func (id *IDField) UnmarshalJSON(rawIDBytes []byte) error {

    // trim quotes and prefix
    trimmedID := bytes.TrimPrefix(bytes.Trim(rawIDBytes, `"`), prefixToTrim)

    // convert back to id field & assign
    *id = IDField(trimmedID)
    return nil
}

//creates a cliemt wallet in connectAPI/wallets
func createKs(userName string, clientPassword string) {

	//create account
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := clientPassword
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex()) // example: 0x20F8D42FB0F667F2E53930fed426f225752453b3
	file := account.URL.String()

	createAUser(userName, file, account.Address.Hex(), password)
}

func createAUser(userName string, file string, account string, password string) *types.Transaction {
	//Setup my account key
	// Use palo's account because we don't have the client yet to connect to his account
	key := `{"address":"4da93b7eece99748afac5fdf3b99138177710ff3","crypto":{"cipher":"aes-128-ctr","ciphertext":"53492bc230964b57761e9aebb288f3e3e0c9311605dc8d74a2417b8d38bda12a","cipherparams":{"iv":"0fa80fd44a285f6871d1b234c1311ff8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"9af823c6e8059c8de80611277125fc050fed31b42686b499df2b44289d3694a9"},"mac":"6e8451f4fe4054a978444bc551af3826e43716f97ed40718e6bdefb38c61db27"},"id":"95daa908-efc1-433a-b402-d5e2e0dc4d19","version":3}`
	
	//"Ornithorynque1" is the password of the account
	auth, err := bind.NewTransactor(strings.NewReader(key), "Ornithorynque1")
	if err != nil {
		fmt.Printf("Cannot bind to account: %v : %v", err, auth)
	}
	auth.GasLimit = 1443312
	//connect to the ethereum node
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/44029fa59f854e649e0a4e9f3691de56")
	if err != nil {
		fmt.Printf("Cannot bind to blockchain: %v : %v", err, blockchain)
	}

	// contract = contracts.NewPayments() permits to use a new function in the contract
	//0x9b9e1cc61f08bd344451f21e4fe6cb67e8f41d17 is the contract's address gicen by the deploySmartContract function in blockchain_main.go
	contract, err := contracts.NewPayments(common.HexToAddress("0x9b9e1cc61f08bd344451f21e4fe6cb67e8f41d17"), blockchain)
	if err != nil {
		fmt.Printf("Unable to bind to deployed instance of contract:%v", err)
	}
	fmt.Println("contrat:%v", contract)

	// contract.ANewClient() activate the function aNewClient in the smart contract to save the client into the blockchain
	clientId, er := contract.ANewClient(auth, userName, file, account)
	if err != nil {
		fmt.Println("Unable to create client:%v :%v", er, clientId)
	}
	fmt.Println("new client:%v", clientId)

	return clientId

}

//Imports the client's wallet
func importPrivateKey(userId int64, password string) string {

	//example of file :  "./tmp/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	//to get the path to the wallet's file we need to import the user from the blockchain 
	//userId is the address of the user inside the blockchain
	_, file, _ := getAUser(userId)

	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("cannot read the file", err)
	}

	//Import the wallet
	//first password is the old one, second the new one
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal("cannot import the account", err)
	}
	fmt.Println(account.Address.Hex()) // example: 0x20F8D42FB0F667F2E53930fed426f225752453b3

	//set the key
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		fmt.Println("json key failed to decrypt: %v", err)
	}
	fmt.Printf("%x", key.PrivateKey.D.Bytes())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", string(jsonBytes))

	return string(key.PrivateKey.D.Bytes())
}


//get the user's information inside the blockchain with his address
func getAUser(userId int64) (string, string, string) {

	// Use palo's account because we don't have the client yet to connect to his account
	key := `{"address":"4da93b7eece99748afac5fdf3b99138177710ff3","crypto":{"cipher":"aes-128-ctr","ciphertext":"53492bc230964b57761e9aebb288f3e3e0c9311605dc8d74a2417b8d38bda12a","cipherparams":{"iv":"0fa80fd44a285f6871d1b234c1311ff8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"9af823c6e8059c8de80611277125fc050fed31b42686b499df2b44289d3694a9"},"mac":"6e8451f4fe4054a978444bc551af3826e43716f97ed40718e6bdefb38c61db27"},"id":"95daa908-efc1-433a-b402-d5e2e0dc4d19","version":3}`

	//connects to the account
	//"Ornithorynque1" is the password of the wallet
	auth, err := bind.NewTransactor(strings.NewReader(key), "Ornithorynque1")
	if err != nil {
		fmt.Printf("Cannot bind to account: %v : %v", err, auth)
	}
	auth.GasLimit = 1443312
	
	//connects to the ethereum node
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/44029fa59f854e649e0a4e9f3691de56")
	if err != nil {
		fmt.Printf("Cannot bind to blockchain: %v : %v", err, blockchain)
	}

	//Connects to the smart contract
	//0x9b9e1cc61f08bd344451f21e4fe6cb67e8f41d17 is the contract's address gicen by the deploySmartContract function in blockchain_main.go
	contract, err := contracts.NewPayments(common.HexToAddress("0x9b9e1cc61f08bd344451f21e4fe6cb67e8f41d17"), blockchain)
	if err != nil {
		fmt.Printf("Unable to bind to deployed instance of contract:%v", err)
	}
	fmt.Println("contrat:%v", contract)

	//calls the getAClient function in the smart contract
	userName, file, clientAccountAddress, err := contract.GetAClient(&bind.CallOpts{Pending: true}, big.NewInt(userId))
	if err != nil {
		fmt.Println("Unable to create client:%v :%v", err)
	}
	fmt.Println("new client:%v : %v : %v", userName, file, clientAccountAddress)
	return userName, file, clientAccountAddress
}


	//NOT TESTED YET
/*
	//send ether to new account
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	unlockedKey, err := keystore.DecryptKey([]byte(key), "paloit2018")
	nonce, err := client.NonceAt(ctx, unlockedKey.Address, nil)
	if err != nil {
		log.Println("error getting latest nonce: ", err)
	}
	goodGasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Println("error getting suggested gas price: ", err)
	}
	log.Println("suggested gas price is: ", goodGasPrice)
	amount := big.NewInt(1)
	gasLimit := uint64(3000000) //big.NewInt(3000000)
	gasPrice := big.NewInt(21000)
	gasPrice.Div(goodGasPrice, big.NewInt(2))
	//common.HexToAddress(`0x56724a9e4d2bb2dca01999acade2e88a92b11a9e`)
	tx := types.NewTransaction(nonce, account.Address, amount, gasLimit, goodGasPrice, nil)
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, unlockedKey.PrivateKey)
	fmt.Println(signTx)
	txErr := client.SendTransaction(context.Background(), signTx)
	if txErr != nil {
		fmt.Println("send tx error:")
		panic(txErr)
	}
	fmt.Printf("send success tx.hash=%s\n", signTx.Hash().String())

*/
