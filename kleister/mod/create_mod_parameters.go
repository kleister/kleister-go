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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// NewCreateModParams creates a new CreateModParams object
// with the default values initialized.
func NewCreateModParams() *CreateModParams {
	var ()
	return &CreateModParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateModParamsWithTimeout creates a new CreateModParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateModParamsWithTimeout(timeout time.Duration) *CreateModParams {
	var ()
	return &CreateModParams{

		timeout: timeout,
	}
}

// NewCreateModParamsWithContext creates a new CreateModParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateModParamsWithContext(ctx context.Context) *CreateModParams {
	var ()
	return &CreateModParams{

		Context: ctx,
	}
}

// NewCreateModParamsWithHTTPClient creates a new CreateModParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateModParamsWithHTTPClient(client *http.Client) *CreateModParams {
	var ()
	return &CreateModParams{
		HTTPClient: client,
	}
}

/*CreateModParams contains all the parameters to send to the API endpoint
for the create mod operation typically these are written to a http.Request
*/
type CreateModParams struct {

	/*Mod
	  The mod data to create

	*/
	Mod *models.Mod

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create mod params
func (o *CreateModParams) WithTimeout(timeout time.Duration) *CreateModParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create mod params
func (o *CreateModParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create mod params
func (o *CreateModParams) WithContext(ctx context.Context) *CreateModParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create mod params
func (o *CreateModParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create mod params
func (o *CreateModParams) WithHTTPClient(client *http.Client) *CreateModParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create mod params
func (o *CreateModParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMod adds the mod to the create mod params
func (o *CreateModParams) WithMod(mod *models.Mod) *CreateModParams {
	o.SetMod(mod)
	return o
}

// SetMod adds the mod to the create mod params
func (o *CreateModParams) SetMod(mod *models.Mod) {
	o.Mod = mod
}

// WriteToRequest writes these params to a swagger request
func (o *CreateModParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Mod != nil {
		if err := r.SetBodyParam(o.Mod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}