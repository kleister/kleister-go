// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// NewDeleteUserFromTeamParams creates a new DeleteUserFromTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserFromTeamParams() *DeleteUserFromTeamParams {
	return &DeleteUserFromTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserFromTeamParamsWithTimeout creates a new DeleteUserFromTeamParams object
// with the ability to set a timeout on a request.
func NewDeleteUserFromTeamParamsWithTimeout(timeout time.Duration) *DeleteUserFromTeamParams {
	return &DeleteUserFromTeamParams{
		timeout: timeout,
	}
}

// NewDeleteUserFromTeamParamsWithContext creates a new DeleteUserFromTeamParams object
// with the ability to set a context for a request.
func NewDeleteUserFromTeamParamsWithContext(ctx context.Context) *DeleteUserFromTeamParams {
	return &DeleteUserFromTeamParams{
		Context: ctx,
	}
}

// NewDeleteUserFromTeamParamsWithHTTPClient creates a new DeleteUserFromTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserFromTeamParamsWithHTTPClient(client *http.Client) *DeleteUserFromTeamParams {
	return &DeleteUserFromTeamParams{
		HTTPClient: client,
	}
}

/* DeleteUserFromTeamParams contains all the parameters to send to the API endpoint
   for the delete user from team operation.

   Typically these are written to a http.Request.
*/
type DeleteUserFromTeamParams struct {

	/* UserID.

	   A user UUID or slug
	*/
	UserID string

	/* UserTeam.

	   The user team data to delete
	*/
	UserTeam *models.UserTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user from team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserFromTeamParams) WithDefaults() *DeleteUserFromTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user from team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserFromTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user from team params
func (o *DeleteUserFromTeamParams) WithTimeout(timeout time.Duration) *DeleteUserFromTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user from team params
func (o *DeleteUserFromTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user from team params
func (o *DeleteUserFromTeamParams) WithContext(ctx context.Context) *DeleteUserFromTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user from team params
func (o *DeleteUserFromTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user from team params
func (o *DeleteUserFromTeamParams) WithHTTPClient(client *http.Client) *DeleteUserFromTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user from team params
func (o *DeleteUserFromTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user from team params
func (o *DeleteUserFromTeamParams) WithUserID(userID string) *DeleteUserFromTeamParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user from team params
func (o *DeleteUserFromTeamParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserTeam adds the userTeam to the delete user from team params
func (o *DeleteUserFromTeamParams) WithUserTeam(userTeam *models.UserTeamParams) *DeleteUserFromTeamParams {
	o.SetUserTeam(userTeam)
	return o
}

// SetUserTeam adds the userTeam to the delete user from team params
func (o *DeleteUserFromTeamParams) SetUserTeam(userTeam *models.UserTeamParams) {
	o.UserTeam = userTeam
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserFromTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}
	if o.UserTeam != nil {
		if err := r.SetBodyParam(o.UserTeam); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
