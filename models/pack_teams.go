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

// PackTeams pack teams
//
// swagger:model pack_teams
type PackTeams struct {

	// pack
	// Read Only: true
	Pack *Pack `json:"pack,omitempty"`

	// teams
	Teams []*TeamPack `json:"teams"`
}

// Validate validates this pack teams
func (m *PackTeams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackTeams) validatePack(formats strfmt.Registry) error {
	if swag.IsZero(m.Pack) { // not required
		return nil
	}

	if m.Pack != nil {
		if err := m.Pack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pack")
			}
			return err
		}
	}

	return nil
}

func (m *PackTeams) validateTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pack teams based on the context it is used
func (m *PackTeams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackTeams) contextValidatePack(ctx context.Context, formats strfmt.Registry) error {

	if m.Pack != nil {

		if swag.IsZero(m.Pack) { // not required
			return nil
		}

		if err := m.Pack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pack")
			}
			return err
		}
	}

	return nil
}

func (m *PackTeams) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Teams); i++ {

		if m.Teams[i] != nil {

			if swag.IsZero(m.Teams[i]) { // not required
				return nil
			}

			if err := m.Teams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackTeams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackTeams) UnmarshalBinary(b []byte) error {
	var res PackTeams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
