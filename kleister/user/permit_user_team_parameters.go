// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewPermitUserTeamParams creates a new PermitUserTeamParams object
// with the default values initialized.
func NewPermitUserTeamParams() *PermitUserTeamParams {
	var ()
	return &PermitUserTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPermitUserTeamParamsWithTimeout creates a new PermitUserTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPermitUserTeamParamsWithTimeout(timeout time.Duration) *PermitUserTeamParams {
	var ()
	return &PermitUserTeamParams{

		timeout: timeout,
	}
}

// NewPermitUserTeamParamsWithContext creates a new PermitUserTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewPermitUserTeamParamsWithContext(ctx context.Context) *PermitUserTeamParams {
	var ()
	return &PermitUserTeamParams{

		Context: ctx,
	}
}

// NewPermitUserTeamParamsWithHTTPClient creates a new PermitUserTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPermitUserTeamParamsWithHTTPClient(client *http.Client) *PermitUserTeamParams {
	var ()
	return &PermitUserTeamParams{
		HTTPClient: client,
	}
}

/*PermitUserTeamParams contains all the parameters to send to the API endpoint
for the permit user team operation typically these are written to a http.Request
*/
type PermitUserTeamParams struct {

	/*UserID
	  A user UUID or slug

	*/
	UserID string
	/*UserTeam
	  The user team data to update

	*/
	UserTeam *models.UserTeamParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the permit user team params
func (o *PermitUserTeamParams) WithTimeout(timeout time.Duration) *PermitUserTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permit user team params
func (o *PermitUserTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permit user team params
func (o *PermitUserTeamParams) WithContext(ctx context.Context) *PermitUserTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permit user team params
func (o *PermitUserTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permit user team params
func (o *PermitUserTeamParams) WithHTTPClient(client *http.Client) *PermitUserTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permit user team params
func (o *PermitUserTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the permit user team params
func (o *PermitUserTeamParams) WithUserID(userID string) *PermitUserTeamParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the permit user team params
func (o *PermitUserTeamParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserTeam adds the userTeam to the permit user team params
func (o *PermitUserTeamParams) WithUserTeam(userTeam *models.UserTeamParams) *PermitUserTeamParams {
	o.SetUserTeam(userTeam)
	return o
}

// SetUserTeam adds the userTeam to the permit user team params
func (o *PermitUserTeamParams) SetUserTeam(userTeam *models.UserTeamParams) {
	o.UserTeam = userTeam
}

// WriteToRequest writes these params to a swagger request
func (o *PermitUserTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if o.UserTeam != nil {
		if err := r.SetBodyParam(o.UserTeam); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
