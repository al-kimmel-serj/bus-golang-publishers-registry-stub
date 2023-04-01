package bus_golang_publishers_registry_stub

import (
	"github.com/al-kimmel-serj/bus-golang"
)

type Client struct {
	endpoints []bus.PublisherEndpoint
}

func (c *Client) Register(_ bus.EventName, _ bus.EventVersion, _ string, _ int) (func() error, error) {
	return func() error {
		return nil
	}, nil
}

func (c *Client) Watch(_ bus.EventName, _ bus.EventVersion, handler func([]bus.PublisherEndpoint)) (func() error, error) {
	handler(c.endpoints)
	return func() error {
		return nil
	}, nil
}

func New(endpoints []bus.PublisherEndpoint) *Client {
	return &Client{
		endpoints: endpoints,
	}
}
