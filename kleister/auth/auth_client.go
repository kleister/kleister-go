// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	LoginUser(params *LoginUserParams, opts ...ClientOption) (*LoginUserOK, error)

	RefreshAuth(params *RefreshAuthParams, opts ...ClientOption) (*RefreshAuthOK, error)

	VerifyAuth(params *VerifyAuthParams, opts ...ClientOption) (*VerifyAuthOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LoginUser authenticates an user by credentials
*/
func (a *Client) LoginUser(params *LoginUserParams, opts ...ClientOption) (*LoginUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LoginUser",
		Method:             "POST",
		PathPattern:        "/auth/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LoginUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LoginUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RefreshAuth refreshes an auth token before it expires
*/
func (a *Client) RefreshAuth(params *RefreshAuthParams, opts ...ClientOption) (*RefreshAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RefreshAuth",
		Method:             "GET",
		PathPattern:        "/auth/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RefreshAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefreshAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RefreshAuthDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  VerifyAuth verifies validity for an authentication token
*/
func (a *Client) VerifyAuth(params *VerifyAuthParams, opts ...ClientOption) (*VerifyAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VerifyAuth",
		Method:             "GET",
		PathPattern:        "/auth/verify/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VerifyAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VerifyAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VerifyAuthDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
