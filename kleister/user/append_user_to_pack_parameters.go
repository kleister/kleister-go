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

// NewAppendUserToPackParams creates a new AppendUserToPackParams object
// with the default values initialized.
func NewAppendUserToPackParams() *AppendUserToPackParams {
	var ()
	return &AppendUserToPackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendUserToPackParamsWithTimeout creates a new AppendUserToPackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendUserToPackParamsWithTimeout(timeout time.Duration) *AppendUserToPackParams {
	var ()
	return &AppendUserToPackParams{

		timeout: timeout,
	}
}

// NewAppendUserToPackParamsWithContext creates a new AppendUserToPackParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendUserToPackParamsWithContext(ctx context.Context) *AppendUserToPackParams {
	var ()
	return &AppendUserToPackParams{

		Context: ctx,
	}
}

// NewAppendUserToPackParamsWithHTTPClient creates a new AppendUserToPackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendUserToPackParamsWithHTTPClient(client *http.Client) *AppendUserToPackParams {
	var ()
	return &AppendUserToPackParams{
		HTTPClient: client,
	}
}

/*AppendUserToPackParams contains all the parameters to send to the API endpoint
for the append user to pack operation typically these are written to a http.Request
*/
type AppendUserToPackParams struct {

	/*UserID
	  A user UUID or slug

	*/
	UserID string
	/*UserPack
	  The user pack data to assign

	*/
	UserPack *models.UserPackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the append user to pack params
func (o *AppendUserToPackParams) WithTimeout(timeout time.Duration) *AppendUserToPackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append user to pack params
func (o *AppendUserToPackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append user to pack params
func (o *AppendUserToPackParams) WithContext(ctx context.Context) *AppendUserToPackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append user to pack params
func (o *AppendUserToPackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append user to pack params
func (o *AppendUserToPackParams) WithHTTPClient(client *http.Client) *AppendUserToPackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append user to pack params
func (o *AppendUserToPackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the append user to pack params
func (o *AppendUserToPackParams) WithUserID(userID string) *AppendUserToPackParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the append user to pack params
func (o *AppendUserToPackParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithUserPack adds the userPack to the append user to pack params
func (o *AppendUserToPackParams) WithUserPack(userPack *models.UserPackParams) *AppendUserToPackParams {
	o.SetUserPack(userPack)
	return o
}

// SetUserPack adds the userPack to the append user to pack params
func (o *AppendUserToPackParams) SetUserPack(userPack *models.UserPackParams) {
	o.UserPack = userPack
}

// WriteToRequest writes these params to a swagger request
func (o *AppendUserToPackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
