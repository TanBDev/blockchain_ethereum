package main

import (
	"math/big"
	"strings"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	//	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/common"
	"gitlab.com/bienestar-mx/api/pkg/blockchain_ethereum/smartContract"
	"github.com/ethereum/go-ethereum/common"
)



	var (
		blockchain     *ethclient.Client
		auth     *bind.TransactOpts
		//contract *contracts.Payment
		address string
	)
 

	//deploys and connects to the function in smart contract
	func deployFunctions(userId int64, password string) {
	//func main(){
		//Setup my account key
		

		key := importPrivateKey(userId, password)
		//Set up the user 
		auth, err := bind.NewTransactor(strings.NewReader(key), password)
		if err != nil {
			fmt.Printf("Cannot bind to account: %v : %v", err, auth)
		}
		auth.GasLimit = 131193

		//Set up the blockchain environnement: connection to a node
		blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/44029fa59f854e649e0a4e9f3691de56")
		if err != nil {
			fmt.Printf("Cannot bind to blockchain: %v : %v", err, blockchain)

		}

		//newPay, err := blockchainStartNewPayment(200, 129, 130)
		//if err != nil {
		//	fmt.Println("Unable to bind to deployed instance of contract:%v", err)
		//}
		//fmt.Printf("new Payment:%v", newPay)
	}



	// ---------------------------  SMART CONTRACT'S FUNCTIONS CALLS  -----------------------------------

		//Create a payment
		func blockchainStartNewPayment(price int64, externalPayerId int64, externalReceiverId int64) (*types.Transaction, error){
			contract, err := contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
			if err != nil {
				fmt.Printf("Unable to bind to deployed instance of contract:%v", err)
			}
	
			// StartNewPayment todavia no funciona porque no subi el smart contract
			hello, er := contract.StartNewPayment(auth, big.NewInt(price), big.NewInt(externalPayerId), big.NewInt(externalReceiverId))
			//hello, er := contract.StartNewPayment(auth, big.NewInt(200), big.NewInt(132), big.NewInt(133))
			if err != nil {
				fmt.Println("Unable to bind to deployed instance of contract:%v :%v", er, hello)
			}
			fmt.Printf("new payment :%v", hello)
			return hello, er
		}
		
	

		//Returns the infos of the payment when we send its ID
		func blockchainGetPaymentById(externalPaymentId int64) (*big.Int, *big.Int, *big.Int, *big.Int, error){
			contract, err := contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
			if err != nil {
				fmt.Println("Unable to bind to deployed instance of contract:%v", err)
			}
	
			price, payerId, receiverId, status, err := contract.GetPaymentById(&bind.CallOpts{Pending: true}, big.NewInt(externalPaymentId))
			//price, payerId, receiverId, err := contract.GetPaymentById(&bind.CallOpts{Pending: true}, big.NewInt(134))
			if err != nil {
				fmt.Printf("Unable to reach the payment:%v", err)
			}
			fmt.Printf("Here is the payment:%v :%v :%v", price, payerId, receiverId)
			return price, payerId, receiverId, status, err
		}





	//returns 2 tabs of the addresses of the payments
	//problem: returns the the addresses and then a bunch of zeros. have to delete the zeros
	func blockchainGetPaymentByClient(externalClientId int64) ([]*big.Int, []*big.Int, error){
		contract, err := contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
		if err != nil {
			fmt.Println("Unable to bind to deployed instance of contract:%v", err)
		}
		tabPaidTransac, tabReceivTransac, err := contract.GetPaymentsByOwner(&bind.CallOpts{Pending: true}, big.NewInt(externalClientId))
		return tabPaidTransac, tabReceivTransac, err
	}  

/*
	
	// Creates a new client inside the blockchain
//	func blockchainANewClient(clientName string) (*types.Transaction, error){
//		contract, err := contracts.NewPayment(common.HexToAddress("0xe662b17cc5aa2423bdb3de750e33141791522c1a"), blockchain)
		if err != nil {
			fmt.Printf("Unable to bind to deployed instance of contract:%v", err)
		}

		fmt.Println("contrat:%v", contract)

		clientId, er := contract.ANewClient(auth, clientName)
		if err != nil {
			fmt.Println("Unable to create client:%v :%v", er)
		}
		return clientId, er
	}
*/