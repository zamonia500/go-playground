package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	fmt.Printf("Process : %s\n", data)
	return data
}

func (lp LogicProvider) Action(data string) string {
	fmt.Printf("Action : %s\n", data)
	return data
}

type Logic interface {
	Process(data string) string
}

type Magic interface {
	Action(data string) string
}

type Client struct {
	// L Logic
	M Magic
}

func (c Client) Program() {
	// c.L.Process("Hi there")
	c.M.Action("Hello world")
}

func main() {
	c := Client{
		// L: LogicProvider{},
		M: LogicProvider{},
	}
	c.Program()

}
