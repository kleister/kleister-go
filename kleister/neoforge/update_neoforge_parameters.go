// Code generated by go-swagger; DO NOT EDIT.

package neoforge

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

// NewUpdateNeoforgeParams creates a new UpdateNeoforgeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNeoforgeParams() *UpdateNeoforgeParams {
	return &UpdateNeoforgeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNeoforgeParamsWithTimeout creates a new UpdateNeoforgeParams object
// with the ability to set a timeout on a request.
func NewUpdateNeoforgeParamsWithTimeout(timeout time.Duration) *UpdateNeoforgeParams {
	return &UpdateNeoforgeParams{
		timeout: timeout,
	}
}

// NewUpdateNeoforgeParamsWithContext creates a new UpdateNeoforgeParams object
// with the ability to set a context for a request.
func NewUpdateNeoforgeParamsWithContext(ctx context.Context) *UpdateNeoforgeParams {
	return &UpdateNeoforgeParams{
		Context: ctx,
	}
}

// NewUpdateNeoforgeParamsWithHTTPClient creates a new UpdateNeoforgeParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNeoforgeParamsWithHTTPClient(client *http.Client) *UpdateNeoforgeParams {
	return &UpdateNeoforgeParams{
		HTTPClient: client,
	}
}

/*
UpdateNeoforgeParams contains all the parameters to send to the API endpoint

	for the update neoforge operation.

	Typically these are written to a http.Request.
*/
type UpdateNeoforgeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update neoforge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNeoforgeParams) WithDefaults() *UpdateNeoforgeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update neoforge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNeoforgeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update neoforge params
func (o *UpdateNeoforgeParams) WithTimeout(timeout time.Duration) *UpdateNeoforgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update neoforge params
func (o *UpdateNeoforgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update neoforge params
func (o *UpdateNeoforgeParams) WithContext(ctx context.Context) *UpdateNeoforgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update neoforge params
func (o *UpdateNeoforgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update neoforge params
func (o *UpdateNeoforgeParams) WithHTTPClient(client *http.Client) *UpdateNeoforgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update neoforge params
func (o *UpdateNeoforgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNeoforgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
