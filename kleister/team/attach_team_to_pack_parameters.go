// Code generated by go-swagger; DO NOT EDIT.

package team

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

// NewAttachTeamToPackParams creates a new AttachTeamToPackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachTeamToPackParams() *AttachTeamToPackParams {
	return &AttachTeamToPackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachTeamToPackParamsWithTimeout creates a new AttachTeamToPackParams object
// with the ability to set a timeout on a request.
func NewAttachTeamToPackParamsWithTimeout(timeout time.Duration) *AttachTeamToPackParams {
	return &AttachTeamToPackParams{
		timeout: timeout,
	}
}

// NewAttachTeamToPackParamsWithContext creates a new AttachTeamToPackParams object
// with the ability to set a context for a request.
func NewAttachTeamToPackParamsWithContext(ctx context.Context) *AttachTeamToPackParams {
	return &AttachTeamToPackParams{
		Context: ctx,
	}
}

// NewAttachTeamToPackParamsWithHTTPClient creates a new AttachTeamToPackParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachTeamToPackParamsWithHTTPClient(client *http.Client) *AttachTeamToPackParams {
	return &AttachTeamToPackParams{
		HTTPClient: client,
	}
}

/*
AttachTeamToPackParams contains all the parameters to send to the API endpoint

	for the attach team to pack operation.

	Typically these are written to a http.Request.
*/
type AttachTeamToPackParams struct {

	/* TeamID.

	   A team UUID or slug
	*/
	TeamID string

	/* TeamPack.

	   The team pack data to attach
	*/
	TeamPack *models.TeamPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attach team to pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachTeamToPackParams) WithDefaults() *AttachTeamToPackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attach team to pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachTeamToPackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attach team to pack params
func (o *AttachTeamToPackParams) WithTimeout(timeout time.Duration) *AttachTeamToPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attach team to pack params
func (o *AttachTeamToPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attach team to pack params
func (o *AttachTeamToPackParams) WithContext(ctx context.Context) *AttachTeamToPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attach team to pack params
func (o *AttachTeamToPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attach team to pack params
func (o *AttachTeamToPackParams) WithHTTPClient(client *http.Client) *AttachTeamToPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attach team to pack params
func (o *AttachTeamToPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the attach team to pack params
func (o *AttachTeamToPackParams) WithTeamID(teamID string) *AttachTeamToPackParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the attach team to pack params
func (o *AttachTeamToPackParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamPack adds the teamPack to the attach team to pack params
func (o *AttachTeamToPackParams) WithTeamPack(teamPack *models.TeamPackParams) *AttachTeamToPackParams {
	o.SetTeamPack(teamPack)
	return o
}

// SetTeamPack adds the teamPack to the attach team to pack params
func (o *AttachTeamToPackParams) SetTeamPack(teamPack *models.TeamPackParams) {
	o.TeamPack = teamPack
}

// WriteToRequest writes these params to a swagger request
func (o *AttachTeamToPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}
	if o.TeamPack != nil {
		if err := r.SetBodyParam(o.TeamPack); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
