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

// NewDeleteModFromUserParams creates a new DeleteModFromUserParams object
// with the default values initialized.
func NewDeleteModFromUserParams() *DeleteModFromUserParams {
	var ()
	return &DeleteModFromUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModFromUserParamsWithTimeout creates a new DeleteModFromUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteModFromUserParamsWithTimeout(timeout time.Duration) *DeleteModFromUserParams {
	var ()
	return &DeleteModFromUserParams{

		timeout: timeout,
	}
}

// NewDeleteModFromUserParamsWithContext creates a new DeleteModFromUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteModFromUserParamsWithContext(ctx context.Context) *DeleteModFromUserParams {
	var ()
	return &DeleteModFromUserParams{

		Context: ctx,
	}
}

// NewDeleteModFromUserParamsWithHTTPClient creates a new DeleteModFromUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteModFromUserParamsWithHTTPClient(client *http.Client) *DeleteModFromUserParams {
	var ()
	return &DeleteModFromUserParams{
		HTTPClient: client,
	}
}

/*DeleteModFromUserParams contains all the parameters to send to the API endpoint
for the delete mod from user operation typically these are written to a http.Request
*/
type DeleteModFromUserParams struct {

	/*ModID
	  A mod UUID or slug

	*/
	ModID string
	/*Params
	  The mod user data to delete

	*/
	Params *models.ModUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
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

// WithParams adds the params to the delete mod from user params
func (o *DeleteModFromUserParams) WithParams(params *models.ModUserParams) *DeleteModFromUserParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the delete mod from user params
func (o *DeleteModFromUserParams) SetParams(params *models.ModUserParams) {
	o.Params = params
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
