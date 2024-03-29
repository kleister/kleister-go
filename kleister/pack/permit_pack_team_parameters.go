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

	"github.com/kleister/kleister-go/models"
)

// NewPermitPackTeamParams creates a new PermitPackTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPermitPackTeamParams() *PermitPackTeamParams {
	return &PermitPackTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPermitPackTeamParamsWithTimeout creates a new PermitPackTeamParams object
// with the ability to set a timeout on a request.
func NewPermitPackTeamParamsWithTimeout(timeout time.Duration) *PermitPackTeamParams {
	return &PermitPackTeamParams{
		timeout: timeout,
	}
}

// NewPermitPackTeamParamsWithContext creates a new PermitPackTeamParams object
// with the ability to set a context for a request.
func NewPermitPackTeamParamsWithContext(ctx context.Context) *PermitPackTeamParams {
	return &PermitPackTeamParams{
		Context: ctx,
	}
}

// NewPermitPackTeamParamsWithHTTPClient creates a new PermitPackTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewPermitPackTeamParamsWithHTTPClient(client *http.Client) *PermitPackTeamParams {
	return &PermitPackTeamParams{
		HTTPClient: client,
	}
}

/*
PermitPackTeamParams contains all the parameters to send to the API endpoint

	for the permit pack team operation.

	Typically these are written to a http.Request.
*/
type PermitPackTeamParams struct {

	/* PackID.

	   A pack UUID or slug
	*/
	PackID string

	/* PackTeam.

	   The pack team data to update
	*/
	PackTeam *models.PackTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the permit pack team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitPackTeamParams) WithDefaults() *PermitPackTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the permit pack team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitPackTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the permit pack team params
func (o *PermitPackTeamParams) WithTimeout(timeout time.Duration) *PermitPackTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permit pack team params
func (o *PermitPackTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permit pack team params
func (o *PermitPackTeamParams) WithContext(ctx context.Context) *PermitPackTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permit pack team params
func (o *PermitPackTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permit pack team params
func (o *PermitPackTeamParams) WithHTTPClient(client *http.Client) *PermitPackTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permit pack team params
func (o *PermitPackTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the permit pack team params
func (o *PermitPackTeamParams) WithPackID(packID string) *PermitPackTeamParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the permit pack team params
func (o *PermitPackTeamParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithPackTeam adds the packTeam to the permit pack team params
func (o *PermitPackTeamParams) WithPackTeam(packTeam *models.PackTeamParams) *PermitPackTeamParams {
	o.SetPackTeam(packTeam)
	return o
}

// SetPackTeam adds the packTeam to the permit pack team params
func (o *PermitPackTeamParams) SetPackTeam(packTeam *models.PackTeamParams) {
	o.PackTeam = packTeam
}

// WriteToRequest writes these params to a swagger request
func (o *PermitPackTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pack_id
	if err := r.SetPathParam("pack_id", o.PackID); err != nil {
		return err
	}
	if o.PackTeam != nil {
		if err := r.SetBodyParam(o.PackTeam); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
