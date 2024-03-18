// Code generated by go-swagger; DO NOT EDIT.

package neoforge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new neoforge API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for neoforge API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AppendNeoforgeToBuild(params *AppendNeoforgeToBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AppendNeoforgeToBuildOK, error)

	DeleteNeoforgeFromBuild(params *DeleteNeoforgeFromBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNeoforgeFromBuildOK, error)

	ListNeoforgeBuilds(params *ListNeoforgeBuildsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNeoforgeBuildsOK, error)

	ListNeoforges(params *ListNeoforgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNeoforgesOK, error)

	SearchNeoforges(params *SearchNeoforgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchNeoforgesOK, error)

	UpdateNeoforge(params *UpdateNeoforgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNeoforgeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AppendNeoforgeToBuild assigns a build to a neoforge version
*/
func (a *Client) AppendNeoforgeToBuild(params *AppendNeoforgeToBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AppendNeoforgeToBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendNeoforgeToBuildParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AppendNeoforgeToBuild",
		Method:             "POST",
		PathPattern:        "/neoforge/{neoforge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AppendNeoforgeToBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*AppendNeoforgeToBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AppendNeoforgeToBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteNeoforgeFromBuild unlinks a build from a neoforge version
*/
func (a *Client) DeleteNeoforgeFromBuild(params *DeleteNeoforgeFromBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNeoforgeFromBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNeoforgeFromBuildParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNeoforgeFromBuild",
		Method:             "DELETE",
		PathPattern:        "/neoforge/{neoforge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNeoforgeFromBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DeleteNeoforgeFromBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNeoforgeFromBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNeoforgeBuilds fetches the builds assigned to a neoforge version
*/
func (a *Client) ListNeoforgeBuilds(params *ListNeoforgeBuildsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNeoforgeBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNeoforgeBuildsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNeoforgeBuilds",
		Method:             "GET",
		PathPattern:        "/neoforge/{neoforge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNeoforgeBuildsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*ListNeoforgeBuildsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNeoforgeBuildsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNeoforges fetches the available neoforge versions
*/
func (a *Client) ListNeoforges(params *ListNeoforgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNeoforgesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNeoforgesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNeoforges",
		Method:             "GET",
		PathPattern:        "/neoforge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNeoforgesReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*ListNeoforgesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNeoforgesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchNeoforges searches for available neoforge versions
*/
func (a *Client) SearchNeoforges(params *SearchNeoforgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchNeoforgesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchNeoforgesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchNeoforges",
		Method:             "GET",
		PathPattern:        "/neoforge/{neoforge_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchNeoforgesReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*SearchNeoforgesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchNeoforgesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateNeoforge updates the available neoforge versions
*/
func (a *Client) UpdateNeoforge(params *UpdateNeoforgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNeoforgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNeoforgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateNeoforge",
		Method:             "PUT",
		PathPattern:        "/neoforge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateNeoforgeReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateNeoforgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateNeoforgeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}