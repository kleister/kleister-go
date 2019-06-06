// Code generated by go-swagger; DO NOT EDIT.

package forge

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

// NewAppendForgeToBuildParams creates a new AppendForgeToBuildParams object
// with the default values initialized.
func NewAppendForgeToBuildParams() *AppendForgeToBuildParams {
	var ()
	return &AppendForgeToBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendForgeToBuildParamsWithTimeout creates a new AppendForgeToBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendForgeToBuildParamsWithTimeout(timeout time.Duration) *AppendForgeToBuildParams {
	var ()
	return &AppendForgeToBuildParams{

		timeout: timeout,
	}
}

// NewAppendForgeToBuildParamsWithContext creates a new AppendForgeToBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendForgeToBuildParamsWithContext(ctx context.Context) *AppendForgeToBuildParams {
	var ()
	return &AppendForgeToBuildParams{

		Context: ctx,
	}
}

// NewAppendForgeToBuildParamsWithHTTPClient creates a new AppendForgeToBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendForgeToBuildParamsWithHTTPClient(client *http.Client) *AppendForgeToBuildParams {
	var ()
	return &AppendForgeToBuildParams{
		HTTPClient: client,
	}
}

/*AppendForgeToBuildParams contains all the parameters to send to the API endpoint
for the append forge to build operation typically these are written to a http.Request
*/
type AppendForgeToBuildParams struct {

	/*ForgeID
	  A forge UUID or slug

	*/
	ForgeID string
	/*Params
	  The build data to append

	*/
	Params *models.ForgeBuildParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the append forge to build params
func (o *AppendForgeToBuildParams) WithTimeout(timeout time.Duration) *AppendForgeToBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append forge to build params
func (o *AppendForgeToBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append forge to build params
func (o *AppendForgeToBuildParams) WithContext(ctx context.Context) *AppendForgeToBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append forge to build params
func (o *AppendForgeToBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append forge to build params
func (o *AppendForgeToBuildParams) WithHTTPClient(client *http.Client) *AppendForgeToBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append forge to build params
func (o *AppendForgeToBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForgeID adds the forgeID to the append forge to build params
func (o *AppendForgeToBuildParams) WithForgeID(forgeID string) *AppendForgeToBuildParams {
	o.SetForgeID(forgeID)
	return o
}

// SetForgeID adds the forgeId to the append forge to build params
func (o *AppendForgeToBuildParams) SetForgeID(forgeID string) {
	o.ForgeID = forgeID
}

// WithParams adds the params to the append forge to build params
func (o *AppendForgeToBuildParams) WithParams(params *models.ForgeBuildParams) *AppendForgeToBuildParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the append forge to build params
func (o *AppendForgeToBuildParams) SetParams(params *models.ForgeBuildParams) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *AppendForgeToBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param forge_id
	if err := r.SetPathParam("forge_id", o.ForgeID); err != nil {
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
