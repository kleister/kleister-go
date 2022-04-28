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

	"github.com/kleister/kleister-go/v1/models"
)

// NewDeleteVersionFromBuildParams creates a new DeleteVersionFromBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVersionFromBuildParams() *DeleteVersionFromBuildParams {
	return &DeleteVersionFromBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVersionFromBuildParamsWithTimeout creates a new DeleteVersionFromBuildParams object
// with the ability to set a timeout on a request.
func NewDeleteVersionFromBuildParamsWithTimeout(timeout time.Duration) *DeleteVersionFromBuildParams {
	return &DeleteVersionFromBuildParams{
		timeout: timeout,
	}
}

// NewDeleteVersionFromBuildParamsWithContext creates a new DeleteVersionFromBuildParams object
// with the ability to set a context for a request.
func NewDeleteVersionFromBuildParamsWithContext(ctx context.Context) *DeleteVersionFromBuildParams {
	return &DeleteVersionFromBuildParams{
		Context: ctx,
	}
}

// NewDeleteVersionFromBuildParamsWithHTTPClient creates a new DeleteVersionFromBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVersionFromBuildParamsWithHTTPClient(client *http.Client) *DeleteVersionFromBuildParams {
	return &DeleteVersionFromBuildParams{
		HTTPClient: client,
	}
}

/* DeleteVersionFromBuildParams contains all the parameters to send to the API endpoint
   for the delete version from build operation.

   Typically these are written to a http.Request.
*/
type DeleteVersionFromBuildParams struct {

	/* ModID.

	   A mod UUID or slug
	*/
	ModID string

	/* VersionBuild.

	   The build data to unlink from version
	*/
	VersionBuild *models.VersionBuildParams

	/* VersionID.

	   A version UUID or slug
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete version from build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionFromBuildParams) WithDefaults() *DeleteVersionFromBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete version from build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionFromBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete version from build params
func (o *DeleteVersionFromBuildParams) WithTimeout(timeout time.Duration) *DeleteVersionFromBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete version from build params
func (o *DeleteVersionFromBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete version from build params
func (o *DeleteVersionFromBuildParams) WithContext(ctx context.Context) *DeleteVersionFromBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete version from build params
func (o *DeleteVersionFromBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete version from build params
func (o *DeleteVersionFromBuildParams) WithHTTPClient(client *http.Client) *DeleteVersionFromBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete version from build params
func (o *DeleteVersionFromBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the delete version from build params
func (o *DeleteVersionFromBuildParams) WithModID(modID string) *DeleteVersionFromBuildParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the delete version from build params
func (o *DeleteVersionFromBuildParams) SetModID(modID string) {
	o.ModID = modID
}

// WithVersionBuild adds the versionBuild to the delete version from build params
func (o *DeleteVersionFromBuildParams) WithVersionBuild(versionBuild *models.VersionBuildParams) *DeleteVersionFromBuildParams {
	o.SetVersionBuild(versionBuild)
	return o
}

// SetVersionBuild adds the versionBuild to the delete version from build params
func (o *DeleteVersionFromBuildParams) SetVersionBuild(versionBuild *models.VersionBuildParams) {
	o.VersionBuild = versionBuild
}

// WithVersionID adds the versionID to the delete version from build params
func (o *DeleteVersionFromBuildParams) WithVersionID(versionID string) *DeleteVersionFromBuildParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete version from build params
func (o *DeleteVersionFromBuildParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVersionFromBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}
	if o.VersionBuild != nil {
		if err := r.SetBodyParam(o.VersionBuild); err != nil {
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
