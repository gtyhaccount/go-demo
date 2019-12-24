package rabbitmq

import "testing"

func TestSendMessage(t *testing.T) {
	SendMessage("Hello world!")

	ReceiveMessage()
}
