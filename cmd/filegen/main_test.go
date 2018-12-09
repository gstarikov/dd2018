package main

import (
	"io/ioutil"
	"strconv"
	"testing"
)

type stubFileWriter struct {
	//
	writes [][]byte
}

func (s *stubFileWriter) Write(p []byte) (int, error) {
	//
	s.writes = append(s.writes, append(([]byte)(nil),p...));
	return len(p), nil
}

func TestWriteRandNumber( t *testing.T) {
	var s stubFileWriter
	_, err := writeRandNumbersToFile(&s)
	if err != nil {
		t.Fatal(err)
	}
	if act, exp := len(s.writes), 1; act != exp {
		t.Fatalf("unxepected write number. calls: %d; wants: %d\n",act, exp)
	}
	p := s.writes[0]
	if n := len(p); n < 1 || p[n-1] != '\n' {
		t.Fatalf("mailformed bytes written")
	}

	num := p[:len(p)-1]
	_, err = strconv.ParseUint(string(num),10,64)
	if err != nil {
		t.Fatalf("invalid nubmer [%s]",num)
	}
	t.Logf("received %+q",s.writes)
}

func BenchmarkWriteRandNumber(b *testing.B) {
	//var s stubFileWriter
	for i:=0; i< b.N; i++ {
		_, _ = writeRandNumbersToFile(ioutil.Discard)

	}
}