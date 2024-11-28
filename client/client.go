package client

import (
	"context"
	"github.com/tsinghua-cel/attacker-client-go/rpc"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string, valIdx int) (*Client, error) {
	return DialContext(context.Background(), rawurl, valIdx)
}

// DialContext connects a client to the given URL with context.
func DialContext(ctx context.Context, rawurl string, valIdx int) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c, valIdx), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, valIdx int) *Client {
	client := &Client{
		c: c,
	}
	return client
}

// Close closes the underlying RPC connection.
func (ec *Client) Close() {
	ec.c.Close()
}

// Client gets the underlying RPC client.
func (ec *Client) Client() *rpc.Client {
	return ec.c
}
