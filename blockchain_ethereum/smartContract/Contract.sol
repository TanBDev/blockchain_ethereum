pragma solidity ^0.5.1;

//import "github.com/Arachnid/solidity-stringutils/strings.sol";

contract owned {
    address public owner;
    constructor() public {
        owner = msg.sender;
    }
    
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }
}

//Contract creates the token
contract ImoDollar is owned { 
    
    string public name; 
    string public symbol; 
    uint8 public decimals;
    uint256 public totalSupply;
    address devAddress;
    uint256 public sellPrice;

    // Balances for each account
    mapping(address => uint256) balances;
    // Owner of account approves the transfer of an amount to another account
    mapping(address => mapping (address => uint256)) allowed;
    
    // Events
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    event Transfer(address indexed from, address indexed to, uint256 value);


    
    
    // This is the constructor and automatically runs when the smart contract is uploaded    
/* 
    function imoDollar() public { // Set the constructor to the same name as the contract name
        name = "Imodollar"; // set the token name here
        symbol = "ImD"; // set the Symbol here
        decimals = 18; // set the number of decimals
        devAddress=0x878307A585Df983D66bA8b310341dD6cB11a6191; // Add the address that you will distribute tokens from here
        uint initialBalance=1000000000000000000*1000000; // 1M tokens
        balances[devAddress]=initialBalance;
        totalSupply+=initialBalance; // Set the total suppy
    }
    */
    
    // Sets basic datas of smart contract
    constructor(
        string memory tokenName, 
        string memory tokenSymbol, 
        uint256 tokenInitialBalance,
        address tokenDevAddress,
        uint setSellPrice) public {
        name = tokenName;                                   // Set the name for display purposes
        symbol = tokenSymbol;                               // Set the symbol for display purposes
        devAddress= tokenDevAddress; 
        uint initialBalance= tokenInitialBalance; 
        balances[devAddress]=initialBalance;
        totalSupply+=initialBalance; // Set the total suppy
        sellPrice = setSellPrice;
    }

    
    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }
    
    
    
    function transferFrom (address _from, address _to, uint256 _amount) private {
        require (balances[_from] >= _amount);
        //require (allowed[_from][msg.sender] >= _amount);
        require (_amount > 0);
        require (balances[_to] + _amount > balances[_to]);
        balances[_from] -= _amount;
        allowed[_from][msg.sender] -= _amount;
        balances[_to] += _amount;
        emit Transfer(_from, _to, _amount);
    }
    
    
    
    //Send only from the accounts using the function
    function transfer(address _to, uint256 _value) public returns (bool success) {
        transferFrom(msg.sender, _to, _value);
        return true;
    }
    

    // Allow _spender to withdraw from your account, multiple times, up to the _value amount.
    // If this function is called again it overwrites the current allowance with _value.
    function approve(address _spender, uint256 _amount) public returns (bool success) {
        allowed[msg.sender][_spender] = _amount;
        emit Approval(msg.sender, _spender, _amount);
        return true;
    }
    
    /// @notice Allow users to buy tokens for `newBuyPrice` eth and sell tokens for `newSellPrice` eth
    /// @param newSellPrice Price the users can sell to the contract
    function setPrices(uint256 newSellPrice) onlyOwner public {
        sellPrice = newSellPrice;
    }
    
    function multiply(uint x, uint y) internal pure returns (uint z) {
        require(y == 0 || (z = x * y) / y == x);
    }

/// @notice Buy tokens from contract by sending ether
    function buy() public payable {
        uint _amount = msg.value*sellPrice;
        //require(_amount == msg.value);
        //require(imbAmount > 0);
        transfer(address(this),  _amount);              // makes the transfers
    }


  
  
    
/// @notice Sell `amount` tokens tsello contract
    /// @param amount amount of tokens to be sold
    function sell (uint256 amount) public {
        address myAddress = address(this);
        require(myAddress.balance >= amount / sellPrice);      // checks if the contract has enough ether to buy
        transferFrom(msg.sender, address(this), amount);              // makes the transfers
        msg.sender.transfer(amount * sellPrice);          // sends ether to the seller. It's important to do this last to avoid recursion attacks
    }
}


