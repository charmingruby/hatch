package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	Client mqtt.Client
}

func New(url string) (*MQTT, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(url)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MQTT{
		Client: client,
	}, nil
}

func (m *MQTT) Disconnect() {
	delay := uint(10000) // 10 seconds

	m.Client.Disconnect(delay)
}
