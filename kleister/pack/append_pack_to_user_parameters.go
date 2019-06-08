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

// NewAppendPackToUserParams creates a new AppendPackToUserParams object
// with the default values initialized.
func NewAppendPackToUserParams() *AppendPackToUserParams {
	var ()
	return &AppendPackToUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendPackToUserParamsWithTimeout creates a new AppendPackToUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendPackToUserParamsWithTimeout(timeout time.Duration) *AppendPackToUserParams {
	var ()
	return &AppendPackToUserParams{

		timeout: timeout,
	}
}

// NewAppendPackToUserParamsWithContext creates a new AppendPackToUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendPackToUserParamsWithContext(ctx context.Context) *AppendPackToUserParams {
	var ()
	return &AppendPackToUserParams{

		Context: ctx,
	}
}

// NewAppendPackToUserParamsWithHTTPClient creates a new AppendPackToUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendPackToUserParamsWithHTTPClient(client *http.Client) *AppendPackToUserParams {
	var ()
	return &AppendPackToUserParams{
		HTTPClient: client,
	}
}

/*AppendPackToUserParams contains all the parameters to send to the API endpoint
for the append pack to user operation typically these are written to a http.Request
*/
type AppendPackToUserParams struct {

	/*PackID
	  A pack UUID or slug

	*/
	PackID string
	/*PackUser
	  The pack user data to assign

	*/
	PackUser *models.PackUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the append pack to user params
func (o *AppendPackToUserParams) WithTimeout(timeout time.Duration) *AppendPackToUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append pack to user params
func (o *AppendPackToUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append pack to user params
func (o *AppendPackToUserParams) WithContext(ctx context.Context) *AppendPackToUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append pack to user params
func (o *AppendPackToUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append pack to user params
func (o *AppendPackToUserParams) WithHTTPClient(client *http.Client) *AppendPackToUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append pack to user params
func (o *AppendPackToUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the append pack to user params
func (o *AppendPackToUserParams) WithPackID(packID string) *AppendPackToUserParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the append pack to user params
func (o *AppendPackToUserParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithPackUser adds the packUser to the append pack to user params
func (o *AppendPackToUserParams) WithPackUser(packUser *models.PackUserParams) *AppendPackToUserParams {
	o.SetPackUser(packUser)
	return o
}

// SetPackUser adds the packUser to the append pack to user params
func (o *AppendPackToUserParams) SetPackUser(packUser *models.PackUserParams) {
	o.PackUser = packUser
}

// WriteToRequest writes these params to a swagger request
func (o *AppendPackToUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pack_id
	if err := r.SetPathParam("pack_id", o.PackID); err != nil {
		return err
	}

	if o.PackUser != nil {
		if err := r.SetBodyParam(o.PackUser); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
