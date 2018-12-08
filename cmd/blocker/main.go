package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	runtime.GOMAXPROCS(1)

	go monitor()
	go cpueater()

	var block chan struct{}
	<-block
}

func cpueater(){
	for now := range time.Tick(time.Second) {
		fmt.Println("eater started", now)
		for i := 0; i < 1e10; i++{
			if i % 1e4 == 0 {
				runtime.Gosched()
			}
		}
		fmt.Println("eater done",now)
	}
}

func monitor() {
	const duration = 500 * time.Millisecond
	lastTick := time.Now()
	for now := range time.Tick(duration) {
		waitDuration := now.Sub(lastTick)
		lastTick=now
		if waitDuration >= 2*duration{
			fmt.Println("overdelay... ",waitDuration)
		}else {
			fmt.Println("time",now)
		}
	}
}