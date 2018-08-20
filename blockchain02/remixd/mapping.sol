pragma  solidity ^0.4.20;

contract test{
    //id => name
    
    mapping(uint => string) id_names;
    
    constructor() public {
        id_names[0x001] = "lily";
        id_names[0x002] = "Jim";
        // 会覆盖上面的
        id_names[0x002] = "lLily";
    }
    
    function getNameById(uint id) constant public returns (string){
        string storage name = id_names[id];
        return name;
    }
}