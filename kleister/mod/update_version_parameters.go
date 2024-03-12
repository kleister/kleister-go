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

	"github.com/kleister/kleister-go/models"
)

// NewUpdateVersionParams creates a new UpdateVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVersionParams() *UpdateVersionParams {
	return &UpdateVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVersionParamsWithTimeout creates a new UpdateVersionParams object
// with the ability to set a timeout on a request.
func NewUpdateVersionParamsWithTimeout(timeout time.Duration) *UpdateVersionParams {
	return &UpdateVersionParams{
		timeout: timeout,
	}
}

// NewUpdateVersionParamsWithContext creates a new UpdateVersionParams object
// with the ability to set a context for a request.
func NewUpdateVersionParamsWithContext(ctx context.Context) *UpdateVersionParams {
	return &UpdateVersionParams{
		Context: ctx,
	}
}

// NewUpdateVersionParamsWithHTTPClient creates a new UpdateVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVersionParamsWithHTTPClient(client *http.Client) *UpdateVersionParams {
	return &UpdateVersionParams{
		HTTPClient: client,
	}
}

/*
UpdateVersionParams contains all the parameters to send to the API endpoint

	for the update version operation.

	Typically these are written to a http.Request.
*/
type UpdateVersionParams struct {

	/* ModID.

	   A mod UUID or slug
	*/
	ModID string

	/* Version.

	   The version data to update
	*/
	Version *models.Version

	/* VersionID.

	   A version UUID or slug
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVersionParams) WithDefaults() *UpdateVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update version params
func (o *UpdateVersionParams) WithTimeout(timeout time.Duration) *UpdateVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update version params
func (o *UpdateVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update version params
func (o *UpdateVersionParams) WithContext(ctx context.Context) *UpdateVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update version params
func (o *UpdateVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update version params
func (o *UpdateVersionParams) WithHTTPClient(client *http.Client) *UpdateVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update version params
func (o *UpdateVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the update version params
func (o *UpdateVersionParams) WithModID(modID string) *UpdateVersionParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the update version params
func (o *UpdateVersionParams) SetModID(modID string) {
	o.ModID = modID
}

// WithVersion adds the version to the update version params
func (o *UpdateVersionParams) WithVersion(version *models.Version) *UpdateVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update version params
func (o *UpdateVersionParams) SetVersion(version *models.Version) {
	o.Version = version
}

// WithVersionID adds the versionID to the update version params
func (o *UpdateVersionParams) WithVersionID(versionID string) *UpdateVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the update version params
func (o *UpdateVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}
	if o.Version != nil {
		if err := r.SetBodyParam(o.Version); err != nil {
			return err
		}
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
