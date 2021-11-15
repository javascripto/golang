package client

import "fmt"

type Client struct {
	Name, Document, Occupation string
}

func (client Client) String() string {
	return fmt.Sprintf(
		`{"Name":"%s","Document":"%s", "Occupation":"%s"}`,
		client.Name,
		client.Document,
		client.Occupation,
	)
}
