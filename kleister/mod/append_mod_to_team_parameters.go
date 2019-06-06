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

// NewAppendModToTeamParams creates a new AppendModToTeamParams object
// with the default values initialized.
func NewAppendModToTeamParams() *AppendModToTeamParams {
	var ()
	return &AppendModToTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendModToTeamParamsWithTimeout creates a new AppendModToTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendModToTeamParamsWithTimeout(timeout time.Duration) *AppendModToTeamParams {
	var ()
	return &AppendModToTeamParams{

		timeout: timeout,
	}
}

// NewAppendModToTeamParamsWithContext creates a new AppendModToTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendModToTeamParamsWithContext(ctx context.Context) *AppendModToTeamParams {
	var ()
	return &AppendModToTeamParams{

		Context: ctx,
	}
}

// NewAppendModToTeamParamsWithHTTPClient creates a new AppendModToTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendModToTeamParamsWithHTTPClient(client *http.Client) *AppendModToTeamParams {
	var ()
	return &AppendModToTeamParams{
		HTTPClient: client,
	}
}

/*AppendModToTeamParams contains all the parameters to send to the API endpoint
for the append mod to team operation typically these are written to a http.Request
*/
type AppendModToTeamParams struct {

	/*ModID
	  A mod UUID or slug

	*/
	ModID string
	/*Params
	  The mod team data to assign

	*/
	Params *models.ModTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the append mod to team params
func (o *AppendModToTeamParams) WithTimeout(timeout time.Duration) *AppendModToTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append mod to team params
func (o *AppendModToTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append mod to team params
func (o *AppendModToTeamParams) WithContext(ctx context.Context) *AppendModToTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append mod to team params
func (o *AppendModToTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append mod to team params
func (o *AppendModToTeamParams) WithHTTPClient(client *http.Client) *AppendModToTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append mod to team params
func (o *AppendModToTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the append mod to team params
func (o *AppendModToTeamParams) WithModID(modID string) *AppendModToTeamParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the append mod to team params
func (o *AppendModToTeamParams) SetModID(modID string) {
	o.ModID = modID
}

// WithParams adds the params to the append mod to team params
func (o *AppendModToTeamParams) WithParams(params *models.ModTeamParams) *AppendModToTeamParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the append mod to team params
func (o *AppendModToTeamParams) SetParams(params *models.ModTeamParams) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *AppendModToTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}

	if o.Params != nil {
		if err := r.SetBodyParam(o.Params); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
