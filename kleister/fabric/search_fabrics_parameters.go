// Code generated by go-swagger; DO NOT EDIT.

package fabric

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

// NewSearchFabricsParams creates a new SearchFabricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchFabricsParams() *SearchFabricsParams {
	return &SearchFabricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchFabricsParamsWithTimeout creates a new SearchFabricsParams object
// with the ability to set a timeout on a request.
func NewSearchFabricsParamsWithTimeout(timeout time.Duration) *SearchFabricsParams {
	return &SearchFabricsParams{
		timeout: timeout,
	}
}

// NewSearchFabricsParamsWithContext creates a new SearchFabricsParams object
// with the ability to set a context for a request.
func NewSearchFabricsParamsWithContext(ctx context.Context) *SearchFabricsParams {
	return &SearchFabricsParams{
		Context: ctx,
	}
}

// NewSearchFabricsParamsWithHTTPClient creates a new SearchFabricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchFabricsParamsWithHTTPClient(client *http.Client) *SearchFabricsParams {
	return &SearchFabricsParams{
		HTTPClient: client,
	}
}

/*
SearchFabricsParams contains all the parameters to send to the API endpoint

	for the search fabrics operation.

	Typically these are written to a http.Request.
*/
type SearchFabricsParams struct {

	/* FabricID.

	   A search token to search fabric versions
	*/
	FabricID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search fabrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchFabricsParams) WithDefaults() *SearchFabricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search fabrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchFabricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search fabrics params
func (o *SearchFabricsParams) WithTimeout(timeout time.Duration) *SearchFabricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search fabrics params
func (o *SearchFabricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search fabrics params
func (o *SearchFabricsParams) WithContext(ctx context.Context) *SearchFabricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search fabrics params
func (o *SearchFabricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search fabrics params
func (o *SearchFabricsParams) WithHTTPClient(client *http.Client) *SearchFabricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search fabrics params
func (o *SearchFabricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFabricID adds the fabricID to the search fabrics params
func (o *SearchFabricsParams) WithFabricID(fabricID string) *SearchFabricsParams {
	o.SetFabricID(fabricID)
	return o
}

// SetFabricID adds the fabricId to the search fabrics params
func (o *SearchFabricsParams) SetFabricID(fabricID string) {
	o.FabricID = fabricID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchFabricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fabric_id
	if err := r.SetPathParam("fabric_id", o.FabricID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}