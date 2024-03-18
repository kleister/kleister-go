// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new forge API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for forge API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AppendForgeToBuild(params *AppendForgeToBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AppendForgeToBuildOK, error)

	DeleteForgeFromBuild(params *DeleteForgeFromBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteForgeFromBuildOK, error)

	ListForgeBuilds(params *ListForgeBuildsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListForgeBuildsOK, error)

	ListForges(params *ListForgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListForgesOK, error)

	SearchForges(params *SearchForgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchForgesOK, error)

	UpdateForge(params *UpdateForgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateForgeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AppendForgeToBuild assigns a build to a forge version
*/
func (a *Client) AppendForgeToBuild(params *AppendForgeToBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AppendForgeToBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendForgeToBuildParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AppendForgeToBuild",
		Method:             "POST",
		PathPattern:        "/forge/{forge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AppendForgeToBuildReader{formats: a.formats},
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
	success, ok := result.(*AppendForgeToBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AppendForgeToBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteForgeFromBuild unlinks a build from a forge version
*/
func (a *Client) DeleteForgeFromBuild(params *DeleteForgeFromBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteForgeFromBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteForgeFromBuildParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteForgeFromBuild",
		Method:             "DELETE",
		PathPattern:        "/forge/{forge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteForgeFromBuildReader{formats: a.formats},
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
	success, ok := result.(*DeleteForgeFromBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteForgeFromBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListForgeBuilds fetches the builds assigned to a forge version
*/
func (a *Client) ListForgeBuilds(params *ListForgeBuildsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListForgeBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListForgeBuildsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListForgeBuilds",
		Method:             "GET",
		PathPattern:        "/forge/{forge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListForgeBuildsReader{formats: a.formats},
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
	success, ok := result.(*ListForgeBuildsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListForgeBuildsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListForges fetches the available forge versions
*/
func (a *Client) ListForges(params *ListForgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListForgesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListForgesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListForges",
		Method:             "GET",
		PathPattern:        "/forge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListForgesReader{formats: a.formats},
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
	success, ok := result.(*ListForgesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListForgesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchForges searches for available forge versions
*/
func (a *Client) SearchForges(params *SearchForgesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchForgesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchForgesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchForges",
		Method:             "GET",
		PathPattern:        "/forge/{forge_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchForgesReader{formats: a.formats},
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
	success, ok := result.(*SearchForgesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchForgesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateForge updates the available forge versions
*/
func (a *Client) UpdateForge(params *UpdateForgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateForgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateForgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateForge",
		Method:             "PUT",
		PathPattern:        "/forge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateForgeReader{formats: a.formats},
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
	success, ok := result.(*UpdateForgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateForgeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}