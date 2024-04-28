package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Person struct {
	Name       string `json:"customName"`
	Age        int    `json:"age,omitempty"`
	CreditCard string
}

func main() {
	p := Person{Name: "tom", CreditCard: "Super secret"}
	pBytes, err := json.Marshal(p)

	log.Print(err)
	log.Print(string(pBytes))

	// p2AsRawJson := json.RawMessage(`{"customName": "mary"}`)
	p2AsRawJson, err3 := ioutil.ReadFile("mary.json")
	log.Print(err3)

	var p2 Person
	err2 := json.Unmarshal(p2AsRawJson, &p2)

	log.Print(err2)
	log.Printf("%+v", p2)
}
