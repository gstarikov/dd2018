package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("cmd line total args count is %d. Wow!\n",len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("\targ[%d] -> %s\n",i,v)
	}
	fmt.Println("hello world")
}