package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Client struct {
	Conn mqtt.Client
}

func New(url string) (*Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(url)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Client{
		Conn: client,
	}, nil
}

func (c *Client) Disconnect() {
	delay := uint(10000) // 10 seconds

	c.Conn.Disconnect(delay)
}
