package messaging

import (
	"errors"
	"messaging/model"
	"messaging/nats"
)

type MessagingService interface {
	Publish()
	Subscribe()
}

func NewMessageService(config model.MessageConfig) (MessagingService, error) {
	if config.Name == "Nats" {
		svc, _ := nats.NewNatsMsgService(config)
		return svc, nil
	}
	return nil, errors.New("Wrong type")
}
