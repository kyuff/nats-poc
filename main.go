package main

import (
	"github.com/nats-io/go-nats"
	"github.com/kyuff/nats-poc/model"
	"math/rand"
	"fmt"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	event := &model.UserCreatedEvent{
		Uid:  fmt.Sprintf("%v", rand.Int31()),
		Name: "Hans Peter Hansen",
		Age:  rand.Intn(67),
	}

	c.Publish("user.created", event)

	fmt.Printf("Event Created: %s of age %d with id %s\n", event.Name, event.Age, event.Uid)

}
