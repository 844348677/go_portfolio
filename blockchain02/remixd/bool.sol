pragma  solidity ^0.4.20;

contract BoolTest{
    uint v1 = 10;
    uint v2 = 20;
    
    string s1 = "hello";
    string s2 = "world";
    
    bool flag1 = true;
    bool flag2 = false;
    
    function f1() constant public returns (bool){
        return !flag1;
    }
    
    function f2() constant public returns (bool){
        return flag1 && flag2;
    }

    function f3() constant public returns (bool){
        return flag1 || flag2;
    }

    function f4() constant public returns (bool){
        return v1 == v2;
    }
    
    function f5() constant public returns (bool){
        return v1 != v2;
    }
    
}