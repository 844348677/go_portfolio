pragma  solidity ^0.4.20;

contract C {
    bytes public _name = new bytes(1);
    bytes public _name2;
    
    // 后续修改长度
    function setLength(uint length){
        _name.length = length;
    }
    function getLength() constant returns (uint){
        return _name.length;
    }
    // 直接写数据
    function setName(bytes name){
        _name = name;
    }
    function changeName(bytes1 name){
        _name[0] = name;
    }
    function setInside(){
        _name = "helloworld";
        _name2 = "HELLOWORLD2";
    }
}