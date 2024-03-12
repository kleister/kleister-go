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

// Pack pack
//
// swagger:model pack
type Pack struct {

	// back
	// Read Only: true
	Back *PackBack `json:"back,omitempty"`

	// builds
	// Read Only: true
	Builds []*Build `json:"builds,omitempty"`

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// icon
	// Read Only: true
	Icon *PackIcon `json:"icon,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// latest
	// Read Only: true
	Latest *Build `json:"latest,omitempty"`

	// latest id
	// Format: uuid
	LatestID *strfmt.UUID `json:"latest_id,omitempty"`

	// logo
	// Read Only: true
	Logo *PackLogo `json:"logo,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// public
	Public *bool `json:"public,omitempty"`

	// recommended
	// Read Only: true
	Recommended *Build `json:"recommended,omitempty"`

	// recommended id
	// Format: uuid
	RecommendedID *strfmt.UUID `json:"recommended_id,omitempty"`

	// slug
	Slug *string `json:"slug,omitempty"`

	// teams
	// Read Only: true
	Teams []*TeamPack `json:"teams,omitempty"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// users
	// Read Only: true
	Users []*UserPack `json:"users,omitempty"`

	// website
	Website *string `json:"website,omitempty"`
}

// Validate validates this pack
func (m *Pack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecommended(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecommendedID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pack) validateBack(formats strfmt.Registry) error {
	if swag.IsZero(m.Back) { // not required
		return nil
	}

	if m.Back != nil {
		if err := m.Back.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("back")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("back")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) validateBuilds(formats strfmt.Registry) error {
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

func (m *Pack) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Pack) validateIcon(formats strfmt.Registry) error {
	if swag.IsZero(m.Icon) { // not required
		return nil
	}

	if m.Icon != nil {
		if err := m.Icon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("icon")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Pack) validateLatest(formats strfmt.Registry) error {
	if swag.IsZero(m.Latest) { // not required
		return nil
	}

	if m.Latest != nil {
		if err := m.Latest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) validateLatestID(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestID) { // not required
		return nil
	}

	if err := validate.FormatOf("latest_id", "body", "uuid", m.LatestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Pack) validateLogo(formats strfmt.Registry) error {
	if swag.IsZero(m.Logo) { // not required
		return nil
	}

	if m.Logo != nil {
		if err := m.Logo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Pack) validateRecommended(formats strfmt.Registry) error {
	if swag.IsZero(m.Recommended) { // not required
		return nil
	}

	if m.Recommended != nil {
		if err := m.Recommended.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recommended")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recommended")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) validateRecommendedID(formats strfmt.Registry) error {
	if swag.IsZero(m.RecommendedID) { // not required
		return nil
	}

	if err := validate.FormatOf("recommended_id", "body", "uuid", m.RecommendedID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Pack) validateTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Pack) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Pack) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pack based on the context it is used
func (m *Pack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuilds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIcon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecommended(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pack) contextValidateBack(ctx context.Context, formats strfmt.Registry) error {

	if m.Back != nil {

		if swag.IsZero(m.Back) { // not required
			return nil
		}

		if err := m.Back.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("back")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("back")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) contextValidateBuilds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "builds", "body", []*Build(m.Builds)); err != nil {
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

func (m *Pack) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Pack) contextValidateIcon(ctx context.Context, formats strfmt.Registry) error {

	if m.Icon != nil {

		if swag.IsZero(m.Icon) { // not required
			return nil
		}

		if err := m.Icon.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("icon")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Pack) contextValidateLatest(ctx context.Context, formats strfmt.Registry) error {

	if m.Latest != nil {

		if swag.IsZero(m.Latest) { // not required
			return nil
		}

		if err := m.Latest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) contextValidateLogo(ctx context.Context, formats strfmt.Registry) error {

	if m.Logo != nil {

		if swag.IsZero(m.Logo) { // not required
			return nil
		}

		if err := m.Logo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) contextValidateRecommended(ctx context.Context, formats strfmt.Registry) error {

	if m.Recommended != nil {

		if swag.IsZero(m.Recommended) { // not required
			return nil
		}

		if err := m.Recommended.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recommended")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recommended")
			}
			return err
		}
	}

	return nil
}

func (m *Pack) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "teams", "body", []*TeamPack(m.Teams)); err != nil {
		return err
	}

	for i := 0; i < len(m.Teams); i++ {

		if m.Teams[i] != nil {

			if swag.IsZero(m.Teams[i]) { // not required
				return nil
			}

			if err := m.Teams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Pack) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Pack) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "users", "body", []*UserPack(m.Users)); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {

			if swag.IsZero(m.Users[i]) { // not required
				return nil
			}

			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Pack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pack) UnmarshalBinary(b []byte) error {
	var res Pack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
