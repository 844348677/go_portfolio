pragma  solidity ^0.4.20;

contract Test{
    
    string str1 = "hello!";
    function loopTest() public{
        // 死循环　uint8 : 0 255之间循环
        // 坑
        for (var i=0; i<2000; i++){
            
        }
        var str2 = str1;
    }
    
    
}