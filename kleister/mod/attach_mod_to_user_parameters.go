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
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// NewAttachModToUserParams creates a new AttachModToUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachModToUserParams() *AttachModToUserParams {
	return &AttachModToUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachModToUserParamsWithTimeout creates a new AttachModToUserParams object
// with the ability to set a timeout on a request.
func NewAttachModToUserParamsWithTimeout(timeout time.Duration) *AttachModToUserParams {
	return &AttachModToUserParams{
		timeout: timeout,
	}
}

// NewAttachModToUserParamsWithContext creates a new AttachModToUserParams object
// with the ability to set a context for a request.
func NewAttachModToUserParamsWithContext(ctx context.Context) *AttachModToUserParams {
	return &AttachModToUserParams{
		Context: ctx,
	}
}

// NewAttachModToUserParamsWithHTTPClient creates a new AttachModToUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachModToUserParamsWithHTTPClient(client *http.Client) *AttachModToUserParams {
	return &AttachModToUserParams{
		HTTPClient: client,
	}
}

/*
AttachModToUserParams contains all the parameters to send to the API endpoint

	for the attach mod to user operation.

	Typically these are written to a http.Request.
*/
type AttachModToUserParams struct {

	/* ModID.

	   A mod UUID or slug
	*/
	ModID string

	/* ModUser.

	   The mod user data to attach
	*/
	ModUser *models.ModUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attach mod to user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachModToUserParams) WithDefaults() *AttachModToUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attach mod to user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachModToUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attach mod to user params
func (o *AttachModToUserParams) WithTimeout(timeout time.Duration) *AttachModToUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach mod to user params
func (o *AttachModToUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach mod to user params
func (o *AttachModToUserParams) WithContext(ctx context.Context) *AttachModToUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach mod to user params
func (o *AttachModToUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach mod to user params
func (o *AttachModToUserParams) WithHTTPClient(client *http.Client) *AttachModToUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach mod to user params
func (o *AttachModToUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the attach mod to user params
func (o *AttachModToUserParams) WithModID(modID string) *AttachModToUserParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the attach mod to user params
func (o *AttachModToUserParams) SetModID(modID string) {
	o.ModID = modID
}

// WithModUser adds the modUser to the attach mod to user params
func (o *AttachModToUserParams) WithModUser(modUser *models.ModUserParams) *AttachModToUserParams {
	o.SetModUser(modUser)
	return o
}

// SetModUser adds the modUser to the attach mod to user params
func (o *AttachModToUserParams) SetModUser(modUser *models.ModUserParams) {
	o.ModUser = modUser
}

// WriteToRequest writes these params to a swagger request
func (o *AttachModToUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}
	if o.ModUser != nil {
		if err := r.SetBodyParam(o.ModUser); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
