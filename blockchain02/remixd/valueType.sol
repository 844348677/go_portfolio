pragma  solidity ^0.4.20;

contract Student{
    uint _score =50;
    uint _num = 0;
    
    function execute() public{
        changeScore(_score);
    }
    
    // 没有被修改，　copy
    function changeScore(uint score) private{
        score = 90;
        _num = 90;
    }
    
    function getScore() constant public returns (uint) {
        return _score;
    }
    function getNum() constant public returns (uint){
        return _num;
    }
}