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
	"github.com/go-openapi/validate"
)

// Version version
//
// swagger:model version
type Version struct {

	// builds
	// Read Only: true
	Builds []*BuildVersion `json:"builds,omitempty"`

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// file
	// Read Only: true
	File *VersionFile `json:"file,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// mod
	// Read Only: true
	Mod *Mod `json:"mod,omitempty"`

	// mod id
	// Format: uuid
	ModID strfmt.UUID `json:"mod_id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// public
	Public *bool `json:"public,omitempty"`

	// slug
	Slug *string `json:"slug,omitempty"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this version
func (m *Version) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *Version) validateBuilds(formats strfmt.Registry) error {
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

func (m *Version) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Version) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *Version) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Version) validateMod(formats strfmt.Registry) error {
	if swag.IsZero(m.Mod) { // not required
		return nil
	}

	if m.Mod != nil {
		if err := m.Mod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mod")
			}
			return err
		}
	}

	return nil
}

func (m *Version) validateModID(formats strfmt.Registry) error {
	if swag.IsZero(m.ModID) { // not required
		return nil
	}

	if err := validate.FormatOf("mod_id", "body", "uuid", m.ModID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Version) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Version) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this version based on the context it is used
func (m *Version) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuilds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Version) contextValidateBuilds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "builds", "body", []*BuildVersion(m.Builds)); err != nil {
		return err
	}

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

func (m *Version) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Version) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if m.File != nil {

		if swag.IsZero(m.File) { // not required
			return nil
		}

		if err := m.File.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *Version) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Version) contextValidateMod(ctx context.Context, formats strfmt.Registry) error {

	if m.Mod != nil {

		if swag.IsZero(m.Mod) { // not required
			return nil
		}

		if err := m.Mod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mod")
			}
			return err
		}
	}

	return nil
}

func (m *Version) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Version) UnmarshalBinary(b []byte) error {
	var res Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
