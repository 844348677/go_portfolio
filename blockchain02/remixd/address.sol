pragma  solidity ^0.4.20;

// 以太坊地址的长度，大小２０个字节，１６０位，可以用一个uint160编码
// 属性　balance
//　方法　send() transfer() call() delegatecall() callcode()

contract AddressTest{
    address _add1 = 0xca35b7d915458ef540ade6068dfe2f44e8fa733c;
    address _add2 = 0x14723a09acff6d2a60dcdf7aa4aff308fddc160c;
    uint160 public _num = 0;
    address public _add;
    // 自动生成请求方法　
    // function _num() constant public returns
    
    function address2uint160() public returns(uint160){
        _num = uint160(_add1);
        return _num;
    }
    
    function uint160ToAddress() public returns(address){
        _add = address(_num);
        return _add;
    }
    
    function comAddress() constant public returns(bool){
        return _add == _add2;
    }
    
    function isGreater() constant public returns(bool){
        return _add1 > _add2;
    }
}

// 1154414090619811796818182302139415280051214250812
// 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c