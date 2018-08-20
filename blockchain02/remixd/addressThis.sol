pragma  solidity ^0.4.20;

contract addressThis {
    
    uint _money ;
    
    //  构造函数, 接受money
    function addressThis() payable public {
        _money = msg.value;
    }
    
    function getThis() public returns (address){
        return this;
    }
    
    function getBalance() constant public returns (uint256){
        return this.balance;
    }
    
    function getMoney() constant public returns(uint){
        return _money;
    }
    
    function getMsgSender() public returns (address){
        return msg.sender;
    }
}