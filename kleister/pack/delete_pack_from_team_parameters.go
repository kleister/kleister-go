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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// NewDeletePackFromTeamParams creates a new DeletePackFromTeamParams object
// with the default values initialized.
func NewDeletePackFromTeamParams() *DeletePackFromTeamParams {
	var ()
	return &DeletePackFromTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePackFromTeamParamsWithTimeout creates a new DeletePackFromTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePackFromTeamParamsWithTimeout(timeout time.Duration) *DeletePackFromTeamParams {
	var ()
	return &DeletePackFromTeamParams{

		timeout: timeout,
	}
}

// NewDeletePackFromTeamParamsWithContext creates a new DeletePackFromTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePackFromTeamParamsWithContext(ctx context.Context) *DeletePackFromTeamParams {
	var ()
	return &DeletePackFromTeamParams{

		Context: ctx,
	}
}

// NewDeletePackFromTeamParamsWithHTTPClient creates a new DeletePackFromTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePackFromTeamParamsWithHTTPClient(client *http.Client) *DeletePackFromTeamParams {
	var ()
	return &DeletePackFromTeamParams{
		HTTPClient: client,
	}
}

/*DeletePackFromTeamParams contains all the parameters to send to the API endpoint
for the delete pack from team operation typically these are written to a http.Request
*/
type DeletePackFromTeamParams struct {

	/*PackID
	  A pack UUID or slug

	*/
	PackID string
	/*PackTeam
	  The pack team data to delete

	*/
	PackTeam *models.PackTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete pack from team params
func (o *DeletePackFromTeamParams) WithTimeout(timeout time.Duration) *DeletePackFromTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pack from team params
func (o *DeletePackFromTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pack from team params
func (o *DeletePackFromTeamParams) WithContext(ctx context.Context) *DeletePackFromTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pack from team params
func (o *DeletePackFromTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pack from team params
func (o *DeletePackFromTeamParams) WithHTTPClient(client *http.Client) *DeletePackFromTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pack from team params
func (o *DeletePackFromTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the delete pack from team params
func (o *DeletePackFromTeamParams) WithPackID(packID string) *DeletePackFromTeamParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the delete pack from team params
func (o *DeletePackFromTeamParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithPackTeam adds the packTeam to the delete pack from team params
func (o *DeletePackFromTeamParams) WithPackTeam(packTeam *models.PackTeamParams) *DeletePackFromTeamParams {
	o.SetPackTeam(packTeam)
	return o
}

// SetPackTeam adds the packTeam to the delete pack from team params
func (o *DeletePackFromTeamParams) SetPackTeam(packTeam *models.PackTeamParams) {
	o.PackTeam = packTeam
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePackFromTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
