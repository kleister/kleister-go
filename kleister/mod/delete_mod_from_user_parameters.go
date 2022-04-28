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
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// NewDeleteModFromUserParams creates a new DeleteModFromUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteModFromUserParams() *DeleteModFromUserParams {
	return &DeleteModFromUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModFromUserParamsWithTimeout creates a new DeleteModFromUserParams object
// with the ability to set a timeout on a request.
func NewDeleteModFromUserParamsWithTimeout(timeout time.Duration) *DeleteModFromUserParams {
	return &DeleteModFromUserParams{
		timeout: timeout,
	}
}

// NewDeleteModFromUserParamsWithContext creates a new DeleteModFromUserParams object
// with the ability to set a context for a request.
func NewDeleteModFromUserParamsWithContext(ctx context.Context) *DeleteModFromUserParams {
	return &DeleteModFromUserParams{
		Context: ctx,
	}
}

// NewDeleteModFromUserParamsWithHTTPClient creates a new DeleteModFromUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteModFromUserParamsWithHTTPClient(client *http.Client) *DeleteModFromUserParams {
	return &DeleteModFromUserParams{
		HTTPClient: client,
	}
}

/* DeleteModFromUserParams contains all the parameters to send to the API endpoint
   for the delete mod from user operation.

   Typically these are written to a http.Request.
*/
type DeleteModFromUserParams struct {

	/* ModID.

	   A mod UUID or slug
	*/
	ModID string

	/* ModUser.

	   The mod user data to delete
	*/
	ModUser *models.ModUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete mod from user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteModFromUserParams) WithDefaults() *DeleteModFromUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete mod from user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteModFromUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete mod from user params
func (o *DeleteModFromUserParams) WithTimeout(timeout time.Duration) *DeleteModFromUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mod from user params
func (o *DeleteModFromUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mod from user params
func (o *DeleteModFromUserParams) WithContext(ctx context.Context) *DeleteModFromUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mod from user params
func (o *DeleteModFromUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mod from user params
func (o *DeleteModFromUserParams) WithHTTPClient(client *http.Client) *DeleteModFromUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mod from user params
func (o *DeleteModFromUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the delete mod from user params
func (o *DeleteModFromUserParams) WithModID(modID string) *DeleteModFromUserParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the delete mod from user params
func (o *DeleteModFromUserParams) SetModID(modID string) {
	o.ModID = modID
}

// WithModUser adds the modUser to the delete mod from user params
func (o *DeleteModFromUserParams) WithModUser(modUser *models.ModUserParams) *DeleteModFromUserParams {
	o.SetModUser(modUser)
	return o
}

// SetModUser adds the modUser to the delete mod from user params
func (o *DeleteModFromUserParams) SetModUser(modUser *models.ModUserParams) {
	o.ModUser = modUser
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteModFromUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}
	if o.ModUser != nil {
		if err := r.SetBodyParam(o.ModUser); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
