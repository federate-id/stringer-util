package main

import (
	"fmt"

	"github.com/federate-id/stringer-util/stringer"
)

type SecretKey struct {
	Key     string `stringer:"masked,length,name=APIKey"`
}

type Client struct {
	ID       string            `stringer:"include"`
	Secret   SecretKey         `stringer:"nested"`
	Metadata map[string]any `stringer:"type"`
	Comment  string            // excluded
}

func (c *Client) String() string {
	return stringer.ToStringWithTags(c)
}

func main() {
	client := Client{
		ID: "uuuuuu",
		Secret: SecretKey{
			Key: "abc1234def",
		},
		Metadata: map[string]any{
			"foo": "hello",
			"bar": "world",
		},
		Comment: "not including this field",
	}
	fmt.Println(client.String())
}
