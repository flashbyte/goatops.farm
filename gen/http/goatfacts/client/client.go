// Code generated by goa v3.4.2, DO NOT EDIT.
//
// goatfacts client HTTP transport
//
// Command:
// $ goa gen github.com/martinohmann/goatops.farm/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the goatfacts service endpoint HTTP clients.
type Client struct {
	// GetFact Doer is the HTTP client used to make requests to the get-fact
	// endpoint.
	GetFactDoer goahttp.Doer

	// ListFacts Doer is the HTTP client used to make requests to the list-facts
	// endpoint.
	ListFactsDoer goahttp.Doer

	// GetRandomFact Doer is the HTTP client used to make requests to the
	// get-random-fact endpoint.
	GetRandomFactDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the goatfacts service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetFactDoer:         doer,
		ListFactsDoer:       doer,
		GetRandomFactDoer:   doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetFact returns an endpoint that makes HTTP requests to the goatfacts
// service get-fact server.
func (c *Client) GetFact() goa.Endpoint {
	var (
		decodeResponse = DecodeGetFactResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetFactRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetFactDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("goatfacts", "get-fact", err)
		}
		return decodeResponse(resp)
	}
}

// ListFacts returns an endpoint that makes HTTP requests to the goatfacts
// service list-facts server.
func (c *Client) ListFacts() goa.Endpoint {
	var (
		decodeResponse = DecodeListFactsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListFactsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListFactsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("goatfacts", "list-facts", err)
		}
		return decodeResponse(resp)
	}
}

// GetRandomFact returns an endpoint that makes HTTP requests to the goatfacts
// service get-random-fact server.
func (c *Client) GetRandomFact() goa.Endpoint {
	var (
		decodeResponse = DecodeGetRandomFactResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetRandomFactRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetRandomFactDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("goatfacts", "get-random-fact", err)
		}
		return decodeResponse(resp)
	}
}
