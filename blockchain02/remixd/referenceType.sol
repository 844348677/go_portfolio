pragma  solidity ^0.4.20;

contract Student{
    string _name = "lily";
    uint _num;
    
    function execute() public{
        changeName(_name);
    }
//  function changeName(string memory name) public{  
    function changeName(string storage name) private{ // internal
        _num = 10;
        bytes(name)[0] = "L";
    }
    
    function getName() constant public returns (string){
        return _name;
    }
    
    function getNum() constant  public returns (uint){
        return _num;
    }
}

// memory(Value Type)  storage( Reference Type)