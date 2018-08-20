pragma  solidity ^0.4.20;

contract C {
    string public _name = "lily";
    
    function nameBytes() constant returns (bytes){
        return bytes(_name);
    }
    
    function nameLength() constant returns (uint){
        // uint len = _name.length;
        return bytes(_name).length;
    }
    function changeName() public{
        bytes(_name)[0] = "L";
        
        // s2[0] = "H"; // error 不支持ｘｉａ
    }
    
    function changeLength(){
        bytes(_name).length = 15;
        bytes(_name)[14] = "x";
        
        // 不会自动追加　越界报错
        // bytes(_name)[19] = "x";
    }
}