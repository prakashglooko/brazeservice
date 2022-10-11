package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Index struct {
	l *log.Logger
}

func NewIndex(l *log.Logger) *Index {
	return &Index{l}
}

func (i *Index) EchoBody(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("In EchoBody")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "%s", d)
}
