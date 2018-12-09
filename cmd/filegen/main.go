package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
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

	res := make(chan error, n)

	for i := 0; i < n; i++ {
		go fileWriter(res, i, m)
	}
	for i := 0; i < n; i++ {
		if err := <-res; err != nil {
			log.Printf("file [%d], err -> %s", i, err)
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
	for j := 0; j < m; j++ {
		if _, err = writeRandNumbersToFile(file); err!= nil {
			res <- err
			return
		}
	}
	res <- nil
}

func writeRandNumbersToFile(w io.Writer) (int, error) {
	v := rand.Uint64()
	//bts := make([]byte, 0, max+1) // 8 is for uint64; 1 is for '\n'.
	bts := pool.Get().(*[]byte)
	defer pool.Put(bts)

	p := *bts
	p = strconv.AppendUint(p, v, 10)
	p = append(p, '\n')
	return w.Write(p)
}

var pool=sync.Pool {
	New: func()interface{}{
		p := make([]byte,0,max+1)
		return &p
	},
}
