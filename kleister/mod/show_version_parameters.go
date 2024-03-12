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

// NewShowVersionParams creates a new ShowVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowVersionParams() *ShowVersionParams {
	return &ShowVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowVersionParamsWithTimeout creates a new ShowVersionParams object
// with the ability to set a timeout on a request.
func NewShowVersionParamsWithTimeout(timeout time.Duration) *ShowVersionParams {
	return &ShowVersionParams{
		timeout: timeout,
	}
}

// NewShowVersionParamsWithContext creates a new ShowVersionParams object
// with the ability to set a context for a request.
func NewShowVersionParamsWithContext(ctx context.Context) *ShowVersionParams {
	return &ShowVersionParams{
		Context: ctx,
	}
}

// NewShowVersionParamsWithHTTPClient creates a new ShowVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowVersionParamsWithHTTPClient(client *http.Client) *ShowVersionParams {
	return &ShowVersionParams{
		HTTPClient: client,
	}
}

/*
ShowVersionParams contains all the parameters to send to the API endpoint

	for the show version operation.

	Typically these are written to a http.Request.
*/
type ShowVersionParams struct {

	/* ModID.

	   A mod UUID or slug
	*/
	ModID string

	/* VersionID.

	   A version UUID or slug
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the show version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowVersionParams) WithDefaults() *ShowVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the show version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the show version params
func (o *ShowVersionParams) WithTimeout(timeout time.Duration) *ShowVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show version params
func (o *ShowVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show version params
func (o *ShowVersionParams) WithContext(ctx context.Context) *ShowVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show version params
func (o *ShowVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show version params
func (o *ShowVersionParams) WithHTTPClient(client *http.Client) *ShowVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show version params
func (o *ShowVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the show version params
func (o *ShowVersionParams) WithModID(modID string) *ShowVersionParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the show version params
func (o *ShowVersionParams) SetModID(modID string) {
	o.ModID = modID
}

// WithVersionID adds the versionID to the show version params
func (o *ShowVersionParams) WithVersionID(versionID string) *ShowVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the show version params
func (o *ShowVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *ShowVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}

	// path param version_id
	if err := r.SetPathParam("version_id", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
