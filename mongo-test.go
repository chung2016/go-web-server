package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"cluster0-shard-00-00.smdsy.mongodb.net"},
		Username: "chung2016",
		Password: "",
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("okokokokokokokok")
	}
	defer session.Close()

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
