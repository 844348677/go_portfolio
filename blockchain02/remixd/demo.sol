// 定义合约
// 向上兼容
pragma solidity ^0.4.21;

// 定义合约
contract SimpleStorage{
    //状态变量　属于成员变量
    uint storedData;
    
    //构造函数
    // function SimpleStorage(){
    constructor() public {
        storedData = 100;
    }
    
    //成员函数
    function set(uint x) public {
        storedData = x;
    }
    
    function get() constant public returns (uint retVal){
        return storedData;
    }
    
    //析构函数
    function destory() public {
        selfdestruct(msg.sender);
    }
}