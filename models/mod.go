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

// Mod mod
//
// swagger:model mod
type Mod struct {

	// author
	Author string `json:"author,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// donate
	Donate string `json:"donate,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// side
	// Enum: [both server client]
	Side string `json:"side,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this mod
func (m *Mod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mod) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mod) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mod) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var modTypeSidePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["both","server","client"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modTypeSidePropEnum = append(modTypeSidePropEnum, v)
	}
}

const (

	// ModSideBoth captures enum value "both"
	ModSideBoth string = "both"

	// ModSideServer captures enum value "server"
	ModSideServer string = "server"

	// ModSideClient captures enum value "client"
	ModSideClient string = "client"
)

// prop value enum
func (m *Mod) validateSideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modTypeSidePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Mod) validateSide(formats strfmt.Registry) error {
	if swag.IsZero(m.Side) { // not required
		return nil
	}

	// value enum
	if err := m.validateSideEnum("side", "body", m.Side); err != nil {
		return err
	}

	return nil
}

func (m *Mod) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mod based on the context it is used
func (m *Mod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mod) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Mod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mod) UnmarshalBinary(b []byte) error {
	var res Mod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
