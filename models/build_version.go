// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildVersion build version
//
// swagger:model build_version
type BuildVersion struct {

	// build
	// Read Only: true
	Build *Build `json:"build,omitempty"`

	// build id
	// Required: true
	// Format: uuid
	BuildID *strfmt.UUID `json:"build_id"`

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// version
	// Read Only: true
	Version *Version `json:"version,omitempty"`

	// version id
	// Required: true
	// Format: uuid
	VersionID *strfmt.UUID `json:"version_id"`
}

// Validate validates this build version
func (m *BuildVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildVersion) validateBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

func (m *BuildVersion) validateBuildID(formats strfmt.Registry) error {

	if err := validate.Required("build_id", "body", m.BuildID); err != nil {
		return err
	}

	if err := validate.FormatOf("build_id", "body", "uuid", m.BuildID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildVersion) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildVersion) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildVersion) validateVersion(formats strfmt.Registry) error {
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

func (m *BuildVersion) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("version_id", "body", m.VersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("version_id", "body", "uuid", m.VersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this build version based on the context it is used
func (m *BuildVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
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

func (m *BuildVersion) contextValidateBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.Build != nil {

		if swag.IsZero(m.Build) { // not required
			return nil
		}

		if err := m.Build.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

func (m *BuildVersion) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *BuildVersion) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *BuildVersion) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BuildVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildVersion) UnmarshalBinary(b []byte) error {
	var res BuildVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
