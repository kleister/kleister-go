// Code generated by go-swagger; DO NOT EDIT.

package profile

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

// NewUpdateProfileParams creates a new UpdateProfileParams object
// with the default values initialized.
func NewUpdateProfileParams() *UpdateProfileParams {
	var ()
	return &UpdateProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProfileParamsWithTimeout creates a new UpdateProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProfileParamsWithTimeout(timeout time.Duration) *UpdateProfileParams {
	var ()
	return &UpdateProfileParams{

		timeout: timeout,
	}
}

// NewUpdateProfileParamsWithContext creates a new UpdateProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProfileParamsWithContext(ctx context.Context) *UpdateProfileParams {
	var ()
	return &UpdateProfileParams{

		Context: ctx,
	}
}

// NewUpdateProfileParamsWithHTTPClient creates a new UpdateProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProfileParamsWithHTTPClient(client *http.Client) *UpdateProfileParams {
	var ()
	return &UpdateProfileParams{
		HTTPClient: client,
	}
}

/*UpdateProfileParams contains all the parameters to send to the API endpoint
for the update profile operation typically these are written to a http.Request
*/
type UpdateProfileParams struct {

	/*Profile
	  The profile data to update

	*/
	Profile *models.Profile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update profile params
func (o *UpdateProfileParams) WithTimeout(timeout time.Duration) *UpdateProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update profile params
func (o *UpdateProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update profile params
func (o *UpdateProfileParams) WithContext(ctx context.Context) *UpdateProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update profile params
func (o *UpdateProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update profile params
func (o *UpdateProfileParams) WithHTTPClient(client *http.Client) *UpdateProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update profile params
func (o *UpdateProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProfile adds the profile to the update profile params
func (o *UpdateProfileParams) WithProfile(profile *models.Profile) *UpdateProfileParams {
	o.SetProfile(profile)
	return o
}

// SetProfile adds the profile to the update profile params
func (o *UpdateProfileParams) SetProfile(profile *models.Profile) {
	o.Profile = profile
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Profile != nil {
		if err := r.SetBodyParam(o.Profile); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
