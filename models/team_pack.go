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

// TeamPack team pack
// swagger:model team_pack
type TeamPack struct {

	// pack id
	// Required: true
	// Format: uuid
	PackID *strfmt.UUID `json:"pack_id"`

	// perm
	// Required: true
	// Enum: [user admin owner]
	Perm *string `json:"perm"`

	// team id
	// Required: true
	// Format: uuid
	TeamID *strfmt.UUID `json:"team_id"`
}

// Validate validates this team pack
func (m *TeamPack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamPack) validatePackID(formats strfmt.Registry) error {

	if err := validate.Required("pack_id", "body", m.PackID); err != nil {
		return err
	}

	if err := validate.FormatOf("pack_id", "body", "uuid", m.PackID.String(), formats); err != nil {
		return err
	}

	return nil
}

var teamPackTypePermPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		teamPackTypePermPropEnum = append(teamPackTypePermPropEnum, v)
	}
}

const (

	// TeamPackPermUser captures enum value "user"
	TeamPackPermUser string = "user"

	// TeamPackPermAdmin captures enum value "admin"
	TeamPackPermAdmin string = "admin"

	// TeamPackPermOwner captures enum value "owner"
	TeamPackPermOwner string = "owner"
)

// prop value enum
func (m *TeamPack) validatePermEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, teamPackTypePermPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TeamPack) validatePerm(formats strfmt.Registry) error {

	if err := validate.Required("perm", "body", m.Perm); err != nil {
		return err
	}

	// value enum
	if err := m.validatePermEnum("perm", "body", *m.Perm); err != nil {
		return err
	}

	return nil
}

func (m *TeamPack) validateTeamID(formats strfmt.Registry) error {

	if err := validate.Required("team_id", "body", m.TeamID); err != nil {
		return err
	}

	if err := validate.FormatOf("team_id", "body", "uuid", m.TeamID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamPack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamPack) UnmarshalBinary(b []byte) error {
	var res TeamPack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}