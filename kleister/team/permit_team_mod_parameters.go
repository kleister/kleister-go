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

	"github.com/kleister/kleister-go/v1/models"
)

// NewPermitTeamModParams creates a new PermitTeamModParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPermitTeamModParams() *PermitTeamModParams {
	return &PermitTeamModParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPermitTeamModParamsWithTimeout creates a new PermitTeamModParams object
// with the ability to set a timeout on a request.
func NewPermitTeamModParamsWithTimeout(timeout time.Duration) *PermitTeamModParams {
	return &PermitTeamModParams{
		timeout: timeout,
	}
}

// NewPermitTeamModParamsWithContext creates a new PermitTeamModParams object
// with the ability to set a context for a request.
func NewPermitTeamModParamsWithContext(ctx context.Context) *PermitTeamModParams {
	return &PermitTeamModParams{
		Context: ctx,
	}
}

// NewPermitTeamModParamsWithHTTPClient creates a new PermitTeamModParams object
// with the ability to set a custom HTTPClient for a request.
func NewPermitTeamModParamsWithHTTPClient(client *http.Client) *PermitTeamModParams {
	return &PermitTeamModParams{
		HTTPClient: client,
	}
}

/* PermitTeamModParams contains all the parameters to send to the API endpoint
   for the permit team mod operation.

   Typically these are written to a http.Request.
*/
type PermitTeamModParams struct {

	/* TeamID.

	   A team UUID or slug
	*/
	TeamID string

	/* TeamMod.

	   The team mod data to update
	*/
	TeamMod *models.TeamModParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the permit team mod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitTeamModParams) WithDefaults() *PermitTeamModParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the permit team mod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitTeamModParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the permit team mod params
func (o *PermitTeamModParams) WithTimeout(timeout time.Duration) *PermitTeamModParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permit team mod params
func (o *PermitTeamModParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permit team mod params
func (o *PermitTeamModParams) WithContext(ctx context.Context) *PermitTeamModParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permit team mod params
func (o *PermitTeamModParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permit team mod params
func (o *PermitTeamModParams) WithHTTPClient(client *http.Client) *PermitTeamModParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permit team mod params
func (o *PermitTeamModParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the permit team mod params
func (o *PermitTeamModParams) WithTeamID(teamID string) *PermitTeamModParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the permit team mod params
func (o *PermitTeamModParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamMod adds the teamMod to the permit team mod params
func (o *PermitTeamModParams) WithTeamMod(teamMod *models.TeamModParams) *PermitTeamModParams {
	o.SetTeamMod(teamMod)
	return o
}

// SetTeamMod adds the teamMod to the permit team mod params
func (o *PermitTeamModParams) SetTeamMod(teamMod *models.TeamModParams) {
	o.TeamMod = teamMod
}

// WriteToRequest writes these params to a swagger request
func (o *PermitTeamModParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}
	if o.TeamMod != nil {
		if err := r.SetBodyParam(o.TeamMod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
