pragma  solidity ^0.4.20;

contract array1Test{
    
    uint len1 = 0;
    function func(){
        uint[] memory value = new uint[](10);
        
        value[0] = 1;
        // value.push(1); //不支持
        len1 = value.length;
        // value.length = 20; // 长度不能修改
        
        
    }
    
    uint [] public arrayValue;
    uint public len2;
    function func2(){
        arrayValue = new uint[](10);
        arrayValue[0] = 99;
        arrayValue.push(100);
        
        len1 = arrayValue.length;
        arrayValue.length = 50;
        len2 = arrayValue.length;
    }
}

// 直接创建
contract array2Test{
    uint [7] value = [1,2,3,4,5,6,7];
    // 内容可变　长度不可变
    uint public sum;
    function func(){
        value[0] = 10;
        uint a = value[1];
        
        for (uint i =0; i < value.length ; i++){
            sum += value[i];
        }
        
        //value.push(8);
    }
}


