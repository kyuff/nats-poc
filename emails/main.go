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

	c.QueueSubscribe("user.created", "emails", handler)

	fmt.Println("Started Emails Service")

	for {
	}

}

func handler(user *model.UserCreatedEvent) {
	fmt.Printf("User created: %s of age %d with id %s\n", user.Name, user.Age, user.Uid)
}
