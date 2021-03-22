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
)

// NewListTeamPacksParams creates a new ListTeamPacksParams object
// with the default values initialized.
func NewListTeamPacksParams() *ListTeamPacksParams {
	var ()
	return &ListTeamPacksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTeamPacksParamsWithTimeout creates a new ListTeamPacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTeamPacksParamsWithTimeout(timeout time.Duration) *ListTeamPacksParams {
	var ()
	return &ListTeamPacksParams{

		timeout: timeout,
	}
}

// NewListTeamPacksParamsWithContext creates a new ListTeamPacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTeamPacksParamsWithContext(ctx context.Context) *ListTeamPacksParams {
	var ()
	return &ListTeamPacksParams{

		Context: ctx,
	}
}

// NewListTeamPacksParamsWithHTTPClient creates a new ListTeamPacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTeamPacksParamsWithHTTPClient(client *http.Client) *ListTeamPacksParams {
	var ()
	return &ListTeamPacksParams{
		HTTPClient: client,
	}
}

/*ListTeamPacksParams contains all the parameters to send to the API endpoint
for the list team packs operation typically these are written to a http.Request
*/
type ListTeamPacksParams struct {

	/*TeamID
	  A team UUID or slug

	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list team packs params
func (o *ListTeamPacksParams) WithTimeout(timeout time.Duration) *ListTeamPacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list team packs params
func (o *ListTeamPacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list team packs params
func (o *ListTeamPacksParams) WithContext(ctx context.Context) *ListTeamPacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list team packs params
func (o *ListTeamPacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list team packs params
func (o *ListTeamPacksParams) WithHTTPClient(client *http.Client) *ListTeamPacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list team packs params
func (o *ListTeamPacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the list team packs params
func (o *ListTeamPacksParams) WithTeamID(teamID string) *ListTeamPacksParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the list team packs params
func (o *ListTeamPacksParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *ListTeamPacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}