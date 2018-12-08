package main

import (
	"flag"
	"os"
)

var (
	n int // Number of files
	m int // Number of entries per file
)

func main(){
	flag.IntVar(&n,"n",1,"number of files to generate")
	flag.IntVar(&m,"m",1,"number of entries to generate")

	for i :=0; i < n; i++ {
		file, err := os.Create("my-file")

	}
}
