package main

import (
	"fmt"
	"messaging/messaging"
	"messaging/model"
)

func main() {
	fmt.Println("starting")
	config := model.MessageConfig{
		Name: "Nats",
		Url:  "My url",
	}
	msgService, err := messaging.NewMessageService(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	msgService.Publish()
	msgService.Subscribe()
	fmt.Println("End...")
}
