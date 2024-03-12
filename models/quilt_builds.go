// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuiltBuilds quilt builds
//
// swagger:model quilt_builds
type QuiltBuilds struct {

	// builds
	Builds []*Build `json:"builds"`

	// quilt
	Quilt *Quilt `json:"quilt,omitempty"`
}

// Validate validates this quilt builds
func (m *QuiltBuilds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuilt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuiltBuilds) validateBuilds(formats strfmt.Registry) error {
	if swag.IsZero(m.Builds) { // not required
		return nil
	}

	for i := 0; i < len(m.Builds); i++ {
		if swag.IsZero(m.Builds[i]) { // not required
			continue
		}

		if m.Builds[i] != nil {
			if err := m.Builds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("builds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("builds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuiltBuilds) validateQuilt(formats strfmt.Registry) error {
	if swag.IsZero(m.Quilt) { // not required
		return nil
	}

	if m.Quilt != nil {
		if err := m.Quilt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quilt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quilt")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quilt builds based on the context it is used
func (m *QuiltBuilds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuilds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuilt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuiltBuilds) contextValidateBuilds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Builds); i++ {

		if m.Builds[i] != nil {

			if swag.IsZero(m.Builds[i]) { // not required
				return nil
			}

			if err := m.Builds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("builds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("builds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuiltBuilds) contextValidateQuilt(ctx context.Context, formats strfmt.Registry) error {

	if m.Quilt != nil {

		if swag.IsZero(m.Quilt) { // not required
			return nil
		}

		if err := m.Quilt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quilt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quilt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuiltBuilds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuiltBuilds) UnmarshalBinary(b []byte) error {
	var res QuiltBuilds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
