pragma  solidity ^0.4.20;

contract FunctionTest{
    uint public v1;
    uint public v2;
    // function v1() constant returns (uint) {
        
    
    
    // 用 internal 修饰
    function internalFunc() internal{
        v1 = 10;
    }
    // 通过　internal 和　external 修饰访问权限
    function externalFunc() external returns (uint res){
        v2 = 20;
        return v2;
    }
    
    function resetV2() public{
        v2 = 0;
    }
    
    function callFunc(){
        // 直接使用内部的方式调用
        internalFunc();
        
        // 不能在内部调用一个外部函数　编译错误
        //externalFunc();
        
        // this 调用相当于外部调用
        // this.internalFunc();
        
        this.externalFunc();
    }
}

// 继承
contract Son is FunctionTest{
    function callInternalFunc() public{
        internalFunc();
        //externalFunc();
    }
}

//　合约本身也是地址
contract FunctionTest1{
    uint public v3;
    function externalCall(FunctionTest ft){
        // 调用另一个合约的外部函数
        v3 = ft.externalFunc();
        
        // 不能调用另一个合约的内部函数
        // ft.internalFunc()
    }
}




