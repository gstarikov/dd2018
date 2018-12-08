package main

type ball struct{}

func main(){
	var b ball

	table := make(chan ball)

	go player(table)
	go player(table)

}

func player(table struct{}){

}
