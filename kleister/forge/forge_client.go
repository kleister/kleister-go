// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new forge API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for forge API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AppendForgeToBuild assigns a build to a forge version
*/
func (a *Client) AppendForgeToBuild(params *AppendForgeToBuildParams) (*AppendForgeToBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendForgeToBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AppendForgeToBuild",
		Method:             "POST",
		PathPattern:        "/forge/{forge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AppendForgeToBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppendForgeToBuildOK), nil

}

/*
DeleteForgeFromBuild unlinks a build from a forge version
*/
func (a *Client) DeleteForgeFromBuild(params *DeleteForgeFromBuildParams) (*DeleteForgeFromBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteForgeFromBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteForgeFromBuild",
		Method:             "DELETE",
		PathPattern:        "/forge/{forge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteForgeFromBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteForgeFromBuildOK), nil

}

/*
ListForgeBuilds fetches the builds assigned to a forge version
*/
func (a *Client) ListForgeBuilds(params *ListForgeBuildsParams) (*ListForgeBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListForgeBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListForgeBuilds",
		Method:             "GET",
		PathPattern:        "/forge/{forge_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListForgeBuildsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListForgeBuildsOK), nil

}

/*
ListForges fetches the available forge versions
*/
func (a *Client) ListForges(params *ListForgesParams) (*ListForgesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListForgesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListForges",
		Method:             "GET",
		PathPattern:        "/forge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListForgesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListForgesOK), nil

}

/*
SearchForges searches for available forge versions
*/
func (a *Client) SearchForges(params *SearchForgesParams) (*SearchForgesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchForgesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SearchForges",
		Method:             "GET",
		PathPattern:        "/forge/{forge_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchForgesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchForgesOK), nil

}

/*
UpdateForge updates the available forge versions
*/
func (a *Client) UpdateForge(params *UpdateForgeParams) (*UpdateForgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateForgeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateForge",
		Method:             "PUT",
		PathPattern:        "/forge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateForgeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateForgeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}