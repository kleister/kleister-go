// Code generated by go-swagger; DO NOT EDIT.

package mod

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

// NewDeleteModFromTeamParams creates a new DeleteModFromTeamParams object
// with the default values initialized.
func NewDeleteModFromTeamParams() *DeleteModFromTeamParams {
	var ()
	return &DeleteModFromTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModFromTeamParamsWithTimeout creates a new DeleteModFromTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteModFromTeamParamsWithTimeout(timeout time.Duration) *DeleteModFromTeamParams {
	var ()
	return &DeleteModFromTeamParams{

		timeout: timeout,
	}
}

// NewDeleteModFromTeamParamsWithContext creates a new DeleteModFromTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteModFromTeamParamsWithContext(ctx context.Context) *DeleteModFromTeamParams {
	var ()
	return &DeleteModFromTeamParams{

		Context: ctx,
	}
}

// NewDeleteModFromTeamParamsWithHTTPClient creates a new DeleteModFromTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteModFromTeamParamsWithHTTPClient(client *http.Client) *DeleteModFromTeamParams {
	var ()
	return &DeleteModFromTeamParams{
		HTTPClient: client,
	}
}

/*DeleteModFromTeamParams contains all the parameters to send to the API endpoint
for the delete mod from team operation typically these are written to a http.Request
*/
type DeleteModFromTeamParams struct {

	/*ModID
	  A mod UUID or slug

	*/
	ModID string
	/*ModTeam
	  The mod team data to delete

	*/
	ModTeam *models.ModTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete mod from team params
func (o *DeleteModFromTeamParams) WithTimeout(timeout time.Duration) *DeleteModFromTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mod from team params
func (o *DeleteModFromTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mod from team params
func (o *DeleteModFromTeamParams) WithContext(ctx context.Context) *DeleteModFromTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mod from team params
func (o *DeleteModFromTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mod from team params
func (o *DeleteModFromTeamParams) WithHTTPClient(client *http.Client) *DeleteModFromTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mod from team params
func (o *DeleteModFromTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the delete mod from team params
func (o *DeleteModFromTeamParams) WithModID(modID string) *DeleteModFromTeamParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the delete mod from team params
func (o *DeleteModFromTeamParams) SetModID(modID string) {
	o.ModID = modID
}

// WithModTeam adds the modTeam to the delete mod from team params
func (o *DeleteModFromTeamParams) WithModTeam(modTeam *models.ModTeamParams) *DeleteModFromTeamParams {
	o.SetModTeam(modTeam)
	return o
}

// SetModTeam adds the modTeam to the delete mod from team params
func (o *DeleteModFromTeamParams) SetModTeam(modTeam *models.ModTeamParams) {
	o.ModTeam = modTeam
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteModFromTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}

	if o.ModTeam != nil {
		if err := r.SetBodyParam(o.ModTeam); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}