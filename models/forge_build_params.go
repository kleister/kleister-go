// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ForgeBuildParams forge build params
// swagger:model forge_build_params
type ForgeBuildParams struct {

	// build
	// Required: true
	Build *string `json:"build"`

	// pack
	// Required: true
	Pack *string `json:"pack"`
}

// Validate validates this forge build params
func (m *ForgeBuildParams) Validate(formats strfmt.Registry) error {
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

func (m *ForgeBuildParams) validateBuild(formats strfmt.Registry) error {

	if err := validate.Required("build", "body", m.Build); err != nil {
		return err
	}

	return nil
}

func (m *ForgeBuildParams) validatePack(formats strfmt.Registry) error {

	if err := validate.Required("pack", "body", m.Pack); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ForgeBuildParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ForgeBuildParams) UnmarshalBinary(b []byte) error {
	var res ForgeBuildParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}