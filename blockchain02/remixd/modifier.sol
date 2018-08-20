pragma solidity ^0.4.20;

contract HasAnOwner{
    address public owner;
    uint public a;
    
    constructor() public{
        owner = msg.sender;
    }
    
    modifier ownerOnly(address addr){
        require(msg.sender == owner);
        
        // 下划线　代码修饰器　所修饰函数的代码
        _;
        // 例如　代表　a=10;
    }
    
    // 业务代码和　和检验分离
    
    function useSuperPowers() ownerOnly(msg.sender) public{
        
        a = 10;
    }
}