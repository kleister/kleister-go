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

// VersionBuilds version builds
//
// swagger:model version_builds
type VersionBuilds struct {

	// builds
	Builds []*BuildVersion `json:"builds"`

	// mod
	// Read Only: true
	Mod *Mod `json:"mod,omitempty"`

	// version
	// Read Only: true
	Version *Version `json:"version,omitempty"`
}

// Validate validates this version builds
func (m *VersionBuilds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionBuilds) validateBuilds(formats strfmt.Registry) error {
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

func (m *VersionBuilds) validateMod(formats strfmt.Registry) error {
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

func (m *VersionBuilds) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this version builds based on the context it is used
func (m *VersionBuilds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuilds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionBuilds) contextValidateBuilds(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VersionBuilds) contextValidateMod(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VersionBuilds) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {

		if swag.IsZero(m.Version) { // not required
			return nil
		}

		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionBuilds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionBuilds) UnmarshalBinary(b []byte) error {
	var res VersionBuilds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