contract Payments is ImoDollar {

    constructor(
        string memory tokenName, 
        string memory tokenSymbol, 
        uint256 tokenInitialBalance,
        address tokenDevAddress,
        uint setSellPrice) ImoDollar(tokenName, tokenSymbol, tokenInitialBalance, tokenDevAddress, setSellPrice) public {}
    
    
    // payment register
    struct Payment {
        uint price;
        uint receiverId;
        uint payerId; //external database payer id
        uint status;   // status = 0 : doesn't exist , 1 : unpaid , 2 : paid 
    }
    
    struct Client{
        string userName;
        string file;
        string clientAccountAddress;
    }

    Payment[] public payments;
    Client[] public clients;
    
    // mappings 
    mapping (uint => address) paymentsMap; //external payment alphanumeric id (gateway) => payment
    mapping (address => uint) ownerPaymentCount;
    mapping (uint => address) clientsMap;

    // events
    event NewPaymentEvent(uint externalPaymentId, uint price, uint externalPayerId, uint receiverId, uint status);
    event NewClientEvent(uint externalClientId, string userName, string file, string clientAccountAddress);
    


    //Need to convert a string to an address
    function parseAddr(string memory _a) internal pure returns (address _parsedAddress) {
        bytes memory tmp = bytes(_a);
        uint160 iaddr = 0;
        uint160 b1;
        uint160 b2;
        for (uint i = 2; i < 2 + 2 * 20; i += 2) {
            iaddr *= 256;
            b1 = uint160(uint8(tmp[i]));
            b2 = uint160(uint8(tmp[i + 1]));
            if ((b1 >= 97) && (b1 <= 102)) {
                b1 -= 87;
            } else if ((b1 >= 65) && (b1 <= 70)) {
                b1 -= 55;
            } else if ((b1 >= 48) && (b1 <= 57)) {
                b1 -= 48;
            }
            if ((b2 >= 97) && (b2 <= 102)) {
                b2 -= 87;
            } else if ((b2 >= 65) && (b2 <= 70)) {
                b2 -= 55;
            } else if ((b2 >= 48) && (b2 <= 57)) {
                b2 -= 48;
            }
            iaddr += (b1 * 16 + b2);
        }
        return address(iaddr);
    }


    function aNewClient(string memory userName, string memory file, string memory clientAccountAddress) public returns (uint){
        uint externalClientId = clients.push(Client(userName, file, clientAccountAddress));
        emit NewClientEvent(externalClientId, userName, file, clientAccountAddress);
        return externalClientId;
    }
    
    function checkIfClientExists(uint clientId) public returns (bool){
        return parseAddr(clients[clientId].clientAccountAddress) != parseAddr("0x0000000000000000000000000000000000000000");
    } 

    function getAClient(uint clientId) public view returns (string memory, string memory, string memory){
        //require(checkIfClientExists(clientId) == true);
        string memory userName = clients[clientId].userName;
        string memory file = clients[clientId].file;
        string memory clientAccountAddress = clients[clientId].clientAccountAddress;
        return (userName, file, clientAccountAddress);
    } 
    
    function startNewPayment(uint price, uint externalPayerId, uint externalReceiverId) public returns(uint) {
        require(price > 0);
        uint externalPaymentId = payments.push(Payment(price, externalReceiverId, externalPayerId, 1));
        //Payment storage myPayment = payments[externalPaymentId];
        address payerDirection = clientsMap[externalPayerId];
        address receiverDirection = clientsMap[externalReceiverId];
        ownerPaymentCount[payerDirection]++;
        ownerPaymentCount[receiverDirection]++;
        emit NewPaymentEvent(externalPaymentId, price, externalPayerId, externalReceiverId, 1);
        return externalPaymentId;
    }
    


    function getPaymentById(uint externalPaymentId) public view returns(uint, uint, uint, uint) {
        uint price = payments[externalPaymentId].price;
        uint payerId = payments[externalPaymentId].payerId;
        uint receiverId = payments[externalPaymentId].receiverId;
        uint status = payments[externalPaymentId].status;
        return (price, payerId, receiverId, status);
    }

    // Returns two tabs of the payments' addresses inside the blockchain
    //first tab are the addresses of the transactions that the owner paid
    //second one the addresses of the payments the owner received
    function getPaymentsByOwner(uint externalOwnerId) public view returns(uint[] memory, uint[] memory) {
        address ownerDirection = clientsMap[externalOwnerId];
        uint[] memory paidTransactions = new uint[](ownerPaymentCount[ownerDirection]);
        uint[] memory receivedTransactions = new uint[](ownerPaymentCount[ownerDirection]);
        uint counterPaid = 0;
        uint counterReceived = 0;
        for (uint i = 0; i < payments.length; i++) {
            if (payments[i].payerId == externalOwnerId) {
                paidTransactions[counterPaid] = i;
                counterPaid++;
            }
            if (payments[i].receiverId == externalOwnerId) {
                receivedTransactions[counterReceived] = i;
                counterReceived++;
          }
        }
    return (paidTransactions, receivedTransactions);
  }
  
  //Pay payment using the tokens of the owner's wallet
  function payPayment(uint externalPaymentId) public{
      uint price = payments[externalPaymentId].price;
      uint receiverId = payments[externalPaymentId].receiverId;
      address receiverAccountAddress = parseAddr(clients[receiverId].clientAccountAddress);
      require(receiverAccountAddress != address(owner));
      transfer(receiverAccountAddress, price);
      payments[externalPaymentId].status = 2;
  }
  
}