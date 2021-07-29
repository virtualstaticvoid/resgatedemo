package main

import (
	"flag"
	"github.com/jirenius/go-res"
	"log"
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

	var natsUrl string
	flag.StringVar(&natsUrl, "nats", "nats://127.0.0.1:4222", "NATS Server URL")
	flag.Parse()

	s := res.NewService("example")

	s.Handle("model",
		res.Access(res.AccessGranted),
		res.GetModel(func(r res.ModelRequest) {
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
				log.Fatal(err)
			}
		}
	}()

	log.Printf("Connecting to %q\n", natsUrl)
	err := s.ListenAndServe(natsUrl)
	if err != nil {
		log.Fatal(err)
	}

}
