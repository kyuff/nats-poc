package main

import "fmt"
import (
	"github.com/nats-io/go-nats"
	"github.com/kyuff/nats-poc/model"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	c.Subscribe("user.created", handler)

	fmt.Println("Started Logs Service")

	for {
	}

}

func handler(channel string, user *model.UserCreatedEvent) {
	fmt.Printf("event logged: %s = %v\n", channel, user)
}
