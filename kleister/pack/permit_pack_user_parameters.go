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
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// NewPermitPackUserParams creates a new PermitPackUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPermitPackUserParams() *PermitPackUserParams {
	return &PermitPackUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPermitPackUserParamsWithTimeout creates a new PermitPackUserParams object
// with the ability to set a timeout on a request.
func NewPermitPackUserParamsWithTimeout(timeout time.Duration) *PermitPackUserParams {
	return &PermitPackUserParams{
		timeout: timeout,
	}
}

// NewPermitPackUserParamsWithContext creates a new PermitPackUserParams object
// with the ability to set a context for a request.
func NewPermitPackUserParamsWithContext(ctx context.Context) *PermitPackUserParams {
	return &PermitPackUserParams{
		Context: ctx,
	}
}

// NewPermitPackUserParamsWithHTTPClient creates a new PermitPackUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewPermitPackUserParamsWithHTTPClient(client *http.Client) *PermitPackUserParams {
	return &PermitPackUserParams{
		HTTPClient: client,
	}
}

/*
PermitPackUserParams contains all the parameters to send to the API endpoint

	for the permit pack user operation.

	Typically these are written to a http.Request.
*/
type PermitPackUserParams struct {

	/* PackID.

	   A pack UUID or slug
	*/
	PackID string

	/* PackUser.

	   The pack user data to update
	*/
	PackUser *models.PackUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the permit pack user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitPackUserParams) WithDefaults() *PermitPackUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the permit pack user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitPackUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the permit pack user params
func (o *PermitPackUserParams) WithTimeout(timeout time.Duration) *PermitPackUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permit pack user params
func (o *PermitPackUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permit pack user params
func (o *PermitPackUserParams) WithContext(ctx context.Context) *PermitPackUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permit pack user params
func (o *PermitPackUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permit pack user params
func (o *PermitPackUserParams) WithHTTPClient(client *http.Client) *PermitPackUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permit pack user params
func (o *PermitPackUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the permit pack user params
func (o *PermitPackUserParams) WithPackID(packID string) *PermitPackUserParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the permit pack user params
func (o *PermitPackUserParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithPackUser adds the packUser to the permit pack user params
func (o *PermitPackUserParams) WithPackUser(packUser *models.PackUserParams) *PermitPackUserParams {
	o.SetPackUser(packUser)
	return o
}

// SetPackUser adds the packUser to the permit pack user params
func (o *PermitPackUserParams) SetPackUser(packUser *models.PackUserParams) {
	o.PackUser = packUser
}

// WriteToRequest writes these params to a swagger request
func (o *PermitPackUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
