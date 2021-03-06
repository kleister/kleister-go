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

// NewAppendTeamToPackParams creates a new AppendTeamToPackParams object
// with the default values initialized.
func NewAppendTeamToPackParams() *AppendTeamToPackParams {
	var ()
	return &AppendTeamToPackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendTeamToPackParamsWithTimeout creates a new AppendTeamToPackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendTeamToPackParamsWithTimeout(timeout time.Duration) *AppendTeamToPackParams {
	var ()
	return &AppendTeamToPackParams{

		timeout: timeout,
	}
}

// NewAppendTeamToPackParamsWithContext creates a new AppendTeamToPackParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendTeamToPackParamsWithContext(ctx context.Context) *AppendTeamToPackParams {
	var ()
	return &AppendTeamToPackParams{

		Context: ctx,
	}
}

// NewAppendTeamToPackParamsWithHTTPClient creates a new AppendTeamToPackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendTeamToPackParamsWithHTTPClient(client *http.Client) *AppendTeamToPackParams {
	var ()
	return &AppendTeamToPackParams{
		HTTPClient: client,
	}
}

/*AppendTeamToPackParams contains all the parameters to send to the API endpoint
for the append team to pack operation typically these are written to a http.Request
*/
type AppendTeamToPackParams struct {

	/*TeamID
	  A team UUID or slug

	*/
	TeamID string
	/*TeamPack
	  The team pack data to assign

	*/
	TeamPack *models.TeamPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the append team to pack params
func (o *AppendTeamToPackParams) WithTimeout(timeout time.Duration) *AppendTeamToPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append team to pack params
func (o *AppendTeamToPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append team to pack params
func (o *AppendTeamToPackParams) WithContext(ctx context.Context) *AppendTeamToPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append team to pack params
func (o *AppendTeamToPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append team to pack params
func (o *AppendTeamToPackParams) WithHTTPClient(client *http.Client) *AppendTeamToPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append team to pack params
func (o *AppendTeamToPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the append team to pack params
func (o *AppendTeamToPackParams) WithTeamID(teamID string) *AppendTeamToPackParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the append team to pack params
func (o *AppendTeamToPackParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamPack adds the teamPack to the append team to pack params
func (o *AppendTeamToPackParams) WithTeamPack(teamPack *models.TeamPackParams) *AppendTeamToPackParams {
	o.SetTeamPack(teamPack)
	return o
}

// SetTeamPack adds the teamPack to the append team to pack params
func (o *AppendTeamToPackParams) SetTeamPack(teamPack *models.TeamPackParams) {
	o.TeamPack = teamPack
}

// WriteToRequest writes these params to a swagger request
func (o *AppendTeamToPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
