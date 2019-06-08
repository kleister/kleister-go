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

// NewCreateVersionParams creates a new CreateVersionParams object
// with the default values initialized.
func NewCreateVersionParams() *CreateVersionParams {
	var ()
	return &CreateVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVersionParamsWithTimeout creates a new CreateVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVersionParamsWithTimeout(timeout time.Duration) *CreateVersionParams {
	var ()
	return &CreateVersionParams{

		timeout: timeout,
	}
}

// NewCreateVersionParamsWithContext creates a new CreateVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVersionParamsWithContext(ctx context.Context) *CreateVersionParams {
	var ()
	return &CreateVersionParams{

		Context: ctx,
	}
}

// NewCreateVersionParamsWithHTTPClient creates a new CreateVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVersionParamsWithHTTPClient(client *http.Client) *CreateVersionParams {
	var ()
	return &CreateVersionParams{
		HTTPClient: client,
	}
}

/*CreateVersionParams contains all the parameters to send to the API endpoint
for the create version operation typically these are written to a http.Request
*/
type CreateVersionParams struct {

	/*ModID
	  A mod UUID or slug

	*/
	ModID string
	/*Version
	  The version data to create

	*/
	Version *models.Version

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create version params
func (o *CreateVersionParams) WithTimeout(timeout time.Duration) *CreateVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create version params
func (o *CreateVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create version params
func (o *CreateVersionParams) WithContext(ctx context.Context) *CreateVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create version params
func (o *CreateVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create version params
func (o *CreateVersionParams) WithHTTPClient(client *http.Client) *CreateVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create version params
func (o *CreateVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the create version params
func (o *CreateVersionParams) WithModID(modID string) *CreateVersionParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the create version params
func (o *CreateVersionParams) SetModID(modID string) {
	o.ModID = modID
}

// WithVersion adds the version to the create version params
func (o *CreateVersionParams) WithVersion(version *models.Version) *CreateVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create version params
func (o *CreateVersionParams) SetVersion(version *models.Version) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}

	if o.Version != nil {
		if err := r.SetBodyParam(o.Version); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
