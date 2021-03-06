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

// NewListBuildsParams creates a new ListBuildsParams object
// with the default values initialized.
func NewListBuildsParams() *ListBuildsParams {
	var ()
	return &ListBuildsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBuildsParamsWithTimeout creates a new ListBuildsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBuildsParamsWithTimeout(timeout time.Duration) *ListBuildsParams {
	var ()
	return &ListBuildsParams{

		timeout: timeout,
	}
}

// NewListBuildsParamsWithContext creates a new ListBuildsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBuildsParamsWithContext(ctx context.Context) *ListBuildsParams {
	var ()
	return &ListBuildsParams{

		Context: ctx,
	}
}

// NewListBuildsParamsWithHTTPClient creates a new ListBuildsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBuildsParamsWithHTTPClient(client *http.Client) *ListBuildsParams {
	var ()
	return &ListBuildsParams{
		HTTPClient: client,
	}
}

/*ListBuildsParams contains all the parameters to send to the API endpoint
for the list builds operation typically these are written to a http.Request
*/
type ListBuildsParams struct {

	/*PackID
	  A pack UUID or slug

	*/
	PackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list builds params
func (o *ListBuildsParams) WithTimeout(timeout time.Duration) *ListBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list builds params
func (o *ListBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list builds params
func (o *ListBuildsParams) WithContext(ctx context.Context) *ListBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list builds params
func (o *ListBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list builds params
func (o *ListBuildsParams) WithHTTPClient(client *http.Client) *ListBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list builds params
func (o *ListBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the list builds params
func (o *ListBuildsParams) WithPackID(packID string) *ListBuildsParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the list builds params
func (o *ListBuildsParams) SetPackID(packID string) {
	o.PackID = packID
}

// WriteToRequest writes these params to a swagger request
func (o *ListBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pack_id
	if err := r.SetPathParam("pack_id", o.PackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
