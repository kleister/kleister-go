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

// NewPermitTeamPackParams creates a new PermitTeamPackParams object
// with the default values initialized.
func NewPermitTeamPackParams() *PermitTeamPackParams {
	var ()
	return &PermitTeamPackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPermitTeamPackParamsWithTimeout creates a new PermitTeamPackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPermitTeamPackParamsWithTimeout(timeout time.Duration) *PermitTeamPackParams {
	var ()
	return &PermitTeamPackParams{

		timeout: timeout,
	}
}

// NewPermitTeamPackParamsWithContext creates a new PermitTeamPackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPermitTeamPackParamsWithContext(ctx context.Context) *PermitTeamPackParams {
	var ()
	return &PermitTeamPackParams{

		Context: ctx,
	}
}

// NewPermitTeamPackParamsWithHTTPClient creates a new PermitTeamPackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPermitTeamPackParamsWithHTTPClient(client *http.Client) *PermitTeamPackParams {
	var ()
	return &PermitTeamPackParams{
		HTTPClient: client,
	}
}

/*PermitTeamPackParams contains all the parameters to send to the API endpoint
for the permit team pack operation typically these are written to a http.Request
*/
type PermitTeamPackParams struct {

	/*TeamID
	  A team UUID or slug

	*/
	TeamID string
	/*TeamPack
	  The team pack data to update

	*/
	TeamPack *models.TeamPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the permit team pack params
func (o *PermitTeamPackParams) WithTimeout(timeout time.Duration) *PermitTeamPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permit team pack params
func (o *PermitTeamPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permit team pack params
func (o *PermitTeamPackParams) WithContext(ctx context.Context) *PermitTeamPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permit team pack params
func (o *PermitTeamPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permit team pack params
func (o *PermitTeamPackParams) WithHTTPClient(client *http.Client) *PermitTeamPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permit team pack params
func (o *PermitTeamPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the permit team pack params
func (o *PermitTeamPackParams) WithTeamID(teamID string) *PermitTeamPackParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the permit team pack params
func (o *PermitTeamPackParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamPack adds the teamPack to the permit team pack params
func (o *PermitTeamPackParams) WithTeamPack(teamPack *models.TeamPackParams) *PermitTeamPackParams {
	o.SetTeamPack(teamPack)
	return o
}

// SetTeamPack adds the teamPack to the permit team pack params
func (o *PermitTeamPackParams) SetTeamPack(teamPack *models.TeamPackParams) {
	o.TeamPack = teamPack
}

// WriteToRequest writes these params to a swagger request
func (o *PermitTeamPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
