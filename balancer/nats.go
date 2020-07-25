package balancer

import (
	"checker/config"
	"log"

	"github.com/nats-io/go-nats"
)

func New(n config.Nats) *nats.Conn {
	nc, err := nats.Connect(n.Host)
	if err != nil {
		log.Fatal(err)
	}

	return nc
}
