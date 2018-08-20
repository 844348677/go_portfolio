pragma solidity ^0.4.20;

contract Test {
    uint public v1 = 10;
    // 常亮不能修改　　值类型　string
    uint constant v2 = 10;
    
    string str1 = "hello!";
    string constant str2 = "test!";
    
    
    function f1() public {
        v1 = 20;
        // v2 = 30;
        
        str1 = "Hello!";
        // str2 = "Test!";
        
    }
    struct Person{
        string name;
        uint age;
    }
    
    // 不能对结构类型使用 constant
    // Person constant p1 ;
    // constant 修饰的值类型，无法　被修改
    // constant 仅可以修饰值类型　无法修饰引用类型　(string 除外)
    
    // 不会修改里面的值　
    // constant 修饰的函数内, 如果修改了状态变量 ,那么状态变量的值是无法改变的，小心
    function f2() constant public{
        v1 = 20;
    }
    
    // view 关键字　和 constant 一样的　
    // view 不能修饰变量　，　只能修饰函数　只能对storage类型的变量读取，无法修改
    
    // pure 不能读也不能写
    function f3() pure public returns (uint){
        //v1 = 29;
        //return v1;
    }
    // pure 　只能修饰函数
    // 它表明该函数内, 无法读写状态变量
    
}