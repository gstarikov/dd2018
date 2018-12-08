package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	addr = flag.String("addr", "0.0.0.0:3030", "addr to bind")
)

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Listeting on %s", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			// FIXME: error may allow repeat
			log.Fatalf("Cant accept %s", err)
		}

		log.Printf(
			"accepted new conection %s -> %s",
			conn.LocalAddr().String(),
			conn.RemoteAddr().String())

		var (
			bts   = make([]byte, 4096)
			lines [][]byte
			line  []byte
		)
		for {
			n, err := conn.Read(bts)
			if err != nil {
				log.Printf("err -> %s\n", err.Error())
				break
			}
			log.Printf("readed %d bytes\n", n)
			log.Printf("message[%s]\n", bts[:n])
			data := bts[:n]
			for {
				i := bytes.IndexByte(data, '\n')
				if i == -1 {
					line = append(line, data...)
					break
				} else {
					line = append(line, data[:i]...)
					lines = append(lines, line)
					line = nil
					data = data[i+1:]
				}
			}

			if !bytes.HasPrefix(lines[0],[]byte("GET")){
				conn.Close()
				continue
			}
			proto := bytes.Index(lines[0],[]byte("HTTP/1.1"))
			if proto == -1 {
				conn.Close()
				continue
			}
			resource := bytes.TrimSpace(lines[0][4 : proto-1])
			log.Printf("client wants %s",resource)

			fmt.Fprintf(conn,"200 OK\n\nContent-Length: 0\n\n")
			conn.Close()
		}

	}

}
