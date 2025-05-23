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

// RegistrationGCPServiceAccountExtV1 registration g c p service account ext v1
//
// swagger:model registration.GCPServiceAccountExtV1
type RegistrationGCPServiceAccountExtV1 struct {

	// client email
	ClientEmail string `json:"client_email,omitempty"`

	// client id
	ClientID string `json:"client_id,omitempty"`

	// project id
	ProjectID string `json:"project_id,omitempty"`

	// service account conditions
	ServiceAccountConditions []*StatemgmtCondition `json:"service_account_conditions"`

	// service account id
	ServiceAccountID int64 `json:"service_account_id,omitempty"`
}

// Validate validates this registration g c p service account ext v1
func (m *RegistrationGCPServiceAccountExtV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceAccountConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationGCPServiceAccountExtV1) validateServiceAccountConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAccountConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceAccountConditions); i++ {
		if swag.IsZero(m.ServiceAccountConditions[i]) { // not required
			continue
		}

		if m.ServiceAccountConditions[i] != nil {
			if err := m.ServiceAccountConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_account_conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_account_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this registration g c p service account ext v1 based on the context it is used
func (m *RegistrationGCPServiceAccountExtV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceAccountConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationGCPServiceAccountExtV1) contextValidateServiceAccountConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceAccountConditions); i++ {

		if m.ServiceAccountConditions[i] != nil {

			if swag.IsZero(m.ServiceAccountConditions[i]) { // not required
				return nil
			}

			if err := m.ServiceAccountConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_account_conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_account_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationGCPServiceAccountExtV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationGCPServiceAccountExtV1) UnmarshalBinary(b []byte) error {
	var res RegistrationGCPServiceAccountExtV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
