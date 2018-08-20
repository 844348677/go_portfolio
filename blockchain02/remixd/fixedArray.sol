pragma  solidity ^0.4.20;

// 内置固定等长数组
contract fixedArray{
    // bytes1 ... bytes2
    
    bytes1 b1 = "x";
    bytes2 b2 = "xy";
    
    bytes3 public b3 = "xy";
    //     0: bytes3: 0x787900
    
    uint public len = b3.length;
    
    // 长度不可修改
    
    bytes8 b8 = "12345678";
    bytes1 public b8_1 = b8[0];
    //     0: bytes1: 0x31 ; ASCII
    
    // b8[1] = "0";
}

contract C {
    bytes10 public b = "helloworld";
    uint public len;
    function f() public{
        len = b.length;
        
        // b.length = 10; // error,　长度为只读，不可修改
    }
    bytes1 public c = b[0];
    
    // b = "HELLO"; // error, 定义之后不可修改
}


