package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	n int // Number of files
	m int // Number of entries per file
)

var max = len(strconv.AppendUint(nil, math.MaxUint64, 10))

func main() {
	flag.IntVar(&n, "n", 1, "number of files to generate")
	flag.IntVar(&m, "m", 5, "number of entries to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	res := make(chan error,n)

	for i := 0; i < n; i++ {
		go fileWriter(res,i,m)
	}
	for i := 0; i < n; i++ {
		if err := <-res; err != nil {
			log.Printf("file [%d], err -> %s",i,err)
		}
	}
}

func fileWriter(res chan<- error, i, m int) {
	var err error
	defer func() {
		if err != nil {
			res <- err
		}
	}()

	fName := fmt.Sprintf("%d.gen", i)
	var file *os.File
	file, err = os.Create(fName)
	if err != nil {
		return
	}
	defer file.Close()
	err = writeRandNumbersToFile(file)
	if err != nil {
		return
	}

	for j := 0; j < m; j++ {
		if err = writeRandNumbersToFile(file); err != nil {
			return
		}
	}
	res <- nil
}

func writeRandNumbersToFile(file *os.File) error {
	bts := randNumberBytes()
	if _, err := file.Write(bts); err != nil {
		return err
	}
	return nil
}

func randNumberBytes() []byte {
	v := rand.Uint64()
	bts := make([]byte, 0, max+1) // 8 is for uint64; 1 is for '\n'.
	bts = strconv.AppendUint(nil, v, 10)
	bts = append(bts, '\n')
	return bts
}
