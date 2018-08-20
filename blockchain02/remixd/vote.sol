pragma solidity ^0.4.20;
pragma experimental ABIEncoderV2;

// ["aaa","bbb","ccc","ddd"] deploy

contract VoteContract{
    
    // 投票者
    struct Voter{
        // 投票号
        uint voteNumber;
        // 是否投过票，已经投过的不允许再次投票
        bool isVoted;
        // 投票权重　初始时均为１，被委托后，权重增加
        uint weight;
        address delegate; // 自定代理人
    }
    
    struct Candidate{
        // 候选人名字
        string name;
        // 获得的投票数
        uint voteCount;
    }
    
    //负责创建合约，授权给地址，使之成为投票人
    address public admin;
    
    // 候选人集合
    Candidate[] public candidates;
    //　投票人集合
    mapping(address => Voter) public voters;
    
    // ["aaa","bbb","ccc","ddd"]
    constructor(string[] candidatesNames) public{
        admin = msg.sender;
        
        // 没一个名字都生成一个候选人　，　添加到候选人集合中
        for(uint i=0; i<candidatesNames.length; i++){
            Candidate memory tmp = Candidate({name:candidatesNames[i],voteCount:0});
            candidates.push(tmp);
        }
    }
    
    // 限定只有管理员有添加投票人的权利
    modifier adminOnly(){
        require(admin == msg.sender);
        _;
    }
    
    // 添加投票人 只能管理员给其他人授权，不能自己给自己授权，需要切换
    function giveVoteRightTo(address addr) adminOnly public {
        if (voters[addr].weight >0)
            revert();
        // 授予　投票者　投票权利
        voters[addr].weight = 1;
    }
    
    // 投票
    function vote(uint voteNum) public{
        //Voter memory voter = voters[msg.sender];
        Voter storage voter = voters[msg.sender];
        // memory 复制值　　storage复制地址
        
        // 不是候选人或已经投过票的　直接返回
        if(voter.weight <= 0 || voter.isVoted)
            revert();
        voter.isVoted = true;
        voter.voteNumber = voteNum;
        
        // 候选人的编号和存储位置一致，第０个候选人存储在第０个位置
        candidates[voteNum].voteCount += voter.weight;
    }
    
    // 代理人
    function delegateFunc(address to) public{
        Voter storage voter = voters[msg.sender];
    
        // 不是候选人或已经投过票的　直接返回
        if(voter.weight <= 0 || voter.isVoted)
            revert();
        
        // 没有被人代理　，　同时不能指会自己
        while (voters[to].delegate != address(0) && voters[to].delegate != msg.sender){
            to = voters[to].delegate;
        }
        // 如果代理人也制定了代理人，那么需要轮询找到最终代理人，并且最终代理人不能是自己，否则死循环
        
        // 代理人是自己　退出
        require(msg.sender != to);
        // 等同
        if (msg.sender == to){
            revert();
        }
        
        voter.isVoted = true;
        voter.delegate = to;
        
        Voter storage finalDelegateVoter = voters[to];
        
        if(finalDelegateVoter.isVoted){
            // 如果代理人已经投票，那么在代理人投票的候选人票数上加上自己的权重
            candidates[finalDelegateVoter.voteNumber].voteCount += voter.weight;
        } else{
            //否则，在代理人权重上加上自己的权重
            finalDelegateVoter.weight += voter.weight;
        }
    }
    
    function whoWin() view public returns (string, uint) {
        string winner;
        uint winnerVoteCount;
        
        for (uint i=0; i<candidates.length; i++){
            if (candidates[i].voteCount > winnerVoteCount){
                winnerVoteCount = candidates[i].voteCount;
                winner = candidates[i].name;
            }
        }
        
        return (winner,winnerVoteCount);
    }
    
}