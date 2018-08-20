pragma solidity ^0.4.20;

contract HashAnOwner{
    address public owner;
    uint public a;
    
    constructor() public {
        owner = msg.sender;
    }
    
    function userSuperPower() public {
        require(msg.sender == owner);
        //if (msg.sender != owner){
        //    throw;
        //}
        a = 10;
    }
}