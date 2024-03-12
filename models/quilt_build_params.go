// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QuiltBuildParams quilt build params
//
// swagger:model quilt_build_params
type QuiltBuildParams struct {

	// build
	// Required: true
	Build *string `json:"build"`

	// pack
	// Required: true
	Pack *string `json:"pack"`
}

// Validate validates this quilt build params
func (m *QuiltBuildParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuiltBuildParams) validateBuild(formats strfmt.Registry) error {

	if err := validate.Required("build", "body", m.Build); err != nil {
		return err
	}

	return nil
}

func (m *QuiltBuildParams) validatePack(formats strfmt.Registry) error {

	if err := validate.Required("pack", "body", m.Pack); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this quilt build params based on context it is used
func (m *QuiltBuildParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuiltBuildParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuiltBuildParams) UnmarshalBinary(b []byte) error {
	var res QuiltBuildParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
