// Code generated by go-swagger; DO NOT EDIT.

package mod

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

// NewListVersionsParams creates a new ListVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVersionsParams() *ListVersionsParams {
	return &ListVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVersionsParamsWithTimeout creates a new ListVersionsParams object
// with the ability to set a timeout on a request.
func NewListVersionsParamsWithTimeout(timeout time.Duration) *ListVersionsParams {
	return &ListVersionsParams{
		timeout: timeout,
	}
}

// NewListVersionsParamsWithContext creates a new ListVersionsParams object
// with the ability to set a context for a request.
func NewListVersionsParamsWithContext(ctx context.Context) *ListVersionsParams {
	return &ListVersionsParams{
		Context: ctx,
	}
}

// NewListVersionsParamsWithHTTPClient creates a new ListVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVersionsParamsWithHTTPClient(client *http.Client) *ListVersionsParams {
	return &ListVersionsParams{
		HTTPClient: client,
	}
}

/*
ListVersionsParams contains all the parameters to send to the API endpoint

	for the list versions operation.

	Typically these are written to a http.Request.
*/
type ListVersionsParams struct {

	/* ModID.

	   A mod UUID or slug
	*/
	ModID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionsParams) WithDefaults() *ListVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list versions params
func (o *ListVersionsParams) WithTimeout(timeout time.Duration) *ListVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list versions params
func (o *ListVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list versions params
func (o *ListVersionsParams) WithContext(ctx context.Context) *ListVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list versions params
func (o *ListVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list versions params
func (o *ListVersionsParams) WithHTTPClient(client *http.Client) *ListVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list versions params
func (o *ListVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the list versions params
func (o *ListVersionsParams) WithModID(modID string) *ListVersionsParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the list versions params
func (o *ListVersionsParams) SetModID(modID string) {
	o.ModID = modID
}

// WriteToRequest writes these params to a swagger request
func (o *ListVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
