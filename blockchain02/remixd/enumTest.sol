pragma solidity ^0.4.21;

contract enumTest{
    // uint8 数字增长　动态调整
    enum ActionChoices { GoLeft, GoRight, GoStraight, SitStill }
    // 错误码
    ActionChoices _choice;
    ActionChoices defaultChoice = ActionChoices.GoStraight;
    
    function setGoStraight(ActionChoices choice) public{
        _choice = choice;
    }
    function getChoice() constant public returns (ActionChoices){
        return _choice;
    }
    function getDefaultChoice() constant public returns (uint){
        return uint(defaultChoice);
    }
    
    function isGoLeft(ActionChoices choice) constant public returns (bool){
        if (choice == ActionChoices.GoLeft){
            return true;
        }
        return false;
    }
}