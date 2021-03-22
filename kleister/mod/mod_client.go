// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new mod API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mod API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AppendModToTeam assigns a team to mod
*/
func (a *Client) AppendModToTeam(params *AppendModToTeamParams) (*AppendModToTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendModToTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AppendModToTeam",
		Method:             "POST",
		PathPattern:        "/mods/{mod_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AppendModToTeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppendModToTeamOK), nil

}

/*
AppendModToUser assigns a user to mod
*/
func (a *Client) AppendModToUser(params *AppendModToUserParams) (*AppendModToUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendModToUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AppendModToUser",
		Method:             "POST",
		PathPattern:        "/mods/{mod_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AppendModToUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppendModToUserOK), nil

}

/*
AppendVersionToBuild assigns a build to a version
*/
func (a *Client) AppendVersionToBuild(params *AppendVersionToBuildParams) (*AppendVersionToBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendVersionToBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AppendVersionToBuild",
		Method:             "POST",
		PathPattern:        "/mods/{mod_id}/versions/{version_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AppendVersionToBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppendVersionToBuildOK), nil

}

/*
CreateMod creates a new mod
*/
func (a *Client) CreateMod(params *CreateModParams) (*CreateModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateModParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateMod",
		Method:             "POST",
		PathPattern:        "/mods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateModReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateModOK), nil

}

/*
CreateVersion creates a new version for a mod
*/
func (a *Client) CreateVersion(params *CreateVersionParams) (*CreateVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateVersion",
		Method:             "POST",
		PathPattern:        "/mods/{mod_id}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateVersionOK), nil

}

/*
DeleteMod deletes a specific mod
*/
func (a *Client) DeleteMod(params *DeleteModParams) (*DeleteModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteMod",
		Method:             "DELETE",
		PathPattern:        "/mods/{mod_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteModReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteModOK), nil

}

/*
DeleteModFromTeam removes a team from mod
*/
func (a *Client) DeleteModFromTeam(params *DeleteModFromTeamParams) (*DeleteModFromTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModFromTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteModFromTeam",
		Method:             "DELETE",
		PathPattern:        "/mods/{mod_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteModFromTeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteModFromTeamOK), nil

}

/*
DeleteModFromUser removes a user from mod
*/
func (a *Client) DeleteModFromUser(params *DeleteModFromUserParams) (*DeleteModFromUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModFromUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteModFromUser",
		Method:             "DELETE",
		PathPattern:        "/mods/{mod_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteModFromUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteModFromUserOK), nil

}

/*
DeleteVersion deletes a specific version for a mod
*/
func (a *Client) DeleteVersion(params *DeleteVersionParams) (*DeleteVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVersion",
		Method:             "DELETE",
		PathPattern:        "/mods/{mod_id}/versions/{version_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteVersionOK), nil

}

/*
DeleteVersionFromBuild unlinks a build from a version
*/
func (a *Client) DeleteVersionFromBuild(params *DeleteVersionFromBuildParams) (*DeleteVersionFromBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionFromBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVersionFromBuild",
		Method:             "DELETE",
		PathPattern:        "/mods/{mod_id}/versions/{version_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteVersionFromBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteVersionFromBuildOK), nil

}

/*
ListModTeams fetches all teams assigned to mod
*/
func (a *Client) ListModTeams(params *ListModTeamsParams) (*ListModTeamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListModTeamsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListModTeams",
		Method:             "GET",
		PathPattern:        "/mods/{mod_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListModTeamsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListModTeamsOK), nil

}

/*
ListModUsers fetches all users assigned to mod
*/
func (a *Client) ListModUsers(params *ListModUsersParams) (*ListModUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListModUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListModUsers",
		Method:             "GET",
		PathPattern:        "/mods/{mod_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListModUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListModUsersOK), nil

}

/*
ListMods fetches all available mods
*/
func (a *Client) ListMods(params *ListModsParams) (*ListModsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListModsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListMods",
		Method:             "GET",
		PathPattern:        "/mods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListModsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListModsOK), nil

}

/*
ListVersionBuilds fetches all builds assigned to version
*/
func (a *Client) ListVersionBuilds(params *ListVersionBuildsParams) (*ListVersionBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVersionBuilds",
		Method:             "GET",
		PathPattern:        "/mods/{mod_id}/versions/{version_id}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListVersionBuildsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVersionBuildsOK), nil

}

/*
ListVersions fetches all available versions for a mod
*/
func (a *Client) ListVersions(params *ListVersionsParams) (*ListVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVersions",
		Method:             "GET",
		PathPattern:        "/mods/{mod_id}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVersionsOK), nil

}

/*
PermitModTeam updates team perms for mod
*/
func (a *Client) PermitModTeam(params *PermitModTeamParams) (*PermitModTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermitModTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PermitModTeam",
		Method:             "PUT",
		PathPattern:        "/mods/{mod_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PermitModTeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PermitModTeamOK), nil

}

/*
PermitModUser updates user perms for mod
*/
func (a *Client) PermitModUser(params *PermitModUserParams) (*PermitModUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermitModUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PermitModUser",
		Method:             "PUT",
		PathPattern:        "/mods/{mod_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PermitModUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PermitModUserOK), nil

}

/*
ShowMod fetches a specific mod
*/
func (a *Client) ShowMod(params *ShowModParams) (*ShowModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowModParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ShowMod",
		Method:             "GET",
		PathPattern:        "/mods/{mod_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ShowModReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShowModOK), nil

}

/*
ShowVersion fetches a specific version for a mod
*/
func (a *Client) ShowVersion(params *ShowVersionParams) (*ShowVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ShowVersion",
		Method:             "GET",
		PathPattern:        "/mods/{mod_id}/versions/{version_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ShowVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShowVersionOK), nil

}

/*
UpdateMod updates a specific mod
*/
func (a *Client) UpdateMod(params *UpdateModParams) (*UpdateModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateModParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateMod",
		Method:             "PUT",
		PathPattern:        "/mods/{mod_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateModReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateModOK), nil

}

/*
UpdateVersion updates a specific version for a mod
*/
func (a *Client) UpdateVersion(params *UpdateVersionParams) (*UpdateVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateVersion",
		Method:             "PUT",
		PathPattern:        "/mods/{mod_id}/versions/{version_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateVersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}