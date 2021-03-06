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

// NewListPackUsersParams creates a new ListPackUsersParams object
// with the default values initialized.
func NewListPackUsersParams() *ListPackUsersParams {
	var ()
	return &ListPackUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPackUsersParamsWithTimeout creates a new ListPackUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPackUsersParamsWithTimeout(timeout time.Duration) *ListPackUsersParams {
	var ()
	return &ListPackUsersParams{

		timeout: timeout,
	}
}

// NewListPackUsersParamsWithContext creates a new ListPackUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPackUsersParamsWithContext(ctx context.Context) *ListPackUsersParams {
	var ()
	return &ListPackUsersParams{

		Context: ctx,
	}
}

// NewListPackUsersParamsWithHTTPClient creates a new ListPackUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPackUsersParamsWithHTTPClient(client *http.Client) *ListPackUsersParams {
	var ()
	return &ListPackUsersParams{
		HTTPClient: client,
	}
}

/*ListPackUsersParams contains all the parameters to send to the API endpoint
for the list pack users operation typically these are written to a http.Request
*/
type ListPackUsersParams struct {

	/*PackID
	  A pack UUID or slug

	*/
	PackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list pack users params
func (o *ListPackUsersParams) WithTimeout(timeout time.Duration) *ListPackUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list pack users params
func (o *ListPackUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list pack users params
func (o *ListPackUsersParams) WithContext(ctx context.Context) *ListPackUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list pack users params
func (o *ListPackUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list pack users params
func (o *ListPackUsersParams) WithHTTPClient(client *http.Client) *ListPackUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list pack users params
func (o *ListPackUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the list pack users params
func (o *ListPackUsersParams) WithPackID(packID string) *ListPackUsersParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the list pack users params
func (o *ListPackUsersParams) SetPackID(packID string) {
	o.PackID = packID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPackUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
