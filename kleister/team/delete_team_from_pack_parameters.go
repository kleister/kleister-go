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

// NewDeleteTeamFromPackParams creates a new DeleteTeamFromPackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTeamFromPackParams() *DeleteTeamFromPackParams {
	return &DeleteTeamFromPackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamFromPackParamsWithTimeout creates a new DeleteTeamFromPackParams object
// with the ability to set a timeout on a request.
func NewDeleteTeamFromPackParamsWithTimeout(timeout time.Duration) *DeleteTeamFromPackParams {
	return &DeleteTeamFromPackParams{
		timeout: timeout,
	}
}

// NewDeleteTeamFromPackParamsWithContext creates a new DeleteTeamFromPackParams object
// with the ability to set a context for a request.
func NewDeleteTeamFromPackParamsWithContext(ctx context.Context) *DeleteTeamFromPackParams {
	return &DeleteTeamFromPackParams{
		Context: ctx,
	}
}

// NewDeleteTeamFromPackParamsWithHTTPClient creates a new DeleteTeamFromPackParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTeamFromPackParamsWithHTTPClient(client *http.Client) *DeleteTeamFromPackParams {
	return &DeleteTeamFromPackParams{
		HTTPClient: client,
	}
}

/*
DeleteTeamFromPackParams contains all the parameters to send to the API endpoint

	for the delete team from pack operation.

	Typically these are written to a http.Request.
*/
type DeleteTeamFromPackParams struct {

	/* TeamID.

	   A team UUID or slug
	*/
	TeamID string

	/* TeamPack.

	   The team pack data to delete
	*/
	TeamPack *models.TeamPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete team from pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamFromPackParams) WithDefaults() *DeleteTeamFromPackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete team from pack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamFromPackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete team from pack params
func (o *DeleteTeamFromPackParams) WithTimeout(timeout time.Duration) *DeleteTeamFromPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team from pack params
func (o *DeleteTeamFromPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team from pack params
func (o *DeleteTeamFromPackParams) WithContext(ctx context.Context) *DeleteTeamFromPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team from pack params
func (o *DeleteTeamFromPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team from pack params
func (o *DeleteTeamFromPackParams) WithHTTPClient(client *http.Client) *DeleteTeamFromPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team from pack params
func (o *DeleteTeamFromPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the delete team from pack params
func (o *DeleteTeamFromPackParams) WithTeamID(teamID string) *DeleteTeamFromPackParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete team from pack params
func (o *DeleteTeamFromPackParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamPack adds the teamPack to the delete team from pack params
func (o *DeleteTeamFromPackParams) WithTeamPack(teamPack *models.TeamPackParams) *DeleteTeamFromPackParams {
	o.SetTeamPack(teamPack)
	return o
}

// SetTeamPack adds the teamPack to the delete team from pack params
func (o *DeleteTeamFromPackParams) SetTeamPack(teamPack *models.TeamPackParams) {
	o.TeamPack = teamPack
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamFromPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
