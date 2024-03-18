// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModTeamParams mod team params
//
// swagger:model mod_team_params
type ModTeamParams struct {

	// perm
	// Required: true
	// Enum: [user admin owner]
	Perm *string `json:"perm"`

	// team
	// Required: true
	Team *string `json:"team"`
}

// Validate validates this mod team params
func (m *ModTeamParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerm(formats); err != nil {
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

var modTeamParamsTypePermPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modTeamParamsTypePermPropEnum = append(modTeamParamsTypePermPropEnum, v)
	}
}

const (

	// ModTeamParamsPermUser captures enum value "user"
	ModTeamParamsPermUser string = "user"

	// ModTeamParamsPermAdmin captures enum value "admin"
	ModTeamParamsPermAdmin string = "admin"

	// ModTeamParamsPermOwner captures enum value "owner"
	ModTeamParamsPermOwner string = "owner"
)

// prop value enum
func (m *ModTeamParams) validatePermEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modTeamParamsTypePermPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModTeamParams) validatePerm(formats strfmt.Registry) error {

	if err := validate.Required("perm", "body", m.Perm); err != nil {
		return err
	}

	// value enum
	if err := m.validatePermEnum("perm", "body", *m.Perm); err != nil {
		return err
	}

	return nil
}

func (m *ModTeamParams) validateTeam(formats strfmt.Registry) error {

	if err := validate.Required("team", "body", m.Team); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mod team params based on context it is used
func (m *ModTeamParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModTeamParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModTeamParams) UnmarshalBinary(b []byte) error {
	var res ModTeamParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}