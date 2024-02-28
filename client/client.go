package client

import (
	"github.com/go-zoox/database-as-a-service/data"
	"github.com/go-zoox/fetch"
)

// Client is the interface for database as a service client
type Client interface {
	Request(payload *data.Request) (result any, err error)
}

// Option is the option for database as a service client
type Option struct {
	Server   string
	Username string
	Password string
}

type client struct {
	opt *Option
}

// New creates a new database as a service client
func New(opts ...func(opt *Option)) Client {
	opt := &Option{}
	for _, o := range opts {
		o(opt)
	}

	return &client{
		opt: opt,
	}
}

func (c *client) Request(payload *data.Request) (result any, err error) {
	response, err := fetch.Post(c.opt.Server, &fetch.Config{
		Headers: map[string]string{
			"Context-Type": "application/json",
		},
		Body: payload,
	})
	if err != nil {
		return nil, err
	}

	if !response.Ok() {
		return nil, response.Error()
	}

	resp := &data.Response{}
	if err = response.UnmarshalJSON(resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
