package handlers

import (
	"context"
	"fmt"
	"glooko/brazeservice/data"
	"log"
	"net/http"
)

type Braze struct {
	l *log.Logger
}

func NewBraze(l *log.Logger) *Braze {
	return &Braze{l}
}

func (i *Braze) ListBrazeCalls(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("In GetBrazeCalls")
	bcb := data.GetBrazeCalls()
	bcb.ToJSON(rw)
}

func (i *Braze) AddBrazeCall(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("In AddBrazeCall")
	bc := r.Context().Value(BrazeCallKey{}).(data.BrazeCall)
	data.AddBrazeCall(&bc)
}

type BrazeCallKey struct{}

func (i *Braze) BrazeCallValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		bc := &data.BrazeCall{}

		err := bc.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to read body", http.StatusBadRequest)
			return
		}
		i.l.Printf("Call: %v", bc.Braze_data)
		err = bc.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Body fails validation %s", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), BrazeCallKey{}, *bc)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
