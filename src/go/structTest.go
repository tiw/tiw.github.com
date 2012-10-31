package main
import (
    "fmt"
    "labix.org/v2/mgo"
)

type Person struct {
    Name string
    Phone string
}

type Order struct {
    Id int
    Customer Person
}

func main() {
    p := Person{"Ting Wang", "1360109827"}
    o := Order{1, p}
    fmt.Println(o.Customer.Name)

    saveOrder(o)
}
func saveOrder(order Order) {
    
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	con := session.DB("test").C("order")
  err = con.Insert(order)
  if err != nil {
      panic(err)
  }
}
