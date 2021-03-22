// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PackTeamParams pack team params
// swagger:model pack_team_params
type PackTeamParams struct {

	// perm
	// Required: true
	// Enum: [user admin owner]
	Perm *string `json:"perm"`

	// team
	// Required: true
	Team *string `json:"team"`
}

// Validate validates this pack team params
func (m *PackTeamParams) Validate(formats strfmt.Registry) error {
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

var packTeamParamsTypePermPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packTeamParamsTypePermPropEnum = append(packTeamParamsTypePermPropEnum, v)
	}
}

const (

	// PackTeamParamsPermUser captures enum value "user"
	PackTeamParamsPermUser string = "user"

	// PackTeamParamsPermAdmin captures enum value "admin"
	PackTeamParamsPermAdmin string = "admin"

	// PackTeamParamsPermOwner captures enum value "owner"
	PackTeamParamsPermOwner string = "owner"
)

// prop value enum
func (m *PackTeamParams) validatePermEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, packTeamParamsTypePermPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PackTeamParams) validatePerm(formats strfmt.Registry) error {

	if err := validate.Required("perm", "body", m.Perm); err != nil {
		return err
	}

	// value enum
	if err := m.validatePermEnum("perm", "body", *m.Perm); err != nil {
		return err
	}

	return nil
}

func (m *PackTeamParams) validateTeam(formats strfmt.Registry) error {

	if err := validate.Required("team", "body", m.Team); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackTeamParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackTeamParams) UnmarshalBinary(b []byte) error {
	var res PackTeamParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}