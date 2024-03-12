// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AttachUserToMod(params *AttachUserToModParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AttachUserToModOK, error)

	AttachUserToPack(params *AttachUserToPackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AttachUserToPackOK, error)

	AttachUserToTeam(params *AttachUserToTeamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AttachUserToTeamOK, error)

	CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserOK, error)

	DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserOK, error)

	DeleteUserFromMod(params *DeleteUserFromModParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserFromModOK, error)

	DeleteUserFromPack(params *DeleteUserFromPackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserFromPackOK, error)

	DeleteUserFromTeam(params *DeleteUserFromTeamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserFromTeamOK, error)

	ListUserMods(params *ListUserModsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserModsOK, error)

	ListUserPacks(params *ListUserPacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserPacksOK, error)

	ListUserTeams(params *ListUserTeamsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserTeamsOK, error)

	ListUsers(params *ListUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUsersOK, error)

	PermitUserMod(params *PermitUserModParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PermitUserModOK, error)

	PermitUserPack(params *PermitUserPackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PermitUserPackOK, error)

	PermitUserTeam(params *PermitUserTeamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PermitUserTeamOK, error)

	ShowUser(params *ShowUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShowUserOK, error)

	UpdateUser(params *UpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AttachUserToMod assigns a mod to user
*/
func (a *Client) AttachUserToMod(params *AttachUserToModParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AttachUserToModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachUserToModParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachUserToMod",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/mods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachUserToModReader{formats: a.formats},
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
	success, ok := result.(*AttachUserToModOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachUserToModDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachUserToPack assigns a pack to user
*/
func (a *Client) AttachUserToPack(params *AttachUserToPackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AttachUserToPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachUserToPackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachUserToPack",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/packs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachUserToPackReader{formats: a.formats},
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
	success, ok := result.(*AttachUserToPackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachUserToPackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachUserToTeam attaches a team to user
*/
func (a *Client) AttachUserToTeam(params *AttachUserToTeamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AttachUserToTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachUserToTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachUserToTeam",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachUserToTeamReader{formats: a.formats},
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
	success, ok := result.(*AttachUserToTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachUserToTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateUser creates a new user
*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateUser",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
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
	success, ok := result.(*CreateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteUser deletes a specific user
*/
func (a *Client) DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUser",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteUserFromMod removes a mod from user
*/
func (a *Client) DeleteUserFromMod(params *DeleteUserFromModParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserFromModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserFromModParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserFromMod",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/mods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserFromModReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserFromModOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteUserFromModDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteUserFromPack removes a pack from user
*/
func (a *Client) DeleteUserFromPack(params *DeleteUserFromPackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserFromPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserFromPackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserFromPack",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/packs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserFromPackReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserFromPackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteUserFromPackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteUserFromTeam removes a team from user
*/
func (a *Client) DeleteUserFromTeam(params *DeleteUserFromTeamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserFromTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserFromTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserFromTeam",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserFromTeamReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserFromTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteUserFromTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListUserMods fetches all mods assigned to user
*/
func (a *Client) ListUserMods(params *ListUserModsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserModsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserModsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListUserMods",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/mods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUserModsReader{formats: a.formats},
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
	success, ok := result.(*ListUserModsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListUserModsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListUserPacks fetches all packs assigned to user
*/
func (a *Client) ListUserPacks(params *ListUserPacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserPacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserPacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListUserPacks",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/packs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUserPacksReader{formats: a.formats},
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
	success, ok := result.(*ListUserPacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListUserPacksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListUserTeams fetches all teams attached to user
*/
func (a *Client) ListUserTeams(params *ListUserTeamsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUserTeamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserTeamsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListUserTeams",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUserTeamsReader{formats: a.formats},
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
	success, ok := result.(*ListUserTeamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListUserTeamsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListUsers fetches all available users
*/
func (a *Client) ListUsers(params *ListUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUsersReader{formats: a.formats},
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
	success, ok := result.(*ListUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PermitUserMod updates mod perms for user
*/
func (a *Client) PermitUserMod(params *PermitUserModParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PermitUserModOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermitUserModParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PermitUserMod",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/mods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PermitUserModReader{formats: a.formats},
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
	success, ok := result.(*PermitUserModOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PermitUserModDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PermitUserPack updates pack perms for user
*/
func (a *Client) PermitUserPack(params *PermitUserPackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PermitUserPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermitUserPackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PermitUserPack",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/packs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PermitUserPackReader{formats: a.formats},
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
	success, ok := result.(*PermitUserPackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PermitUserPackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PermitUserTeam updates team perms for user
*/
func (a *Client) PermitUserTeam(params *PermitUserTeamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PermitUserTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermitUserTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PermitUserTeam",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PermitUserTeamReader{formats: a.formats},
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
	success, ok := result.(*PermitUserTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PermitUserTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ShowUser fetches a specific user
*/
func (a *Client) ShowUser(params *ShowUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShowUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ShowUser",
		Method:             "GET",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ShowUserReader{formats: a.formats},
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
	success, ok := result.(*ShowUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ShowUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateUser updates a specific user
*/
func (a *Client) UpdateUser(params *UpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateUser",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
