package pool

import (
	"log"
	nats "github.com/nats-io/nats.go"
)

var Nats = new(nats.EncodedConn)

func NatsConnect() {
	log.Print(Config.Nats.Url)
	nc, err := nats.Connect(Config.Nats.Url)
	if err != nil {
		log.Fatal("nats connection failed", err)
	}
	Nats, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal("nats encoded connection failed", err)
	}
	log.Print("nats successfully connected!")
}
