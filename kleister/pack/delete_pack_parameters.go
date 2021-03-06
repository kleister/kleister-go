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

// NewDeletePackParams creates a new DeletePackParams object
// with the default values initialized.
func NewDeletePackParams() *DeletePackParams {
	var ()
	return &DeletePackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePackParamsWithTimeout creates a new DeletePackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePackParamsWithTimeout(timeout time.Duration) *DeletePackParams {
	var ()
	return &DeletePackParams{

		timeout: timeout,
	}
}

// NewDeletePackParamsWithContext creates a new DeletePackParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePackParamsWithContext(ctx context.Context) *DeletePackParams {
	var ()
	return &DeletePackParams{

		Context: ctx,
	}
}

// NewDeletePackParamsWithHTTPClient creates a new DeletePackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePackParamsWithHTTPClient(client *http.Client) *DeletePackParams {
	var ()
	return &DeletePackParams{
		HTTPClient: client,
	}
}

/*DeletePackParams contains all the parameters to send to the API endpoint
for the delete pack operation typically these are written to a http.Request
*/
type DeletePackParams struct {

	/*PackID
	  A pack UUID or slug

	*/
	PackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete pack params
func (o *DeletePackParams) WithTimeout(timeout time.Duration) *DeletePackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pack params
func (o *DeletePackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pack params
func (o *DeletePackParams) WithContext(ctx context.Context) *DeletePackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pack params
func (o *DeletePackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pack params
func (o *DeletePackParams) WithHTTPClient(client *http.Client) *DeletePackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pack params
func (o *DeletePackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the delete pack params
func (o *DeletePackParams) WithPackID(packID string) *DeletePackParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the delete pack params
func (o *DeletePackParams) SetPackID(packID string) {
	o.PackID = packID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
