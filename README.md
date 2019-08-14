# BLOCKCHAIN’S IMPLEMENTATION, AND FURTHER STEPS


The project is composed of two parts : 
- The smart contract, coded in solidity and composed, after compilation, of golang files (/smartContract)
- The api in golang (/connectAPI)


After developing the smart contract in solidity, and testing its function on remix, I can compile it inside my project. I first put my solidity code in a file of my text editor and call it Contract.sol. 
In my editor terminal :
```
solc --abi Contract.sol -o build
```

This will create one .abi file per contract in my solidity file (for instance in this project there are 3 : owned, ImoDollar, Payments)
Then I want to create the golang file that will contains all my smart contract’s function, and link it to my abi file :
```
abigen --abi=./build/Contract_sol_ImoDollar.abi --pkg=ImoDollar --out=ImoDollar.go
```

note : This is only for the ImoDollar file, I want to do it for every of my .abi file.
Let’s create the .bin file :
```
solc --bin Contract.sol -o build
```

And connect it to the other files :
```
abigen --bin=./build/ Contract_sol_ImoDollar.bin --abi=./build/ Contract_sol_ImoDollar.abi --pkg=ImoDollar --out=ImoDollar.go
```

note : Once again this is only for the ImoDollar file. I have to connect every of my .bin files
Once I have executed all these steps I can use my smart contract using the golang API

In the next part I will Describe the API in golang.

# Deployment of the smart contract
First the we need to deploy the contract, by calling deploySmartContract() function (blockchainAPI/blockchain_main.go). This function is called everytime the app is opened.
This function returns an address that we need every time we call a smart contract’s function. We need to enter it in the function :
```
contract, err := contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
```
Where "0xfb7974b7616a0c0dbd08c7fee7f1291548045e33" is the address



# Create a user’s keystore
The keystore is a file containing all the information of the user’s wallet/ethereum account, which he uses complete transactions.
The function ```createKs(userName string, clientPassword string)``` Permits to auto generate a keystore and create a wallet for a new user. The user is automaticaly saved in the blockchain. ClientPassword is the password the user wants in order to delock its wallet. The wallet file is created in a /wallet folder.
This function should return an integer, which is the address of the user in the blockchain (called clientId in the smart contract), that we will use to import this user or use its data in the smart contract. However it only returns an hash for now.

Implementation: When the createKs function works correctly, we will want to call it every time a new user is created. We need to save the returned integer, userId, of this function in the client’s data because we want to be able to access this integer easily every time the client connects in order to access his wallet.

# Import a user’s keystore
To import a user’s keystore we need to call ```importPrivateKey(userId int64, password string)``` function, with the user ID (the integer we spoke above), and the password created when calling the createKs function.
Implementation : Every time we call a smart contract’s funtion, one of the argument will be the clientId. We need to be able to access it while the client is connected

# Smart contract functions
To call a smart contract’s function we need to go to /blockchain_transactions.go.

Here is an example:
```
func blockchainStartNewPayment(price int64, externalPayerId int64, externalReceiverId int64) (*types.Transaction, error){
	contract, err := contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
	if err != nil {
		fmt.Printf("Unable to bind to deployed instance of contract:%v", err)
	}
	payment, er := contract.StartNewPayment(auth, big.NewInt(price), big.NewInt(externalPayerId), big.NewInt(externalReceiverId))
	if err != nil {
		fmt.Println("Unable to bind to deployed instance of contract:%v :%v", er, hello)
	}
	fmt.Printf("new payment :%v", hello)
	return payment, er
}
```

We first use  the function ```NewPayments(address, blockchain)``` to connect the user’s wallet to the blockchain. Here, the 0xfb7974b7616a0c0dbd08c7fee7f1291548045e33 address is the smart contract's address, given when ```deploySmartContract``` function is called in /blockchain_main.go .
```
contracts.NewPayments(common.HexToAddress("0xfb7974b7616a0c0dbd08c7fee7f1291548045e33"), blockchain)
```
Blockchain was previously defined to connect to an ethereum node with  
```
blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/44029fa59f854e649e0a4e9f3691de56")
```

then, we call the startNewPayment function of the smart contract 
```
payment, er := contract.StartNewPayment(auth, big.NewInt(price), big.NewInt(externalPayerId), big.NewInt(externalReceiverId))
```
auth was previously declared to link to the user’s wallet
```
auth, err := bind.NewTransactor(strings.NewReader(key), password)
```
auth contains the private key of the wallet, it is only used for function that modify the blockchain and so has a transaction fee. Function that doesn’t have a fee only use &bind.CallOpts{Pending: true} instead

The client and payment functions :
```
aNewClient (string username, string file, string clientAccountAddress) returns (int userId)
```
with file = the path of the keystore file of the client
```
checkIfClientExists (int clientId) returns (bool)
```
```
getAClient (int clientId) returns (string username, string file, string clientAccountAddress)
```
```
startNewPayment (int price, int payerId, int receiverId) returns (int paymentId)
```
withpayerId and receiverId are the clientId variables of the client who has to pay the payment and the person who shall receive the money.
```
getPaymentById (int paymentId) returns (int price, int payerId, int receiverId, int status)
```
```
getPaymentByOwner (int ownerId) returns (int[] paidPayments, int[] receivedPayments)
```
which returns a tab of the paymentId variable : all the ids of the payments that the owner paid (paidPayments) or received (receivedPayment)
```
payPayment (int externalPaymentId) 
```
this function pays the payment online with the token we have created as the cryptocurrency for imobly. The user that uses this function can pay all the payments he wants, but he will automatically pay with its own wallet.

The token functions:
A token is like a cryptocurrency, only that it is not native on an app : it uses the native cryptocurrency in  order to by, sell and make exchanges.
```
transferFrom (address _from, address _to, int amount)
```
Transfer tokens from a user to another. Used to pay a payment.
```
transfer (address _to, int amount)
```
In case we are sure that the one that calls the function is the one that shall pay the payment
```
setPrice (int newSellPrice)
```
The function that sets how much costs an imobly token compared to the native cryptocurrency (ethers) : equal the price we buy/sell each token in ethers.
```
Buy ()
```
To buy tokens with ethreum. This is a “payable” functions, which means we can send ethers to the smart contract. The amount of ethers sent to the contract is how much we want to buy
```
Sell (int amount)
```
To sell tokens and get ethers


# TODO
- Connect every function in the smart contract (fast : see in connectAPI/blockchain_transaction.go how to connect to the functions in go, always the same)
- createKs function should return an int (userId) and not an hash of the transaction
- unmarshallJSON function shall work in order to use importPrivateKey function. Indeed when we create the client in the blockchain, createKs always adds "keystore://" before the file path string. This provokes a mistake. Thoses caracters need to be deleted, which is the function of unmarshallJSON (more specified inside the code)
- If the user has its own ethereum account : create a function that saves its keystore and will be able to use it.
- How to add safely ETHERS in the client’s wallet?
- check if the buy/sell token function works well (some problems during testing but I think it is due to the remix testing environment)
- Configure a quorum node so that we can run this smart contract on a private blockchain and so delete the transaction cost
