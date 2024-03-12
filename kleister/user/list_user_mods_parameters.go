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
)

// NewListUserModsParams creates a new ListUserModsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUserModsParams() *ListUserModsParams {
	return &ListUserModsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUserModsParamsWithTimeout creates a new ListUserModsParams object
// with the ability to set a timeout on a request.
func NewListUserModsParamsWithTimeout(timeout time.Duration) *ListUserModsParams {
	return &ListUserModsParams{
		timeout: timeout,
	}
}

// NewListUserModsParamsWithContext creates a new ListUserModsParams object
// with the ability to set a context for a request.
func NewListUserModsParamsWithContext(ctx context.Context) *ListUserModsParams {
	return &ListUserModsParams{
		Context: ctx,
	}
}

// NewListUserModsParamsWithHTTPClient creates a new ListUserModsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListUserModsParamsWithHTTPClient(client *http.Client) *ListUserModsParams {
	return &ListUserModsParams{
		HTTPClient: client,
	}
}

/*
ListUserModsParams contains all the parameters to send to the API endpoint

	for the list user mods operation.

	Typically these are written to a http.Request.
*/
type ListUserModsParams struct {

	/* UserID.

	   A user UUID or slug
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list user mods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUserModsParams) WithDefaults() *ListUserModsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list user mods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUserModsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list user mods params
func (o *ListUserModsParams) WithTimeout(timeout time.Duration) *ListUserModsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user mods params
func (o *ListUserModsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user mods params
func (o *ListUserModsParams) WithContext(ctx context.Context) *ListUserModsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user mods params
func (o *ListUserModsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user mods params
func (o *ListUserModsParams) WithHTTPClient(client *http.Client) *ListUserModsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user mods params
func (o *ListUserModsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the list user mods params
func (o *ListUserModsParams) WithUserID(userID string) *ListUserModsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the list user mods params
func (o *ListUserModsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserModsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
