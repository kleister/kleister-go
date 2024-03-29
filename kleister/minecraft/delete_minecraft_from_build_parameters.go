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

	"github.com/kleister/kleister-go/models"
)

// NewDeleteMinecraftFromBuildParams creates a new DeleteMinecraftFromBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMinecraftFromBuildParams() *DeleteMinecraftFromBuildParams {
	return &DeleteMinecraftFromBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMinecraftFromBuildParamsWithTimeout creates a new DeleteMinecraftFromBuildParams object
// with the ability to set a timeout on a request.
func NewDeleteMinecraftFromBuildParamsWithTimeout(timeout time.Duration) *DeleteMinecraftFromBuildParams {
	return &DeleteMinecraftFromBuildParams{
		timeout: timeout,
	}
}

// NewDeleteMinecraftFromBuildParamsWithContext creates a new DeleteMinecraftFromBuildParams object
// with the ability to set a context for a request.
func NewDeleteMinecraftFromBuildParamsWithContext(ctx context.Context) *DeleteMinecraftFromBuildParams {
	return &DeleteMinecraftFromBuildParams{
		Context: ctx,
	}
}

// NewDeleteMinecraftFromBuildParamsWithHTTPClient creates a new DeleteMinecraftFromBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMinecraftFromBuildParamsWithHTTPClient(client *http.Client) *DeleteMinecraftFromBuildParams {
	return &DeleteMinecraftFromBuildParams{
		HTTPClient: client,
	}
}

/*
DeleteMinecraftFromBuildParams contains all the parameters to send to the API endpoint

	for the delete minecraft from build operation.

	Typically these are written to a http.Request.
*/
type DeleteMinecraftFromBuildParams struct {

	/* MinecraftBuild.

	   The build data to unlink
	*/
	MinecraftBuild *models.MinecraftBuildParams

	/* MinecraftID.

	   A minecraft UUID or slug
	*/
	MinecraftID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete minecraft from build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMinecraftFromBuildParams) WithDefaults() *DeleteMinecraftFromBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete minecraft from build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMinecraftFromBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) WithTimeout(timeout time.Duration) *DeleteMinecraftFromBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) WithContext(ctx context.Context) *DeleteMinecraftFromBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) WithHTTPClient(client *http.Client) *DeleteMinecraftFromBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMinecraftBuild adds the minecraftBuild to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) WithMinecraftBuild(minecraftBuild *models.MinecraftBuildParams) *DeleteMinecraftFromBuildParams {
	o.SetMinecraftBuild(minecraftBuild)
	return o
}

// SetMinecraftBuild adds the minecraftBuild to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) SetMinecraftBuild(minecraftBuild *models.MinecraftBuildParams) {
	o.MinecraftBuild = minecraftBuild
}

// WithMinecraftID adds the minecraftID to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) WithMinecraftID(minecraftID string) *DeleteMinecraftFromBuildParams {
	o.SetMinecraftID(minecraftID)
	return o
}

// SetMinecraftID adds the minecraftId to the delete minecraft from build params
func (o *DeleteMinecraftFromBuildParams) SetMinecraftID(minecraftID string) {
	o.MinecraftID = minecraftID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMinecraftFromBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.MinecraftBuild != nil {
		if err := r.SetBodyParam(o.MinecraftBuild); err != nil {
			return err
		}
	}

	// path param minecraft_id
	if err := r.SetPathParam("minecraft_id", o.MinecraftID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
