pragma solidity ^0.4.20;

contract deleteTest {
    string public str1 = "hello";
    
    // delete 操作符可以用于　任何变量(mapping除外), 将其设置成默认值
    // delete 删除　设置成默认值
    function delStr() public {
        delete str1;
    }
    
    function setStr() public{
        str1 = "HELLO";
    }
    
    // 静态数组　，　动态数组
    uint[10] public staticArray = [4,2,4,5,6,7,8];
    uint[] public dynamicArray = new uint[](10);
    
    function intDynamicArray () public {
        for (uint i=0; i<10 ; i++){
            dynamicArray.push(i);
        }
    }
    function delStaticArray() public{
        delete staticArray;
    }
    function delDynamicArray() public{
        delete dynamicArray;
    }
    function getArrayLength() view public returns (uint,uint){
        return (staticArray.length, dynamicArray.length);
    }
    
    // 动态数组　整个元素个数清零　长度归０
    // 静态数组　个数不变　每个元素赋值为０
    
    mapping(uint => bool) public map1;
    
    function initMap() public{
        map1[1] = true;
        map1[2] = true;
        map1[3] = false;
        // delete map1;
    }
    //delete map1;
    function deleMapByKey(uint key) public{
        delete map1[key];
    }
    
    struct Person{
        string name;
        mapping(string => uint) nameScore;
    }
    
    Person public p1;
    
    function initP1(){
        p1.name = "duke";
        p1.nameScore["duke"] = 80;
    }
    
    function returnP1() view public returns (string,uint){
        return (p1.name, p1.nameScore["duke"]);
    }
    
    // name会被清空　map没有表现
    function deleteP1() public{
        delete p1;
    }
    
}