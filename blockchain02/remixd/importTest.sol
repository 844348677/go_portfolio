pragma  solidity ^0.4.20;

// import 分别deploy　hello 和　importTest
import "./hello.sol";
// import * as symbolName from "filename"

contract ImportTest{
    
    function setString(Test test, string str) public {
        test.setValue(str);
    }
}


