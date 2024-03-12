// Code generated by go-swagger; DO NOT EDIT.

package quilt

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

// NewListQuiltBuildsParams creates a new ListQuiltBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListQuiltBuildsParams() *ListQuiltBuildsParams {
	return &ListQuiltBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListQuiltBuildsParamsWithTimeout creates a new ListQuiltBuildsParams object
// with the ability to set a timeout on a request.
func NewListQuiltBuildsParamsWithTimeout(timeout time.Duration) *ListQuiltBuildsParams {
	return &ListQuiltBuildsParams{
		timeout: timeout,
	}
}

// NewListQuiltBuildsParamsWithContext creates a new ListQuiltBuildsParams object
// with the ability to set a context for a request.
func NewListQuiltBuildsParamsWithContext(ctx context.Context) *ListQuiltBuildsParams {
	return &ListQuiltBuildsParams{
		Context: ctx,
	}
}

// NewListQuiltBuildsParamsWithHTTPClient creates a new ListQuiltBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListQuiltBuildsParamsWithHTTPClient(client *http.Client) *ListQuiltBuildsParams {
	return &ListQuiltBuildsParams{
		HTTPClient: client,
	}
}

/*
ListQuiltBuildsParams contains all the parameters to send to the API endpoint

	for the list quilt builds operation.

	Typically these are written to a http.Request.
*/
type ListQuiltBuildsParams struct {

	/* QuiltID.

	   A quilt UUID or slug
	*/
	QuiltID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list quilt builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListQuiltBuildsParams) WithDefaults() *ListQuiltBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list quilt builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListQuiltBuildsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list quilt builds params
func (o *ListQuiltBuildsParams) WithTimeout(timeout time.Duration) *ListQuiltBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list quilt builds params
func (o *ListQuiltBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list quilt builds params
func (o *ListQuiltBuildsParams) WithContext(ctx context.Context) *ListQuiltBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list quilt builds params
func (o *ListQuiltBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list quilt builds params
func (o *ListQuiltBuildsParams) WithHTTPClient(client *http.Client) *ListQuiltBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list quilt builds params
func (o *ListQuiltBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuiltID adds the quiltID to the list quilt builds params
func (o *ListQuiltBuildsParams) WithQuiltID(quiltID string) *ListQuiltBuildsParams {
	o.SetQuiltID(quiltID)
	return o
}

// SetQuiltID adds the quiltId to the list quilt builds params
func (o *ListQuiltBuildsParams) SetQuiltID(quiltID string) {
	o.QuiltID = quiltID
}

// WriteToRequest writes these params to a swagger request
func (o *ListQuiltBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param quilt_id
	if err := r.SetPathParam("quilt_id", o.QuiltID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
