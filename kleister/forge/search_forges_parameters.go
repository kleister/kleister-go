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
	"github.com/go-openapi/strfmt"
)

// NewSearchForgesParams creates a new SearchForgesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchForgesParams() *SearchForgesParams {
	return &SearchForgesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchForgesParamsWithTimeout creates a new SearchForgesParams object
// with the ability to set a timeout on a request.
func NewSearchForgesParamsWithTimeout(timeout time.Duration) *SearchForgesParams {
	return &SearchForgesParams{
		timeout: timeout,
	}
}

// NewSearchForgesParamsWithContext creates a new SearchForgesParams object
// with the ability to set a context for a request.
func NewSearchForgesParamsWithContext(ctx context.Context) *SearchForgesParams {
	return &SearchForgesParams{
		Context: ctx,
	}
}

// NewSearchForgesParamsWithHTTPClient creates a new SearchForgesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchForgesParamsWithHTTPClient(client *http.Client) *SearchForgesParams {
	return &SearchForgesParams{
		HTTPClient: client,
	}
}

/*
SearchForgesParams contains all the parameters to send to the API endpoint

	for the search forges operation.

	Typically these are written to a http.Request.
*/
type SearchForgesParams struct {

	/* ForgeID.

	   A search token to search Forge versions
	*/
	ForgeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search forges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchForgesParams) WithDefaults() *SearchForgesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search forges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchForgesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search forges params
func (o *SearchForgesParams) WithTimeout(timeout time.Duration) *SearchForgesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search forges params
func (o *SearchForgesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search forges params
func (o *SearchForgesParams) WithContext(ctx context.Context) *SearchForgesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search forges params
func (o *SearchForgesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search forges params
func (o *SearchForgesParams) WithHTTPClient(client *http.Client) *SearchForgesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search forges params
func (o *SearchForgesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForgeID adds the forgeID to the search forges params
func (o *SearchForgesParams) WithForgeID(forgeID string) *SearchForgesParams {
	o.SetForgeID(forgeID)
	return o
}

// SetForgeID adds the forgeId to the search forges params
func (o *SearchForgesParams) SetForgeID(forgeID string) {
	o.ForgeID = forgeID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchForgesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param forge_id
	if err := r.SetPathParam("forge_id", o.ForgeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
