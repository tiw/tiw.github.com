package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

type Order struct {
    OrderId int
    customer Person
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
  person := Person{"Ale", "+55 12 2314 2145"}
  order := Order{1, person}
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8456 1330"})
	if err != nil {
		panic(err)
	}
  c1 := session.DB("test").C("order")
  err = c1.Insert(order)
  if err != nil {
      panic(err)
  }

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)

  foundOrder := Order{}
  err = c1.Find(bson.M{"orderid":1}).One(&foundOrder)
  if err != nil {
      panic(err)
  }
  //fmt.Println("order:")
}
