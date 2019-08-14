package main

import (
	"math/big"
	"strings"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/bienestar-mx/api/pkg/blockchain_ethereum/smartContract"
	"github.com/ethereum/go-ethereum/common"
)



	var (
		blockchain     *ethclient.Client
		auth     *bind.TransactOpts
		address string
	)
 

	//deploys and connects to the function in smart contract
	func deployFunctions(userId int64, password string) {
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
	}



	// ---------------------------  SMART CONTRACT'S FUNCTIONS CALLS  -----------------------------------

		//Create a payment into the blockchain
		func blockchainStartNewPayment(price int64, externalPayerId int64, externalReceiverId int64) (*types.Transaction, error){
			//0xfb7974b7616a0c0dbd08c7fee7f1291548045e33 is the contract's address gicen by the deploySmartContract function in blockchain_main.go
			contract, err := contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
			if err != nil {
				fmt.Printf("Unable to bind to deployed instance of contract:%v", err)
			}
	
			payment, err := contract.StartNewPayment(auth, big.NewInt(price), big.NewInt(externalPayerId), big.NewInt(externalReceiverId))
			if err != nil {
				fmt.Println("Unable to bind to deployed instance of contract:%v :%v", err, payment)
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
