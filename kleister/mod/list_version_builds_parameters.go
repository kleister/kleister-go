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

// NewListVersionBuildsParams creates a new ListVersionBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVersionBuildsParams() *ListVersionBuildsParams {
	return &ListVersionBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVersionBuildsParamsWithTimeout creates a new ListVersionBuildsParams object
// with the ability to set a timeout on a request.
func NewListVersionBuildsParamsWithTimeout(timeout time.Duration) *ListVersionBuildsParams {
	return &ListVersionBuildsParams{
		timeout: timeout,
	}
}

// NewListVersionBuildsParamsWithContext creates a new ListVersionBuildsParams object
// with the ability to set a context for a request.
func NewListVersionBuildsParamsWithContext(ctx context.Context) *ListVersionBuildsParams {
	return &ListVersionBuildsParams{
		Context: ctx,
	}
}

// NewListVersionBuildsParamsWithHTTPClient creates a new ListVersionBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVersionBuildsParamsWithHTTPClient(client *http.Client) *ListVersionBuildsParams {
	return &ListVersionBuildsParams{
		HTTPClient: client,
	}
}

/*
ListVersionBuildsParams contains all the parameters to send to the API endpoint

	for the list version builds operation.

	Typically these are written to a http.Request.
*/
type ListVersionBuildsParams struct {

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

// WithDefaults hydrates default values in the list version builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionBuildsParams) WithDefaults() *ListVersionBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list version builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionBuildsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list version builds params
func (o *ListVersionBuildsParams) WithTimeout(timeout time.Duration) *ListVersionBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list version builds params
func (o *ListVersionBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list version builds params
func (o *ListVersionBuildsParams) WithContext(ctx context.Context) *ListVersionBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list version builds params
func (o *ListVersionBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list version builds params
func (o *ListVersionBuildsParams) WithHTTPClient(client *http.Client) *ListVersionBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list version builds params
func (o *ListVersionBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the list version builds params
func (o *ListVersionBuildsParams) WithModID(modID string) *ListVersionBuildsParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the list version builds params
func (o *ListVersionBuildsParams) SetModID(modID string) {
	o.ModID = modID
}

// WithVersionID adds the versionID to the list version builds params
func (o *ListVersionBuildsParams) WithVersionID(versionID string) *ListVersionBuildsParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the list version builds params
func (o *ListVersionBuildsParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *ListVersionBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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