pragma solidity ^0.4.21;

contract EthUnit{
    uint a = 1 ether;
    uint b = 10 ** 18 wei;
    uint c = 1000 finney;
    uint d = 1000000 szabo;
    
    // 货币单位　时间单位
    function f1() constant public returns (bool){
        return a == b;
    }
    function f2() constant public returns (bool){
        return a == c;
    }
    function f3() constant public returns (bool){
        return a == d;
    }
    function f4() pure public returns (bool){
        return 1 ether == 100 wei;
    }
}

// 时间单位
contract TimeUnit{
    function f1() pure public returns (bool){
        return 1 == 1 seconds;
    }
    function f2() pure public returns (bool){
        return 1 minutes == 60 seconds;
    }
    function f3() pure public returns (bool){
        return 1 hours == 60 minutes;
    }
    function f4() pure public returns (bool){
        return 1 days == 24 hours;
    }
    function f5() pure public returns (bool){
        return 1 weeks == 7 days;
    }
    function f6() pure public returns (bool){
        return 1 years == 365 days;
    }
}
