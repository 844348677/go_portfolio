package data

import (
	"gopkg.in/mgo.v2"
	"time"
)

// MongoStore is a MongoDB data store which implements the Store interface
type MongoStore struct {
	session *mgo.Session
}
// 又是在哪找的　包
// https://godoc.org/labix.org/v2/mgo
// https://godoc.org/labix.org/v2/mgo
// go get gopkg.in/mgo.v2
// http://labix.org/mgo
// 待看文档

func NewMongoStore(connection string) (*MongoStore,error){
	sesseion,err := mgo.Dial(connection)
	if err != nil{
		return nil,err
	}
	return &MongoStore{session:sesseion},nil
}

// the Dial method attempts to connect ot he instance of MongoDB specified in the connection string


// 这个方法不在这里
func waitForDB(){
	for i:=0 ; i<10 ;i++{
		_,err := NewMongoStore("localhost")
		if err == nil{
			break
		}
		time.Sleep(1*time.Second)
	}
}