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

// WebhooktriggerAuthConfig webhooktrigger auth config
//
// swagger:model webhooktrigger.AuthConfig
type WebhooktriggerAuthConfig struct {

	// API key configuration
	APIKeyConfig *WebhooktriggerAPIKeyConfig `json:"api_key_config,omitempty"`

	// The type of authentication to use for the webhook trigger
	// Required: true
	AuthType *string `json:"auth_type"`

	// Basic authentication configuration
	BasicAuthConfig *WebhooktriggerBasicAuthConfig `json:"basic_auth_config,omitempty"`

	// HMAC authentication configuration
	HmacConfig *WebhooktriggerHMACConfig `json:"hmac_config,omitempty"`
}

// Validate validates this webhooktrigger auth config
func (m *WebhooktriggerAuthConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKeyConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBasicAuthConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHmacConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhooktriggerAuthConfig) validateAPIKeyConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.APIKeyConfig) { // not required
		return nil
	}

	if m.APIKeyConfig != nil {
		if err := m.APIKeyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_key_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api_key_config")
			}
			return err
		}
	}

	return nil
}

func (m *WebhooktriggerAuthConfig) validateAuthType(formats strfmt.Registry) error {

	if err := validate.Required("auth_type", "body", m.AuthType); err != nil {
		return err
	}

	return nil
}

func (m *WebhooktriggerAuthConfig) validateBasicAuthConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.BasicAuthConfig) { // not required
		return nil
	}

	if m.BasicAuthConfig != nil {
		if err := m.BasicAuthConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth_config")
			}
			return err
		}
	}

	return nil
}

func (m *WebhooktriggerAuthConfig) validateHmacConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HmacConfig) { // not required
		return nil
	}

	if m.HmacConfig != nil {
		if err := m.HmacConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hmac_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hmac_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this webhooktrigger auth config based on the context it is used
func (m *WebhooktriggerAuthConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIKeyConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBasicAuthConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHmacConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhooktriggerAuthConfig) contextValidateAPIKeyConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.APIKeyConfig != nil {

		if swag.IsZero(m.APIKeyConfig) { // not required
			return nil
		}

		if err := m.APIKeyConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_key_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api_key_config")
			}
			return err
		}
	}

	return nil
}

func (m *WebhooktriggerAuthConfig) contextValidateBasicAuthConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.BasicAuthConfig != nil {

		if swag.IsZero(m.BasicAuthConfig) { // not required
			return nil
		}

		if err := m.BasicAuthConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth_config")
			}
			return err
		}
	}

	return nil
}

func (m *WebhooktriggerAuthConfig) contextValidateHmacConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HmacConfig != nil {

		if swag.IsZero(m.HmacConfig) { // not required
			return nil
		}

		if err := m.HmacConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hmac_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hmac_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebhooktriggerAuthConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhooktriggerAuthConfig) UnmarshalBinary(b []byte) error {
	var res WebhooktriggerAuthConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
