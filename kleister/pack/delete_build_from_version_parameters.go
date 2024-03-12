// Code generated by go-swagger; DO NOT EDIT.

package pack

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

// NewDeleteBuildFromVersionParams creates a new DeleteBuildFromVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBuildFromVersionParams() *DeleteBuildFromVersionParams {
	return &DeleteBuildFromVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildFromVersionParamsWithTimeout creates a new DeleteBuildFromVersionParams object
// with the ability to set a timeout on a request.
func NewDeleteBuildFromVersionParamsWithTimeout(timeout time.Duration) *DeleteBuildFromVersionParams {
	return &DeleteBuildFromVersionParams{
		timeout: timeout,
	}
}

// NewDeleteBuildFromVersionParamsWithContext creates a new DeleteBuildFromVersionParams object
// with the ability to set a context for a request.
func NewDeleteBuildFromVersionParamsWithContext(ctx context.Context) *DeleteBuildFromVersionParams {
	return &DeleteBuildFromVersionParams{
		Context: ctx,
	}
}

// NewDeleteBuildFromVersionParamsWithHTTPClient creates a new DeleteBuildFromVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBuildFromVersionParamsWithHTTPClient(client *http.Client) *DeleteBuildFromVersionParams {
	return &DeleteBuildFromVersionParams{
		HTTPClient: client,
	}
}

/*
DeleteBuildFromVersionParams contains all the parameters to send to the API endpoint

	for the delete build from version operation.

	Typically these are written to a http.Request.
*/
type DeleteBuildFromVersionParams struct {

	/* BuildID.

	   A build UUID or slug
	*/
	BuildID string

	/* BuildVersion.

	   The version data to unlink from build
	*/
	BuildVersion *models.BuildVersionParams

	/* PackID.

	   A pack UUID or slug
	*/
	PackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete build from version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildFromVersionParams) WithDefaults() *DeleteBuildFromVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete build from version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBuildFromVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete build from version params
func (o *DeleteBuildFromVersionParams) WithTimeout(timeout time.Duration) *DeleteBuildFromVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build from version params
func (o *DeleteBuildFromVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build from version params
func (o *DeleteBuildFromVersionParams) WithContext(ctx context.Context) *DeleteBuildFromVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build from version params
func (o *DeleteBuildFromVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build from version params
func (o *DeleteBuildFromVersionParams) WithHTTPClient(client *http.Client) *DeleteBuildFromVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build from version params
func (o *DeleteBuildFromVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the delete build from version params
func (o *DeleteBuildFromVersionParams) WithBuildID(buildID string) *DeleteBuildFromVersionParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the delete build from version params
func (o *DeleteBuildFromVersionParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithBuildVersion adds the buildVersion to the delete build from version params
func (o *DeleteBuildFromVersionParams) WithBuildVersion(buildVersion *models.BuildVersionParams) *DeleteBuildFromVersionParams {
	o.SetBuildVersion(buildVersion)
	return o
}

// SetBuildVersion adds the buildVersion to the delete build from version params
func (o *DeleteBuildFromVersionParams) SetBuildVersion(buildVersion *models.BuildVersionParams) {
	o.BuildVersion = buildVersion
}

// WithPackID adds the packID to the delete build from version params
func (o *DeleteBuildFromVersionParams) WithPackID(packID string) *DeleteBuildFromVersionParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the delete build from version params
func (o *DeleteBuildFromVersionParams) SetPackID(packID string) {
	o.PackID = packID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildFromVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}
	if o.BuildVersion != nil {
		if err := r.SetBodyParam(o.BuildVersion); err != nil {
			return err
		}
	}

	// path param pack_id
	if err := r.SetPathParam("pack_id", o.PackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
