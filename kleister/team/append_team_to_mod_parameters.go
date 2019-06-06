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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// NewAppendTeamToModParams creates a new AppendTeamToModParams object
// with the default values initialized.
func NewAppendTeamToModParams() *AppendTeamToModParams {
	var ()
	return &AppendTeamToModParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendTeamToModParamsWithTimeout creates a new AppendTeamToModParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendTeamToModParamsWithTimeout(timeout time.Duration) *AppendTeamToModParams {
	var ()
	return &AppendTeamToModParams{

		timeout: timeout,
	}
}

// NewAppendTeamToModParamsWithContext creates a new AppendTeamToModParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendTeamToModParamsWithContext(ctx context.Context) *AppendTeamToModParams {
	var ()
	return &AppendTeamToModParams{

		Context: ctx,
	}
}

// NewAppendTeamToModParamsWithHTTPClient creates a new AppendTeamToModParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendTeamToModParamsWithHTTPClient(client *http.Client) *AppendTeamToModParams {
	var ()
	return &AppendTeamToModParams{
		HTTPClient: client,
	}
}

/*AppendTeamToModParams contains all the parameters to send to the API endpoint
for the append team to mod operation typically these are written to a http.Request
*/
type AppendTeamToModParams struct {

	/*Params
	  The team mod data to assign

	*/
	Params *models.TeamModParams
	/*TeamID
	  A team UUID or slug

	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the append team to mod params
func (o *AppendTeamToModParams) WithTimeout(timeout time.Duration) *AppendTeamToModParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append team to mod params
func (o *AppendTeamToModParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append team to mod params
func (o *AppendTeamToModParams) WithContext(ctx context.Context) *AppendTeamToModParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append team to mod params
func (o *AppendTeamToModParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append team to mod params
func (o *AppendTeamToModParams) WithHTTPClient(client *http.Client) *AppendTeamToModParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append team to mod params
func (o *AppendTeamToModParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParams adds the params to the append team to mod params
func (o *AppendTeamToModParams) WithParams(params *models.TeamModParams) *AppendTeamToModParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the append team to mod params
func (o *AppendTeamToModParams) SetParams(params *models.TeamModParams) {
	o.Params = params
}

// WithTeamID adds the teamID to the append team to mod params
func (o *AppendTeamToModParams) WithTeamID(teamID string) *AppendTeamToModParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the append team to mod params
func (o *AppendTeamToModParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *AppendTeamToModParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Params != nil {
		if err := r.SetBodyParam(o.Params); err != nil {
			return err
		}
	}

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
