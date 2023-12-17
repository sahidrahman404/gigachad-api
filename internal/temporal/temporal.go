package temporal

import (
	"go.temporal.io/sdk/client"
)

func NewClient() (client.Client, error) {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		// log.Fatalln("Unable to create client", err)
		return nil, err
	}
	defer c.Close()
	return c, nil
}
