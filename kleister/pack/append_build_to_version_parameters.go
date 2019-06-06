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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// NewAppendBuildToVersionParams creates a new AppendBuildToVersionParams object
// with the default values initialized.
func NewAppendBuildToVersionParams() *AppendBuildToVersionParams {
	var ()
	return &AppendBuildToVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendBuildToVersionParamsWithTimeout creates a new AppendBuildToVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendBuildToVersionParamsWithTimeout(timeout time.Duration) *AppendBuildToVersionParams {
	var ()
	return &AppendBuildToVersionParams{

		timeout: timeout,
	}
}

// NewAppendBuildToVersionParamsWithContext creates a new AppendBuildToVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendBuildToVersionParamsWithContext(ctx context.Context) *AppendBuildToVersionParams {
	var ()
	return &AppendBuildToVersionParams{

		Context: ctx,
	}
}

// NewAppendBuildToVersionParamsWithHTTPClient creates a new AppendBuildToVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendBuildToVersionParamsWithHTTPClient(client *http.Client) *AppendBuildToVersionParams {
	var ()
	return &AppendBuildToVersionParams{
		HTTPClient: client,
	}
}

/*AppendBuildToVersionParams contains all the parameters to send to the API endpoint
for the append build to version operation typically these are written to a http.Request
*/
type AppendBuildToVersionParams struct {

	/*BuildID
	  A build UUID or slug

	*/
	BuildID string
	/*PackID
	  A pack UUID or slug

	*/
	PackID string
	/*Params
	  The version data to append to build

	*/
	Params *models.BuildVersionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
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

// WithPackID adds the packID to the append build to version params
func (o *AppendBuildToVersionParams) WithPackID(packID string) *AppendBuildToVersionParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the append build to version params
func (o *AppendBuildToVersionParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithParams adds the params to the append build to version params
func (o *AppendBuildToVersionParams) WithParams(params *models.BuildVersionParams) *AppendBuildToVersionParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the append build to version params
func (o *AppendBuildToVersionParams) SetParams(params *models.BuildVersionParams) {
	o.Params = params
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

	// path param pack_id
	if err := r.SetPathParam("pack_id", o.PackID); err != nil {
		return err
	}

	if o.Params != nil {
		if err := r.SetBodyParam(o.Params); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
