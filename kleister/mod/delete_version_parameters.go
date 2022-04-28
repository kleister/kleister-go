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

// NewDeleteVersionParams creates a new DeleteVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVersionParams() *DeleteVersionParams {
	return &DeleteVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVersionParamsWithTimeout creates a new DeleteVersionParams object
// with the ability to set a timeout on a request.
func NewDeleteVersionParamsWithTimeout(timeout time.Duration) *DeleteVersionParams {
	return &DeleteVersionParams{
		timeout: timeout,
	}
}

// NewDeleteVersionParamsWithContext creates a new DeleteVersionParams object
// with the ability to set a context for a request.
func NewDeleteVersionParamsWithContext(ctx context.Context) *DeleteVersionParams {
	return &DeleteVersionParams{
		Context: ctx,
	}
}

// NewDeleteVersionParamsWithHTTPClient creates a new DeleteVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVersionParamsWithHTTPClient(client *http.Client) *DeleteVersionParams {
	return &DeleteVersionParams{
		HTTPClient: client,
	}
}

/* DeleteVersionParams contains all the parameters to send to the API endpoint
   for the delete version operation.

   Typically these are written to a http.Request.
*/
type DeleteVersionParams struct {

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

// WithDefaults hydrates default values in the delete version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionParams) WithDefaults() *DeleteVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete version params
func (o *DeleteVersionParams) WithTimeout(timeout time.Duration) *DeleteVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete version params
func (o *DeleteVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete version params
func (o *DeleteVersionParams) WithContext(ctx context.Context) *DeleteVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete version params
func (o *DeleteVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete version params
func (o *DeleteVersionParams) WithHTTPClient(client *http.Client) *DeleteVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete version params
func (o *DeleteVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the delete version params
func (o *DeleteVersionParams) WithModID(modID string) *DeleteVersionParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the delete version params
func (o *DeleteVersionParams) SetModID(modID string) {
	o.ModID = modID
}

// WithVersionID adds the versionID to the delete version params
func (o *DeleteVersionParams) WithVersionID(versionID string) *DeleteVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete version params
func (o *DeleteVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
