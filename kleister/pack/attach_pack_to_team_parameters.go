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

// NewAttachPackToTeamParams creates a new AttachPackToTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachPackToTeamParams() *AttachPackToTeamParams {
	return &AttachPackToTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachPackToTeamParamsWithTimeout creates a new AttachPackToTeamParams object
// with the ability to set a timeout on a request.
func NewAttachPackToTeamParamsWithTimeout(timeout time.Duration) *AttachPackToTeamParams {
	return &AttachPackToTeamParams{
		timeout: timeout,
	}
}

// NewAttachPackToTeamParamsWithContext creates a new AttachPackToTeamParams object
// with the ability to set a context for a request.
func NewAttachPackToTeamParamsWithContext(ctx context.Context) *AttachPackToTeamParams {
	return &AttachPackToTeamParams{
		Context: ctx,
	}
}

// NewAttachPackToTeamParamsWithHTTPClient creates a new AttachPackToTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachPackToTeamParamsWithHTTPClient(client *http.Client) *AttachPackToTeamParams {
	return &AttachPackToTeamParams{
		HTTPClient: client,
	}
}

/*
AttachPackToTeamParams contains all the parameters to send to the API endpoint

	for the attach pack to team operation.

	Typically these are written to a http.Request.
*/
type AttachPackToTeamParams struct {

	/* PackID.

	   A pack UUID or slug
	*/
	PackID string

	/* PackTeam.

	   The pack team data to attach
	*/
	PackTeam *models.PackTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attach pack to team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachPackToTeamParams) WithDefaults() *AttachPackToTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attach pack to team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachPackToTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attach pack to team params
func (o *AttachPackToTeamParams) WithTimeout(timeout time.Duration) *AttachPackToTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach pack to team params
func (o *AttachPackToTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach pack to team params
func (o *AttachPackToTeamParams) WithContext(ctx context.Context) *AttachPackToTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach pack to team params
func (o *AttachPackToTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach pack to team params
func (o *AttachPackToTeamParams) WithHTTPClient(client *http.Client) *AttachPackToTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach pack to team params
func (o *AttachPackToTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the attach pack to team params
func (o *AttachPackToTeamParams) WithPackID(packID string) *AttachPackToTeamParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the attach pack to team params
func (o *AttachPackToTeamParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithPackTeam adds the packTeam to the attach pack to team params
func (o *AttachPackToTeamParams) WithPackTeam(packTeam *models.PackTeamParams) *AttachPackToTeamParams {
	o.SetPackTeam(packTeam)
	return o
}

// SetPackTeam adds the packTeam to the attach pack to team params
func (o *AttachPackToTeamParams) SetPackTeam(packTeam *models.PackTeamParams) {
	o.PackTeam = packTeam
}

// WriteToRequest writes these params to a swagger request
func (o *AttachPackToTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
