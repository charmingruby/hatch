// Package mqtt provides the MQTT connection
package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

// Client is a wrapper to mqtt.Client, extending with default disconnect method.
type Client struct {
	Conn mqtt.Client
}

// New creates a Client instance
//
// Parameters:
//   - url: MQTT url (e.g.:"mqtt://localhost:1883")
//
// Returns :
//   - *Client: MQTT client wrapper instance
//   - error: if there is any on error on connecting to
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

// Disconnect closes the connection with MQTT, with a timeout of 10 seconds.
func (c *Client) Disconnect() {
	delay := uint(10000) // 10 seconds

	c.Conn.Disconnect(delay)
}
