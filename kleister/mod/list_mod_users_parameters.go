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
)

// NewListModUsersParams creates a new ListModUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListModUsersParams() *ListModUsersParams {
	return &ListModUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListModUsersParamsWithTimeout creates a new ListModUsersParams object
// with the ability to set a timeout on a request.
func NewListModUsersParamsWithTimeout(timeout time.Duration) *ListModUsersParams {
	return &ListModUsersParams{
		timeout: timeout,
	}
}

// NewListModUsersParamsWithContext creates a new ListModUsersParams object
// with the ability to set a context for a request.
func NewListModUsersParamsWithContext(ctx context.Context) *ListModUsersParams {
	return &ListModUsersParams{
		Context: ctx,
	}
}

// NewListModUsersParamsWithHTTPClient creates a new ListModUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListModUsersParamsWithHTTPClient(client *http.Client) *ListModUsersParams {
	return &ListModUsersParams{
		HTTPClient: client,
	}
}

/*
ListModUsersParams contains all the parameters to send to the API endpoint

	for the list mod users operation.

	Typically these are written to a http.Request.
*/
type ListModUsersParams struct {

	/* ModID.

	   A mod UUID or slug
	*/
	ModID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list mod users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListModUsersParams) WithDefaults() *ListModUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list mod users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListModUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list mod users params
func (o *ListModUsersParams) WithTimeout(timeout time.Duration) *ListModUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list mod users params
func (o *ListModUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list mod users params
func (o *ListModUsersParams) WithContext(ctx context.Context) *ListModUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list mod users params
func (o *ListModUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list mod users params
func (o *ListModUsersParams) WithHTTPClient(client *http.Client) *ListModUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list mod users params
func (o *ListModUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModID adds the modID to the list mod users params
func (o *ListModUsersParams) WithModID(modID string) *ListModUsersParams {
	o.SetModID(modID)
	return o
}

// SetModID adds the modId to the list mod users params
func (o *ListModUsersParams) SetModID(modID string) {
	o.ModID = modID
}

// WriteToRequest writes these params to a swagger request
func (o *ListModUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mod_id
	if err := r.SetPathParam("mod_id", o.ModID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
