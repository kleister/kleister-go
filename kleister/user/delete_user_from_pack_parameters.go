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

// NewDeleteUserFromPackParams creates a new DeleteUserFromPackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserFromPackParams() *DeleteUserFromPackParams {
	return &DeleteUserFromPackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserFromPackParamsWithTimeout creates a new DeleteUserFromPackParams object
// with the ability to set a timeout on a request.
func NewDeleteUserFromPackParamsWithTimeout(timeout time.Duration) *DeleteUserFromPackParams {
	return &DeleteUserFromPackParams{
		timeout: timeout,
	}
}

// NewDeleteUserFromPackParamsWithContext creates a new DeleteUserFromPackParams object
// with the ability to set a context for a request.
func NewDeleteUserFromPackParamsWithContext(ctx context.Context) *DeleteUserFromPackParams {
	return &DeleteUserFromPackParams{
		Context: ctx,
	}
}

// NewDeleteUserFromPackParamsWithHTTPClient creates a new DeleteUserFromPackParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserFromPackParamsWithHTTPClient(client *http.Client) *DeleteUserFromPackParams {
	return &DeleteUserFromPackParams{
		HTTPClient: client,
	}
}

/*
DeleteUserFromPackParams contains all the parameters to send to the API endpoint

	for the delete user from pack operation.

	Typically these are written to a http.Request.
*/
type DeleteUserFromPackParams struct {

	/* UserID.

	   A user UUID or slug
	*/
	UserID string

	/* UserPack.

	   The user pack data to delete
	*/
	UserPack *models.UserPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user from pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserFromPackParams) WithDefaults() *DeleteUserFromPackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user from pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserFromPackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user from pack params
func (o *DeleteUserFromPackParams) WithTimeout(timeout time.Duration) *DeleteUserFromPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user from pack params
func (o *DeleteUserFromPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user from pack params
func (o *DeleteUserFromPackParams) WithContext(ctx context.Context) *DeleteUserFromPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user from pack params
func (o *DeleteUserFromPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user from pack params
func (o *DeleteUserFromPackParams) WithHTTPClient(client *http.Client) *DeleteUserFromPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user from pack params
func (o *DeleteUserFromPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user from pack params
func (o *DeleteUserFromPackParams) WithUserID(userID string) *DeleteUserFromPackParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user from pack params
func (o *DeleteUserFromPackParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserPack adds the userPack to the delete user from pack params
func (o *DeleteUserFromPackParams) WithUserPack(userPack *models.UserPackParams) *DeleteUserFromPackParams {
	o.SetUserPack(userPack)
	return o
}

// SetUserPack adds the userPack to the delete user from pack params
func (o *DeleteUserFromPackParams) SetUserPack(userPack *models.UserPackParams) {
	o.UserPack = userPack
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserFromPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}
	if o.UserPack != nil {
		if err := r.SetBodyParam(o.UserPack); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
