// Code generated by go-swagger; DO NOT EDIT.

package minecraft

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

// NewListMinecraftsParams creates a new ListMinecraftsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMinecraftsParams() *ListMinecraftsParams {
	return &ListMinecraftsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMinecraftsParamsWithTimeout creates a new ListMinecraftsParams object
// with the ability to set a timeout on a request.
func NewListMinecraftsParamsWithTimeout(timeout time.Duration) *ListMinecraftsParams {
	return &ListMinecraftsParams{
		timeout: timeout,
	}
}

// NewListMinecraftsParamsWithContext creates a new ListMinecraftsParams object
// with the ability to set a context for a request.
func NewListMinecraftsParamsWithContext(ctx context.Context) *ListMinecraftsParams {
	return &ListMinecraftsParams{
		Context: ctx,
	}
}

// NewListMinecraftsParamsWithHTTPClient creates a new ListMinecraftsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMinecraftsParamsWithHTTPClient(client *http.Client) *ListMinecraftsParams {
	return &ListMinecraftsParams{
		HTTPClient: client,
	}
}

/*
ListMinecraftsParams contains all the parameters to send to the API endpoint

	for the list minecrafts operation.

	Typically these are written to a http.Request.
*/
type ListMinecraftsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list minecrafts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMinecraftsParams) WithDefaults() *ListMinecraftsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list minecrafts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMinecraftsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list minecrafts params
func (o *ListMinecraftsParams) WithTimeout(timeout time.Duration) *ListMinecraftsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list minecrafts params
func (o *ListMinecraftsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list minecrafts params
func (o *ListMinecraftsParams) WithContext(ctx context.Context) *ListMinecraftsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list minecrafts params
func (o *ListMinecraftsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list minecrafts params
func (o *ListMinecraftsParams) WithHTTPClient(client *http.Client) *ListMinecraftsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list minecrafts params
func (o *ListMinecraftsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListMinecraftsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
