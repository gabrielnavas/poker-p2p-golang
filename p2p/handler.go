package p2p

import (
	"fmt"
	"io"
)

type Handler interface {
	HandleMessage(*Message) error
}

type DefaultHandler struct{}

func NewDefaultHandler() Handler {
	return &DefaultHandler{}
}

func (h *DefaultHandler) HandleMessage(message *Message) error {
	msgBytes, err := io.ReadAll(message.Payload)
	if err != nil {
		return err
	}
	fmt.Printf("handleing the message from %s: %s", message.From, msgBytes)
	return nil
}
