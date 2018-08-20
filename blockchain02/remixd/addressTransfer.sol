pragma  solidity ^0.4.20;

// account 1 0x6000d643e3fb9b1fa180b19f5f5d33d2b646109a : int112
// account 4 0xb685ed019fc04b40454c56db1c2102083a5b47dc


// deploy 的时候　value　填空
// 构造的不用传money过去
// deploy 完毕的时候调用transfer的方法　填写value
contract TransferTest {
    address AliceAddress = 0xb685ed019fc04b40454c56db1c2102083a5b47dc;
    
    function transfer() {
        AliceAddress.transfer(msg.value);
    }
}

// send transfer
// send 不会抛出异常　
// send 相对于　transfer 语言更底层

// send 执行有一些风险　　如果调用栈的　深度或者超过１０２４或ｇａｓ耗光，交易都会失败





