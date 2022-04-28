// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mod API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mod API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AppendModToTeam(params *AppendModToTeamParams, opts ...ClientOption) (*AppendModToTeamOK, error)

	AppendModToUser(params *AppendModToUserParams, opts ...ClientOption) (*AppendModToUserOK, error)

	AppendVersionToBuild(params *AppendVersionToBuildParams, opts ...ClientOption) (*AppendVersionToBuildOK, error)

	CreateMod(params *CreateModParams, opts ...ClientOption) (*CreateModOK, error)

	CreateVersion(params *CreateVersionParams, opts ...ClientOption) (*CreateVersionOK, error)

	DeleteMod(params *DeleteModParams, opts ...ClientOption) (*DeleteModOK, error)

	DeleteModFromTeam(params *DeleteModFromTeamParams, opts ...ClientOption) (*DeleteModFromTeamOK, error)

	DeleteModFromUser(params *DeleteModFromUserParams, opts ...ClientOption) (*DeleteModFromUserOK, error)

	DeleteVersion(params *DeleteVersionParams, opts ...ClientOption) (*DeleteVersionOK, error)

	DeleteVersionFromBuild(params *DeleteVersionFromBuildParams, opts ...ClientOption) (*DeleteVersionFromBuildOK, error)

	ListModTeams(params *ListModTeamsParams, opts ...ClientOption) (*ListModTeamsOK, error)

	ListModUsers(params *ListModUsersParams, opts ...ClientOption) (*ListModUsersOK, error)

	ListMods(params *ListModsParams, opts ...ClientOption) (*ListModsOK, error)

	ListVersionBuilds(params *ListVersionBuildsParams, opts ...ClientOption) (*ListVersionBuildsOK, error)

	ListVersions(params *ListVersionsParams, opts ...ClientOption) (*ListVersionsOK, error)

	PermitModTeam(params *PermitModTeamParams, opts ...ClientOption) (*PermitModTeamOK, error)

	PermitModUser(params *PermitModUserParams, opts ...ClientOption) (*PermitModUserOK, error)

	ShowMod(params *ShowModParams, opts ...ClientOption) (*ShowModOK, error)

	ShowVersion(params *ShowVersionParams, opts ...ClientOption) (*ShowVersionOK, error)

	UpdateMod(params *UpdateModParams, opts ...ClientOption) (*UpdateModOK, error)

	UpdateVersion(params *UpdateVersionParams, opts ...ClientOption) (*UpdateVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AppendModToTeam assigns a team to mod
*/
func (a *Client) AppendModToTeam(params *AppendModToTeamParams, opts ...ClientOption) (*AppendModToTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendModToTeamParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AppendModToTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AppendModToTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AppendModToUser assigns a user to mod
*/
func (a *Client) AppendModToUser(params *AppendModToUserParams, opts ...ClientOption) (*AppendModToUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendModToUserParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AppendModToUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AppendModToUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AppendVersionToBuild assigns a build to a version
*/
func (a *Client) AppendVersionToBuild(params *AppendVersionToBuildParams, opts ...ClientOption) (*AppendVersionToBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppendVersionToBuildParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AppendVersionToBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AppendVersionToBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateMod creates a new mod
*/
func (a *Client) CreateMod(params *CreateModParams, opts ...ClientOption) (*CreateModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateModParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateModOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateModDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateVersion creates a new version for a mod
*/
func (a *Client) CreateVersion(params *CreateVersionParams, opts ...ClientOption) (*CreateVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVersionParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteMod deletes a specific mod
*/
func (a *Client) DeleteMod(params *DeleteModParams, opts ...ClientOption) (*DeleteModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteModOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteModDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteModFromTeam removes a team from mod
*/
func (a *Client) DeleteModFromTeam(params *DeleteModFromTeamParams, opts ...ClientOption) (*DeleteModFromTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModFromTeamParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteModFromTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteModFromTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteModFromUser removes a user from mod
*/
func (a *Client) DeleteModFromUser(params *DeleteModFromUserParams, opts ...ClientOption) (*DeleteModFromUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModFromUserParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteModFromUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteModFromUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteVersion deletes a specific version for a mod
*/
func (a *Client) DeleteVersion(params *DeleteVersionParams, opts ...ClientOption) (*DeleteVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteVersionFromBuild unlinks a build from a version
*/
func (a *Client) DeleteVersionFromBuild(params *DeleteVersionFromBuildParams, opts ...ClientOption) (*DeleteVersionFromBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionFromBuildParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVersionFromBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVersionFromBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListModTeams fetches all teams assigned to mod
*/
func (a *Client) ListModTeams(params *ListModTeamsParams, opts ...ClientOption) (*ListModTeamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListModTeamsParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListModTeamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListModTeamsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListModUsers fetches all users assigned to mod
*/
func (a *Client) ListModUsers(params *ListModUsersParams, opts ...ClientOption) (*ListModUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListModUsersParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListModUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListModUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListMods fetches all available mods
*/
func (a *Client) ListMods(params *ListModsParams, opts ...ClientOption) (*ListModsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListModsParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListModsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListModsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVersionBuilds fetches all builds assigned to version
*/
func (a *Client) ListVersionBuilds(params *ListVersionBuildsParams, opts ...ClientOption) (*ListVersionBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionBuildsParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVersionBuildsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVersionBuildsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVersions fetches all available versions for a mod
*/
func (a *Client) ListVersions(params *ListVersionsParams, opts ...ClientOption) (*ListVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionsParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PermitModTeam updates team perms for mod
*/
func (a *Client) PermitModTeam(params *PermitModTeamParams, opts ...ClientOption) (*PermitModTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermitModTeamParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PermitModTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PermitModTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PermitModUser updates user perms for mod
*/
func (a *Client) PermitModUser(params *PermitModUserParams, opts ...ClientOption) (*PermitModUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermitModUserParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PermitModUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PermitModUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ShowMod fetches a specific mod
*/
func (a *Client) ShowMod(params *ShowModParams, opts ...ClientOption) (*ShowModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowModParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ShowModOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ShowModDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ShowVersion fetches a specific version for a mod
*/
func (a *Client) ShowVersion(params *ShowVersionParams, opts ...ClientOption) (*ShowVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowVersionParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ShowVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ShowVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateMod updates a specific mod
*/
func (a *Client) UpdateMod(params *UpdateModParams, opts ...ClientOption) (*UpdateModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateModParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateModOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateModDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateVersion updates a specific version for a mod
*/
func (a *Client) UpdateVersion(params *UpdateVersionParams, opts ...ClientOption) (*UpdateVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVersionParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
