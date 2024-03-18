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

	"github.com/kleister/kleister-go/models"
)

// NewAttachUserToTeamParams creates a new AttachUserToTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachUserToTeamParams() *AttachUserToTeamParams {
	return &AttachUserToTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachUserToTeamParamsWithTimeout creates a new AttachUserToTeamParams object
// with the ability to set a timeout on a request.
func NewAttachUserToTeamParamsWithTimeout(timeout time.Duration) *AttachUserToTeamParams {
	return &AttachUserToTeamParams{
		timeout: timeout,
	}
}

// NewAttachUserToTeamParamsWithContext creates a new AttachUserToTeamParams object
// with the ability to set a context for a request.
func NewAttachUserToTeamParamsWithContext(ctx context.Context) *AttachUserToTeamParams {
	return &AttachUserToTeamParams{
		Context: ctx,
	}
}

// NewAttachUserToTeamParamsWithHTTPClient creates a new AttachUserToTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachUserToTeamParamsWithHTTPClient(client *http.Client) *AttachUserToTeamParams {
	return &AttachUserToTeamParams{
		HTTPClient: client,
	}
}

/*
AttachUserToTeamParams contains all the parameters to send to the API endpoint

	for the attach user to team operation.

	Typically these are written to a http.Request.
*/
type AttachUserToTeamParams struct {

	/* UserID.

	   A user UUID or slug
	*/
	UserID string

	/* UserTeam.

	   The user team data to attach
	*/
	UserTeam *models.UserTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attach user to team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachUserToTeamParams) WithDefaults() *AttachUserToTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attach user to team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachUserToTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attach user to team params
func (o *AttachUserToTeamParams) WithTimeout(timeout time.Duration) *AttachUserToTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach user to team params
func (o *AttachUserToTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach user to team params
func (o *AttachUserToTeamParams) WithContext(ctx context.Context) *AttachUserToTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach user to team params
func (o *AttachUserToTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach user to team params
func (o *AttachUserToTeamParams) WithHTTPClient(client *http.Client) *AttachUserToTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach user to team params
func (o *AttachUserToTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the attach user to team params
func (o *AttachUserToTeamParams) WithUserID(userID string) *AttachUserToTeamParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the attach user to team params
func (o *AttachUserToTeamParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserTeam adds the userTeam to the attach user to team params
func (o *AttachUserToTeamParams) WithUserTeam(userTeam *models.UserTeamParams) *AttachUserToTeamParams {
	o.SetUserTeam(userTeam)
	return o
}

// SetUserTeam adds the userTeam to the attach user to team params
func (o *AttachUserToTeamParams) SetUserTeam(userTeam *models.UserTeamParams) {
	o.UserTeam = userTeam
}

// WriteToRequest writes these params to a swagger request
func (o *AttachUserToTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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