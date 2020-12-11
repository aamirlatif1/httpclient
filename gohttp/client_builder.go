package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers           http.Header
	maxIdleConnection int
	connectionTimeout time.Duration
	responseTimeout   time.Duration
	disableTimeouts   bool
	client            *http.Client
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableTimeouts(disabled bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) Build() Client {
	return &httpClient{
		builder: c,
	}
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnection = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disabled bool) ClientBuilder {
	c.disableTimeouts = disabled
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}
