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

// NewUpdatePackParams creates a new UpdatePackParams object
// with the default values initialized.
func NewUpdatePackParams() *UpdatePackParams {
	var ()
	return &UpdatePackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePackParamsWithTimeout creates a new UpdatePackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePackParamsWithTimeout(timeout time.Duration) *UpdatePackParams {
	var ()
	return &UpdatePackParams{

		timeout: timeout,
	}
}

// NewUpdatePackParamsWithContext creates a new UpdatePackParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePackParamsWithContext(ctx context.Context) *UpdatePackParams {
	var ()
	return &UpdatePackParams{

		Context: ctx,
	}
}

// NewUpdatePackParamsWithHTTPClient creates a new UpdatePackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePackParamsWithHTTPClient(client *http.Client) *UpdatePackParams {
	var ()
	return &UpdatePackParams{
		HTTPClient: client,
	}
}

/*UpdatePackParams contains all the parameters to send to the API endpoint
for the update pack operation typically these are written to a http.Request
*/
type UpdatePackParams struct {

	/*PackID
	  A pack UUID or slug

	*/
	PackID string
	/*Params
	  The pack data to update

	*/
	Params *models.Pack

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update pack params
func (o *UpdatePackParams) WithTimeout(timeout time.Duration) *UpdatePackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pack params
func (o *UpdatePackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pack params
func (o *UpdatePackParams) WithContext(ctx context.Context) *UpdatePackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pack params
func (o *UpdatePackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pack params
func (o *UpdatePackParams) WithHTTPClient(client *http.Client) *UpdatePackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pack params
func (o *UpdatePackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the update pack params
func (o *UpdatePackParams) WithPackID(packID string) *UpdatePackParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the update pack params
func (o *UpdatePackParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithParams adds the params to the update pack params
func (o *UpdatePackParams) WithParams(params *models.Pack) *UpdatePackParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the update pack params
func (o *UpdatePackParams) SetParams(params *models.Pack) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
