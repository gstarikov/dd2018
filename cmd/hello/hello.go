package main

import (
	"fmt"
	"os"
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


func main() {
	var op, name string

	if len (os.Args)>=2{
		op = os.Args[1]
	}
	if len(os.Args)>=3 {
		op = os.Args[2]
	}
	var g Greeter
	switch op {
	case "prefix":
		g = PrefixGreeter{}
	case "suffix":
		g = SuffixGreeter{}
	default:
		fmt.Printf("unexpected greeter[%s]\n",op)
		return
	}

	g.Greet(name)
}