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

// NewAppendBuildToVersionParams creates a new AppendBuildToVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppendBuildToVersionParams() *AppendBuildToVersionParams {
	return &AppendBuildToVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppendBuildToVersionParamsWithTimeout creates a new AppendBuildToVersionParams object
// with the ability to set a timeout on a request.
func NewAppendBuildToVersionParamsWithTimeout(timeout time.Duration) *AppendBuildToVersionParams {
	return &AppendBuildToVersionParams{
		timeout: timeout,
	}
}

// NewAppendBuildToVersionParamsWithContext creates a new AppendBuildToVersionParams object
// with the ability to set a context for a request.
func NewAppendBuildToVersionParamsWithContext(ctx context.Context) *AppendBuildToVersionParams {
	return &AppendBuildToVersionParams{
		Context: ctx,
	}
}

// NewAppendBuildToVersionParamsWithHTTPClient creates a new AppendBuildToVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppendBuildToVersionParamsWithHTTPClient(client *http.Client) *AppendBuildToVersionParams {
	return &AppendBuildToVersionParams{
		HTTPClient: client,
	}
}

/* AppendBuildToVersionParams contains all the parameters to send to the API endpoint
   for the append build to version operation.

   Typically these are written to a http.Request.
*/
type AppendBuildToVersionParams struct {

	/* BuildID.

	   A build UUID or slug
	*/
	BuildID string

	/* BuildVersion.

	   The version data to append to build
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

// WithDefaults hydrates default values in the append build to version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppendBuildToVersionParams) WithDefaults() *AppendBuildToVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the append build to version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppendBuildToVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the append build to version params
func (o *AppendBuildToVersionParams) WithTimeout(timeout time.Duration) *AppendBuildToVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append build to version params
func (o *AppendBuildToVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append build to version params
func (o *AppendBuildToVersionParams) WithContext(ctx context.Context) *AppendBuildToVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append build to version params
func (o *AppendBuildToVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append build to version params
func (o *AppendBuildToVersionParams) WithHTTPClient(client *http.Client) *AppendBuildToVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append build to version params
func (o *AppendBuildToVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the append build to version params
func (o *AppendBuildToVersionParams) WithBuildID(buildID string) *AppendBuildToVersionParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the append build to version params
func (o *AppendBuildToVersionParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithBuildVersion adds the buildVersion to the append build to version params
func (o *AppendBuildToVersionParams) WithBuildVersion(buildVersion *models.BuildVersionParams) *AppendBuildToVersionParams {
	o.SetBuildVersion(buildVersion)
	return o
}

// SetBuildVersion adds the buildVersion to the append build to version params
func (o *AppendBuildToVersionParams) SetBuildVersion(buildVersion *models.BuildVersionParams) {
	o.BuildVersion = buildVersion
}

// WithPackID adds the packID to the append build to version params
func (o *AppendBuildToVersionParams) WithPackID(packID string) *AppendBuildToVersionParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the append build to version params
func (o *AppendBuildToVersionParams) SetPackID(packID string) {
	o.PackID = packID
}

// WriteToRequest writes these params to a swagger request
func (o *AppendBuildToVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
