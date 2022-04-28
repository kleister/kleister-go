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
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// NewDeletePackFromUserParams creates a new DeletePackFromUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePackFromUserParams() *DeletePackFromUserParams {
	return &DeletePackFromUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePackFromUserParamsWithTimeout creates a new DeletePackFromUserParams object
// with the ability to set a timeout on a request.
func NewDeletePackFromUserParamsWithTimeout(timeout time.Duration) *DeletePackFromUserParams {
	return &DeletePackFromUserParams{
		timeout: timeout,
	}
}

// NewDeletePackFromUserParamsWithContext creates a new DeletePackFromUserParams object
// with the ability to set a context for a request.
func NewDeletePackFromUserParamsWithContext(ctx context.Context) *DeletePackFromUserParams {
	return &DeletePackFromUserParams{
		Context: ctx,
	}
}

// NewDeletePackFromUserParamsWithHTTPClient creates a new DeletePackFromUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePackFromUserParamsWithHTTPClient(client *http.Client) *DeletePackFromUserParams {
	return &DeletePackFromUserParams{
		HTTPClient: client,
	}
}

/* DeletePackFromUserParams contains all the parameters to send to the API endpoint
   for the delete pack from user operation.

   Typically these are written to a http.Request.
*/
type DeletePackFromUserParams struct {

	/* PackID.

	   A pack UUID or slug
	*/
	PackID string

	/* PackUser.

	   The pack user data to delete
	*/
	PackUser *models.PackUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete pack from user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePackFromUserParams) WithDefaults() *DeletePackFromUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete pack from user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePackFromUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete pack from user params
func (o *DeletePackFromUserParams) WithTimeout(timeout time.Duration) *DeletePackFromUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pack from user params
func (o *DeletePackFromUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pack from user params
func (o *DeletePackFromUserParams) WithContext(ctx context.Context) *DeletePackFromUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pack from user params
func (o *DeletePackFromUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pack from user params
func (o *DeletePackFromUserParams) WithHTTPClient(client *http.Client) *DeletePackFromUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pack from user params
func (o *DeletePackFromUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackID adds the packID to the delete pack from user params
func (o *DeletePackFromUserParams) WithPackID(packID string) *DeletePackFromUserParams {
	o.SetPackID(packID)
	return o
}

// SetPackID adds the packId to the delete pack from user params
func (o *DeletePackFromUserParams) SetPackID(packID string) {
	o.PackID = packID
}

// WithPackUser adds the packUser to the delete pack from user params
func (o *DeletePackFromUserParams) WithPackUser(packUser *models.PackUserParams) *DeletePackFromUserParams {
	o.SetPackUser(packUser)
	return o
}

// SetPackUser adds the packUser to the delete pack from user params
func (o *DeletePackFromUserParams) SetPackUser(packUser *models.PackUserParams) {
	o.PackUser = packUser
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePackFromUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pack_id
	if err := r.SetPathParam("pack_id", o.PackID); err != nil {
		return err
	}
	if o.PackUser != nil {
		if err := r.SetBodyParam(o.PackUser); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
