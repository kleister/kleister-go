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

// NewAttachUserToPackParams creates a new AttachUserToPackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachUserToPackParams() *AttachUserToPackParams {
	return &AttachUserToPackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachUserToPackParamsWithTimeout creates a new AttachUserToPackParams object
// with the ability to set a timeout on a request.
func NewAttachUserToPackParamsWithTimeout(timeout time.Duration) *AttachUserToPackParams {
	return &AttachUserToPackParams{
		timeout: timeout,
	}
}

// NewAttachUserToPackParamsWithContext creates a new AttachUserToPackParams object
// with the ability to set a context for a request.
func NewAttachUserToPackParamsWithContext(ctx context.Context) *AttachUserToPackParams {
	return &AttachUserToPackParams{
		Context: ctx,
	}
}

// NewAttachUserToPackParamsWithHTTPClient creates a new AttachUserToPackParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachUserToPackParamsWithHTTPClient(client *http.Client) *AttachUserToPackParams {
	return &AttachUserToPackParams{
		HTTPClient: client,
	}
}

/*
AttachUserToPackParams contains all the parameters to send to the API endpoint

	for the attach user to pack operation.

	Typically these are written to a http.Request.
*/
type AttachUserToPackParams struct {

	/* UserID.

	   A user UUID or slug
	*/
	UserID string

	/* UserPack.

	   The user pack data to assign
	*/
	UserPack *models.UserPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attach user to pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachUserToPackParams) WithDefaults() *AttachUserToPackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attach user to pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachUserToPackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attach user to pack params
func (o *AttachUserToPackParams) WithTimeout(timeout time.Duration) *AttachUserToPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach user to pack params
func (o *AttachUserToPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach user to pack params
func (o *AttachUserToPackParams) WithContext(ctx context.Context) *AttachUserToPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach user to pack params
func (o *AttachUserToPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach user to pack params
func (o *AttachUserToPackParams) WithHTTPClient(client *http.Client) *AttachUserToPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach user to pack params
func (o *AttachUserToPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the attach user to pack params
func (o *AttachUserToPackParams) WithUserID(userID string) *AttachUserToPackParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the attach user to pack params
func (o *AttachUserToPackParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserPack adds the userPack to the attach user to pack params
func (o *AttachUserToPackParams) WithUserPack(userPack *models.UserPackParams) *AttachUserToPackParams {
	o.SetUserPack(userPack)
	return o
}

// SetUserPack adds the userPack to the attach user to pack params
func (o *AttachUserToPackParams) SetUserPack(userPack *models.UserPackParams) {
	o.UserPack = userPack
}

// WriteToRequest writes these params to a swagger request
func (o *AttachUserToPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
