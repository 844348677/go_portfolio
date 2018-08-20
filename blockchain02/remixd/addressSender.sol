pragma  solidity ^0.4.20;

contract senderTest{
    address public _owner;
    
    function senderTest() public {
        _owner = msg.sender;
    }
    
    function getOwnerBalance() public returns (uint256){
        
        //return _owner.balance;
        return msg.sender.balance;
    }
    
    // 调用者的　地址
    function getInvoker() constant public returns (address){
        return msg.sender;
    }
}