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

// UserPack user pack
//
// swagger:model user_pack
type UserPack struct {

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// pack
	// Read Only: true
	Pack *Pack `json:"pack,omitempty"`

	// pack id
	// Required: true
	// Format: uuid
	PackID *strfmt.UUID `json:"pack_id"`

	// perm
	// Required: true
	// Enum: [user admin owner]
	Perm *string `json:"perm"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// user
	// Read Only: true
	User *User `json:"user,omitempty"`

	// user id
	// Required: true
	// Format: uuid
	UserID *strfmt.UUID `json:"user_id"`
}

// Validate validates this user pack
func (m *UserPack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPack) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserPack) validatePack(formats strfmt.Registry) error {
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

func (m *UserPack) validatePackID(formats strfmt.Registry) error {

	if err := validate.Required("pack_id", "body", m.PackID); err != nil {
		return err
	}

	if err := validate.FormatOf("pack_id", "body", "uuid", m.PackID.String(), formats); err != nil {
		return err
	}

	return nil
}

var userPackTypePermPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userPackTypePermPropEnum = append(userPackTypePermPropEnum, v)
	}
}

const (

	// UserPackPermUser captures enum value "user"
	UserPackPermUser string = "user"

	// UserPackPermAdmin captures enum value "admin"
	UserPackPermAdmin string = "admin"

	// UserPackPermOwner captures enum value "owner"
	UserPackPermOwner string = "owner"
)

// prop value enum
func (m *UserPack) validatePermEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userPackTypePermPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserPack) validatePerm(formats strfmt.Registry) error {

	if err := validate.Required("perm", "body", m.Perm); err != nil {
		return err
	}

	// value enum
	if err := m.validatePermEnum("perm", "body", *m.Perm); err != nil {
		return err
	}

	return nil
}

func (m *UserPack) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserPack) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *UserPack) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user pack based on the context it is used
func (m *UserPack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPack) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *UserPack) contextValidatePack(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserPack) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *UserPack) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserPack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPack) UnmarshalBinary(b []byte) error {
	var res UserPack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
