package nats

import (
	"fmt"
	"messaging/model"
)

type NatsMsgService struct {
	js string
	c  model.MessageConfig
}

func NewNatsMsgService(config model.MessageConfig) (*NatsMsgService, error) {
	return &NatsMsgService{
		js: "nats js",
		c:  config,
	}, nil
}

func (n *NatsMsgService) Publish() {
	fmt.Println("Nats publishing")
}

func (n *NatsMsgService) Subscribe() {
	fmt.Println("Nats subscribing")
}
