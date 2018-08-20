pragma  solidity ^0.4.20;

contract Test{
    bytes10 b10 = 0x68656c6c6f776f726c64;     // helloworld
    // bytes bs10 = b10;
    
    // copy 赋值 
    bytes public bs10 = new bytes(b10.length);
    // 1 固定橙子姐数据转动态字节数据
    function fixedBytesToBytes(){
        for(uint i =0;i<b10.length;i++){
            bs10[i] = b10[i];
        }
    }
    
    //2 . 动态字节数组转string
    string greeting = "helloworld"; 
    bytes public b1;
    // 将 greeting 赋值给 b1
    function StringToBytes(){
        b1 = bytes(greeting);
    }
    
    // 3 . bytes 转成string
    string public str3;
    function BytesToString() public{
        str3 = string(bs10);
    }
    
    // 固定长度数组　-> 动态长度数组　<- -> string
    // bytes1, bytes2   , bytes , string
    
    // 4 . fixed bytes  to string
    function FixedBytesToString(){
        // 不能强转
        //string tmp = string(b10);
    }
    
}