package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type ball struct{}

func main(){
	var b ball
	var wg sync.WaitGroup

	fmt.Println("threads -> ", runtime.GOMAXPROCS(0))
	fmt.Println("threads -> ", runtime.GOMAXPROCS(1))
	fmt.Println("threads -> ", runtime.GOMAXPROCS(0))

	table := make(chan ball)

	for _, name := range []string{"Pert","Ivan"} {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			player(name, table)
		}(name)
	}

	table <- b

	time.Sleep(5*time.Millisecond)
	fmt.Println("Stop the game. take a ball")
	<-table
	fmt.Println("close table")
	close(table)
	fmt.Println("waiting for shutdown")
	wg.Wait()
	fmt.Println("done add routines are down")

}

func player(name string, table chan ball){
	for b := range table {
		fmt.Printf("player[%s] -> %s\n", name, "YAY! Got the ball")
		table <- b
		time.Sleep(time.Millisecond)
	}
}
