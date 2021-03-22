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
)

// NewDeleteBuildParams creates a new DeleteBuildParams object
// with the default values initialized.
func NewDeleteBuildParams() *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildParamsWithTimeout creates a new DeleteBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBuildParamsWithTimeout(timeout time.Duration) *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{

		timeout: timeout,
	}
}

// NewDeleteBuildParamsWithContext creates a new DeleteBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBuildParamsWithContext(ctx context.Context) *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{

		Context: ctx,
	}
}

// NewDeleteBuildParamsWithHTTPClient creates a new DeleteBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBuildParamsWithHTTPClient(client *http.Client) *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{
		HTTPClient: client,
	}
}

/*DeleteBuildParams contains all the parameters to send to the API endpoint
for the delete build operation typically these are written to a http.Request
*/
type DeleteBuildParams struct {

	/*BuildID
	  A build UUID or slug

	*/
	BuildID string
	/*PackID
	  A pack UUID or slug

	*/
	PackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete build params
func (o *DeleteBuildParams) WithTimeout(timeout time.Duration) *DeleteBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build params
func (o *DeleteBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build params
func (o *DeleteBuildParams) WithContext(ctx context.Context) *DeleteBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build params
func (o *DeleteBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build params
func (o *DeleteBuildParams) WithHTTPClient(client *http.Client) *DeleteBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build params
func (o *DeleteBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the delete build params
func (o *DeleteBuildParams) WithBuildID(buildID string) *DeleteBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the delete build params
func (o *DeleteBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithPackID adds the packID to the delete build params
func (o *DeleteBuildParams) WithPackID(packID string) *DeleteBuildParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the delete build params
func (o *DeleteBuildParams) SetPackID(packID string) {
	o.PackID = packID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}