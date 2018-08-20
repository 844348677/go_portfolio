pragma  solidity ^0.4.20;

contract Test {
    struct Student{
        string name;
        uint age;
        uint score;
        string sex;
    }
    
    Student public stu1 = Student("lily",18,90,"girl");
    Student public stu2 = Student({name:"Jim",age:20,score:80,sex:"boy"});
    
    Student [] public Students;
    
    function assign() public{
        Students.push(stu1);
        Students.push(stu2);
        
        stu1.name = "Lily";
        
    }
}