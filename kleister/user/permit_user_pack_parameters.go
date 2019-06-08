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

// NewPermitUserPackParams creates a new PermitUserPackParams object
// with the default values initialized.
func NewPermitUserPackParams() *PermitUserPackParams {
	var ()
	return &PermitUserPackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPermitUserPackParamsWithTimeout creates a new PermitUserPackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPermitUserPackParamsWithTimeout(timeout time.Duration) *PermitUserPackParams {
	var ()
	return &PermitUserPackParams{

		timeout: timeout,
	}
}

// NewPermitUserPackParamsWithContext creates a new PermitUserPackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPermitUserPackParamsWithContext(ctx context.Context) *PermitUserPackParams {
	var ()
	return &PermitUserPackParams{

		Context: ctx,
	}
}

// NewPermitUserPackParamsWithHTTPClient creates a new PermitUserPackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPermitUserPackParamsWithHTTPClient(client *http.Client) *PermitUserPackParams {
	var ()
	return &PermitUserPackParams{
		HTTPClient: client,
	}
}

/*PermitUserPackParams contains all the parameters to send to the API endpoint
for the permit user pack operation typically these are written to a http.Request
*/
type PermitUserPackParams struct {

	/*UserID
	  A user UUID or slug

	*/
	UserID string
	/*UserPack
	  The user pack data to update

	*/
	UserPack *models.UserPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the permit user pack params
func (o *PermitUserPackParams) WithTimeout(timeout time.Duration) *PermitUserPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permit user pack params
func (o *PermitUserPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permit user pack params
func (o *PermitUserPackParams) WithContext(ctx context.Context) *PermitUserPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permit user pack params
func (o *PermitUserPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permit user pack params
func (o *PermitUserPackParams) WithHTTPClient(client *http.Client) *PermitUserPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permit user pack params
func (o *PermitUserPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the permit user pack params
func (o *PermitUserPackParams) WithUserID(userID string) *PermitUserPackParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the permit user pack params
func (o *PermitUserPackParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserPack adds the userPack to the permit user pack params
func (o *PermitUserPackParams) WithUserPack(userPack *models.UserPackParams) *PermitUserPackParams {
	o.SetUserPack(userPack)
	return o
}

// SetUserPack adds the userPack to the permit user pack params
func (o *PermitUserPackParams) SetUserPack(userPack *models.UserPackParams) {
	o.UserPack = userPack
}

// WriteToRequest writes these params to a swagger request
func (o *PermitUserPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if o.UserPack != nil {
		if err := r.SetBodyParam(o.UserPack); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
