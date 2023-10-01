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

// DomainAzureRoleAssignment domain azure role assignment
//
// swagger:model domain.AzureRoleAssignment
type DomainAzureRoleAssignment struct {

	// name
	Name string `json:"name,omitempty"`

	// role definition id
	RoleDefinitionID string `json:"role_definition_id,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`
}

// Validate validates this domain azure role assignment
func (m *DomainAzureRoleAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAzureRoleAssignment) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain azure role assignment based on context it is used
func (m *DomainAzureRoleAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAzureRoleAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAzureRoleAssignment) UnmarshalBinary(b []byte) error {
	var res DomainAzureRoleAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}