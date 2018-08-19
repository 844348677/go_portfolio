pragma solidity ^0.4.21;

contract Test{
    string str1;
    
    function setValue (string para) public{
        str1 = para;
    }
    
    function getValue() constant public returns (string){
        return str1;
    }
}