package main

import (
	"fmt"
	"github.com/jirenius/go-res"
	"time"
)

type Model struct {
	Message   string `json:"message"`
	UpdatedAt string `json:"updated-at"`
}

var model = &Model{
	Message:   "Hello, World!",
	UpdatedAt: time.Now().UTC().String(),
}

func main() {

	fmt.Println("Starting...")

	s := res.NewService("example")

	s.Handle("model",
		res.Access(res.AccessGranted),
		res.GetModel(func(r res.ModelRequest) {
			fmt.Println("Request")
			r.Model(model)
		}),
	)

	go func() {
		for _ = range time.Tick(time.Second * 10) {
			err := s.With("example.model", func(r res.Resource) {
				model.UpdatedAt = time.Now().UTC().String()
				r.Event("tick", model)
			})
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
				break
			}
		}
	}()

	err := s.ListenAndServe("nats://nats:4222")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

}
