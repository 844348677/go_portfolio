pragma  solidity ^0.4.20;

contract IntegerTest {
    uint8 _v1 = 10;
    // uint8 0 - 255
    //  1111 1111 
    // 10000 0000 - 1
    // int8 -127 - 127
    //  1111 1111 
    // 10000 0000-1 -> -(10000 0000 - 1)
    //  0111 1111
    //  1000 0000 - 1
    
    // 数值越界　值不可控　不报错
    function IntegerTest(uint8 para) public{
        _v1 = para;
    }
    
    function getValue() constant public returns(uint8){
        return _v1;
    }
    
    function com1() constant public returns (bool){
        return _v1 >= 10;
    }
    // & , | , ^ , ~ , 与　或　异或　位取反
    // 10  0000 1010
    // 4   0000 0100
    // &   0000 0000
    // |   0000 1110
    // ^   0000 1110
    
    // 3   0000 1100
    // ~3  1111 0011
    
    uint8 a = 10;
    uint8 b = 4;
    uint8 c = 3;
    
    function f1() constant public returns (bool){
        return a & b == 0;
    }
    
    function f2() constant public returns (bool){
        return a | b == 14;
    }
    function f3() constant public returns (bool){
        return a ^ b == 14;
    }
    function f4() constant public returns (bool){
        return ~c == 252;
    }
}