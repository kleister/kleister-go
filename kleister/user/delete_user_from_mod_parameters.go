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

// NewDeleteUserFromModParams creates a new DeleteUserFromModParams object
// with the default values initialized.
func NewDeleteUserFromModParams() *DeleteUserFromModParams {
	var ()
	return &DeleteUserFromModParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserFromModParamsWithTimeout creates a new DeleteUserFromModParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserFromModParamsWithTimeout(timeout time.Duration) *DeleteUserFromModParams {
	var ()
	return &DeleteUserFromModParams{

		timeout: timeout,
	}
}

// NewDeleteUserFromModParamsWithContext creates a new DeleteUserFromModParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserFromModParamsWithContext(ctx context.Context) *DeleteUserFromModParams {
	var ()
	return &DeleteUserFromModParams{

		Context: ctx,
	}
}

// NewDeleteUserFromModParamsWithHTTPClient creates a new DeleteUserFromModParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserFromModParamsWithHTTPClient(client *http.Client) *DeleteUserFromModParams {
	var ()
	return &DeleteUserFromModParams{
		HTTPClient: client,
	}
}

/*DeleteUserFromModParams contains all the parameters to send to the API endpoint
for the delete user from mod operation typically these are written to a http.Request
*/
type DeleteUserFromModParams struct {

	/*UserID
	  A user UUID or slug

	*/
	UserID string
	/*UserMod
	  The user mod data to delete

	*/
	UserMod *models.UserModParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user from mod params
func (o *DeleteUserFromModParams) WithTimeout(timeout time.Duration) *DeleteUserFromModParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user from mod params
func (o *DeleteUserFromModParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user from mod params
func (o *DeleteUserFromModParams) WithContext(ctx context.Context) *DeleteUserFromModParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user from mod params
func (o *DeleteUserFromModParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user from mod params
func (o *DeleteUserFromModParams) WithHTTPClient(client *http.Client) *DeleteUserFromModParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user from mod params
func (o *DeleteUserFromModParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user from mod params
func (o *DeleteUserFromModParams) WithUserID(userID string) *DeleteUserFromModParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user from mod params
func (o *DeleteUserFromModParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserMod adds the userMod to the delete user from mod params
func (o *DeleteUserFromModParams) WithUserMod(userMod *models.UserModParams) *DeleteUserFromModParams {
	o.SetUserMod(userMod)
	return o
}

// SetUserMod adds the userMod to the delete user from mod params
func (o *DeleteUserFromModParams) SetUserMod(userMod *models.UserModParams) {
	o.UserMod = userMod
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserFromModParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if o.UserMod != nil {
		if err := r.SetBodyParam(o.UserMod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
