// Code generated by go-swagger; DO NOT EDIT.

package quilt

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

// NewSearchQuiltsParams creates a new SearchQuiltsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchQuiltsParams() *SearchQuiltsParams {
	return &SearchQuiltsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchQuiltsParamsWithTimeout creates a new SearchQuiltsParams object
// with the ability to set a timeout on a request.
func NewSearchQuiltsParamsWithTimeout(timeout time.Duration) *SearchQuiltsParams {
	return &SearchQuiltsParams{
		timeout: timeout,
	}
}

// NewSearchQuiltsParamsWithContext creates a new SearchQuiltsParams object
// with the ability to set a context for a request.
func NewSearchQuiltsParamsWithContext(ctx context.Context) *SearchQuiltsParams {
	return &SearchQuiltsParams{
		Context: ctx,
	}
}

// NewSearchQuiltsParamsWithHTTPClient creates a new SearchQuiltsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchQuiltsParamsWithHTTPClient(client *http.Client) *SearchQuiltsParams {
	return &SearchQuiltsParams{
		HTTPClient: client,
	}
}

/*
SearchQuiltsParams contains all the parameters to send to the API endpoint

	for the search quilts operation.

	Typically these are written to a http.Request.
*/
type SearchQuiltsParams struct {

	/* QuiltID.

	   A search token to search quilt versions
	*/
	QuiltID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search quilts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchQuiltsParams) WithDefaults() *SearchQuiltsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search quilts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchQuiltsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search quilts params
func (o *SearchQuiltsParams) WithTimeout(timeout time.Duration) *SearchQuiltsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search quilts params
func (o *SearchQuiltsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search quilts params
func (o *SearchQuiltsParams) WithContext(ctx context.Context) *SearchQuiltsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search quilts params
func (o *SearchQuiltsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search quilts params
func (o *SearchQuiltsParams) WithHTTPClient(client *http.Client) *SearchQuiltsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search quilts params
func (o *SearchQuiltsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuiltID adds the quiltID to the search quilts params
func (o *SearchQuiltsParams) WithQuiltID(quiltID string) *SearchQuiltsParams {
	o.SetQuiltID(quiltID)
	return o
}

// SetQuiltID adds the quiltId to the search quilts params
func (o *SearchQuiltsParams) SetQuiltID(quiltID string) {
	o.QuiltID = quiltID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchQuiltsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param quilt_id
	if err := r.SetPathParam("quilt_id", o.QuiltID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
