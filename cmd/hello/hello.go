package main

import (
	"flag"
	"fmt"
)

// Greeter greets someone by name
type Greeter interface {
	Greet(name string)
}

type PrefixGreeter struct {

}

func (p PrefixGreeter) Greet(name string) {
	fmt.Println("hello",name)
}

type SuffixGreeter struct {

}

func (p SuffixGreeter) Greet(name string) {
	fmt.Println(name,"hello")
}

var greeterType = flag.String("greeter","suffix","type of the greeter")

func main() {
	flag.Parse()


	var g Greeter
	switch *greeterType {
	case "prefix":
		g = PrefixGreeter{}
	case "suffix":
		g = SuffixGreeter{}
	default:
		fmt.Printf("unexpected greeter\n")
		return
	}

	g.Greet("world")
}