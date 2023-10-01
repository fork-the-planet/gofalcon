// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PreventionCreatePolicyReqV1 prevention create policy req v1
//
// swagger:model prevention.CreatePolicyReqV1
type PreventionCreatePolicyReqV1 struct {

	// If specified the settings of the prevention policy identified by the id will be used
	CloneID string `json:"clone_id,omitempty"`

	// The description to use when creating the policy
	Description string `json:"description,omitempty"`

	// The name to use when creating the policy
	// Required: true
	Name *string `json:"name"`

	// The name of the platform
	// Required: true
	// Enum: [Windows Mac Linux]
	PlatformName *string `json:"platform_name"`

	// The settings to create the policy with
	Settings []*PreventionSettingReqV1 `json:"settings"`
}

// Validate validates this prevention create policy req v1
func (m *PreventionCreatePolicyReqV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreventionCreatePolicyReqV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var preventionCreatePolicyReqV1TypePlatformNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Windows","Mac","Linux"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		preventionCreatePolicyReqV1TypePlatformNamePropEnum = append(preventionCreatePolicyReqV1TypePlatformNamePropEnum, v)
	}
}

const (

	// PreventionCreatePolicyReqV1PlatformNameWindows captures enum value "Windows"
	PreventionCreatePolicyReqV1PlatformNameWindows string = "Windows"

	// PreventionCreatePolicyReqV1PlatformNameMac captures enum value "Mac"
	PreventionCreatePolicyReqV1PlatformNameMac string = "Mac"

	// PreventionCreatePolicyReqV1PlatformNameLinux captures enum value "Linux"
	PreventionCreatePolicyReqV1PlatformNameLinux string = "Linux"
)

// prop value enum
func (m *PreventionCreatePolicyReqV1) validatePlatformNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, preventionCreatePolicyReqV1TypePlatformNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PreventionCreatePolicyReqV1) validatePlatformName(formats strfmt.Registry) error {

	if err := validate.Required("platform_name", "body", m.PlatformName); err != nil {
		return err
	}

	// value enum
	if err := m.validatePlatformNameEnum("platform_name", "body", *m.PlatformName); err != nil {
		return err
	}

	return nil
}

func (m *PreventionCreatePolicyReqV1) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	for i := 0; i < len(m.Settings); i++ {
		if swag.IsZero(m.Settings[i]) { // not required
			continue
		}

		if m.Settings[i] != nil {
			if err := m.Settings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this prevention create policy req v1 based on the context it is used
func (m *PreventionCreatePolicyReqV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreventionCreatePolicyReqV1) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Settings); i++ {

		if m.Settings[i] != nil {

			if swag.IsZero(m.Settings[i]) { // not required
				return nil
			}

			if err := m.Settings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreventionCreatePolicyReqV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreventionCreatePolicyReqV1) UnmarshalBinary(b []byte) error {
	var res PreventionCreatePolicyReqV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}