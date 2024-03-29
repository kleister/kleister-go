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

// TeamPacks team packs
//
// swagger:model team_packs
type TeamPacks struct {

	// packs
	Packs []*TeamPack `json:"packs"`

	// team
	// Read Only: true
	Team *Team `json:"team,omitempty"`
}

// Validate validates this team packs
func (m *TeamPacks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamPacks) validatePacks(formats strfmt.Registry) error {
	if swag.IsZero(m.Packs) { // not required
		return nil
	}

	for i := 0; i < len(m.Packs); i++ {
		if swag.IsZero(m.Packs[i]) { // not required
			continue
		}

		if m.Packs[i] != nil {
			if err := m.Packs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamPacks) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this team packs based on the context it is used
func (m *TeamPacks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamPacks) contextValidatePacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Packs); i++ {

		if m.Packs[i] != nil {

			if swag.IsZero(m.Packs[i]) { // not required
				return nil
			}

			if err := m.Packs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamPacks) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {

		if swag.IsZero(m.Team) { // not required
			return nil
		}

		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamPacks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamPacks) UnmarshalBinary(b []byte) error {
	var res TeamPacks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
