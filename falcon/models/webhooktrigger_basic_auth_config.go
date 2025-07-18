// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhooktriggerBasicAuthConfig webhooktrigger basic auth config
//
// swagger:model webhooktrigger.BasicAuthConfig
type WebhooktriggerBasicAuthConfig struct {

	// Password for basic authentication
	Password string `json:"password,omitempty"`

	// Username for basic authentication
	Username string `json:"username,omitempty"`
}

// Validate validates this webhooktrigger basic auth config
func (m *WebhooktriggerBasicAuthConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webhooktrigger basic auth config based on context it is used
func (m *WebhooktriggerBasicAuthConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhooktriggerBasicAuthConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhooktriggerBasicAuthConfig) UnmarshalBinary(b []byte) error {
	var res WebhooktriggerBasicAuthConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
