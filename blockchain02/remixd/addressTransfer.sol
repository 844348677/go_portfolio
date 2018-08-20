pragma  solidity ^0.4.20;

// account 1 0x6000d643e3fb9b1fa180b19f5f5d33d2b646109a : int112
// account 4 0xb685ed019fc04b40454c56db1c2102083a5b47dc


// 这个转账有问题　搞死我了　！！！！
contract TransferTest {
    address AliceAddress = 0xb685ed019fc04b40454c56db1c2102083a5b47dc;
    
    function transfer() {
        AliceAddress.transfer(msg.value);
    }
}