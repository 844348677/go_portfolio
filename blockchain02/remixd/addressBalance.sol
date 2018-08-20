pragma  solidity ^0.4.20;

contract addressBalance{
    
    function getBalance(address addr) constant public returns (uint){
        return addr.balance;
        
    }
}