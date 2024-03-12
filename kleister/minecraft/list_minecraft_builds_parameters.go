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

// NewListMinecraftBuildsParams creates a new ListMinecraftBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMinecraftBuildsParams() *ListMinecraftBuildsParams {
	return &ListMinecraftBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMinecraftBuildsParamsWithTimeout creates a new ListMinecraftBuildsParams object
// with the ability to set a timeout on a request.
func NewListMinecraftBuildsParamsWithTimeout(timeout time.Duration) *ListMinecraftBuildsParams {
	return &ListMinecraftBuildsParams{
		timeout: timeout,
	}
}

// NewListMinecraftBuildsParamsWithContext creates a new ListMinecraftBuildsParams object
// with the ability to set a context for a request.
func NewListMinecraftBuildsParamsWithContext(ctx context.Context) *ListMinecraftBuildsParams {
	return &ListMinecraftBuildsParams{
		Context: ctx,
	}
}

// NewListMinecraftBuildsParamsWithHTTPClient creates a new ListMinecraftBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMinecraftBuildsParamsWithHTTPClient(client *http.Client) *ListMinecraftBuildsParams {
	return &ListMinecraftBuildsParams{
		HTTPClient: client,
	}
}

/*
ListMinecraftBuildsParams contains all the parameters to send to the API endpoint

	for the list minecraft builds operation.

	Typically these are written to a http.Request.
*/
type ListMinecraftBuildsParams struct {

	/* MinecraftID.

	   A minecraft UUID or slug
	*/
	MinecraftID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list minecraft builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMinecraftBuildsParams) WithDefaults() *ListMinecraftBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list minecraft builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMinecraftBuildsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list minecraft builds params
func (o *ListMinecraftBuildsParams) WithTimeout(timeout time.Duration) *ListMinecraftBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list minecraft builds params
func (o *ListMinecraftBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list minecraft builds params
func (o *ListMinecraftBuildsParams) WithContext(ctx context.Context) *ListMinecraftBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list minecraft builds params
func (o *ListMinecraftBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list minecraft builds params
func (o *ListMinecraftBuildsParams) WithHTTPClient(client *http.Client) *ListMinecraftBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list minecraft builds params
func (o *ListMinecraftBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMinecraftID adds the minecraftID to the list minecraft builds params
func (o *ListMinecraftBuildsParams) WithMinecraftID(minecraftID string) *ListMinecraftBuildsParams {
	o.SetMinecraftID(minecraftID)
	return o
}

// SetMinecraftID adds the minecraftId to the list minecraft builds params
func (o *ListMinecraftBuildsParams) SetMinecraftID(minecraftID string) {
	o.MinecraftID = minecraftID
}

// WriteToRequest writes these params to a swagger request
func (o *ListMinecraftBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param minecraft_id
	if err := r.SetPathParam("minecraft_id", o.MinecraftID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
