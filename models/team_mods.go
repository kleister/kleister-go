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

// TeamMods team mods
//
// swagger:model team_mods
type TeamMods struct {

	// mods
	Mods []*TeamMod `json:"mods"`

	// team
	// Read Only: true
	Team *Team `json:"team,omitempty"`
}

// Validate validates this team mods
func (m *TeamMods) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMods(formats); err != nil {
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

func (m *TeamMods) validateMods(formats strfmt.Registry) error {
	if swag.IsZero(m.Mods) { // not required
		return nil
	}

	for i := 0; i < len(m.Mods); i++ {
		if swag.IsZero(m.Mods[i]) { // not required
			continue
		}

		if m.Mods[i] != nil {
			if err := m.Mods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamMods) validateTeam(formats strfmt.Registry) error {
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

// ContextValidate validate this team mods based on the context it is used
func (m *TeamMods) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMods(ctx, formats); err != nil {
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

func (m *TeamMods) contextValidateMods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mods); i++ {

		if m.Mods[i] != nil {

			if swag.IsZero(m.Mods[i]) { // not required
				return nil
			}

			if err := m.Mods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamMods) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TeamMods) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamMods) UnmarshalBinary(b []byte) error {
	var res TeamMods
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
