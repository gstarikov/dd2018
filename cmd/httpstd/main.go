package main

import (
	"context"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"sync"
)


var (
	addr = flag.String("addr", "0.0.0.0:3030", "addr to bind")
)

type Message struct {
	User, Message string
}

type MessageStorage interface {
	Put(context.Context, Message)error
	List(ctx context.Context) ([]Message, error)
}
type InMemoryStrage struct {
	mu sync.Mutex
	messages []Message
}

func (t *InMemoryStrage) Put(_ context.Context, msg Message)error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.messages = append(t.messages,msg)
	return nil
}

func (t *InMemoryStrage) List(_ context.Context)([]Message, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.messages,nil
}

type IndexPage struct {
	Title string
	Messages []Message
}

const indexText = `asd`

var index = template.Must(template.New("index").Parse(indexText))

func main() {
	flag.Parse()

	var st MessageStorage

	http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
		log.Printf("got [%s]-> %s\n",req.Method,req.RequestURI)

		switch req.Method {
		case http.MethodGet:
		case http.MethodPost:
			err := req.ParseForm()
			if err != nil {
				log.Printf("Formarse error -> %s",err.Error())
				io.WriteString(res,err.Error())
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			user := req.Form.Get("user")
			message := req.Form.Get("message")
			if user == "" || message == "" {
				io.WriteString(res,"user or message are empty")
				log.Printf("user or message are empty\n")
				res.WriteHeader(http.StatusBadRequest)
			}

			if err := st.Put(nil,Message{User:user,Message: message}); err != nil {
				log.Printf("cant write fo storage")
				io.WriteString(res, "cant write to storage")
				res.WriteHeader(http.StatusInternalServerError)
			}

		}


		io.WriteString(res,"ehlo")
	})
	http.ListenAndServe(*addr,nil)
}
