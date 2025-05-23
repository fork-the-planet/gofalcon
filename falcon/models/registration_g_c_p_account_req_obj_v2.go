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

// RegistrationGCPAccountReqObjV2 registration g c p account req obj v2
//
// swagger:model registration.GCPAccountReqObjV2
type RegistrationGCPAccountReqObjV2 struct {

	// client email
	ClientEmail string `json:"client_email,omitempty"`

	// client id
	ClientID string `json:"client_id,omitempty"`

	// parent id
	// Required: true
	ParentID *string `json:"parent_id"`

	// parent type
	ParentType string `json:"parent_type,omitempty"`

	// private key
	PrivateKey string `json:"private_key,omitempty"`

	// private key id
	PrivateKeyID string `json:"private_key_id,omitempty"`

	// project id
	ProjectID string `json:"project_id,omitempty"`

	// service account conditions
	ServiceAccountConditions []*StatemgmtCondition `json:"service_account_conditions"`

	// service account id
	ServiceAccountID int64 `json:"service_account_id,omitempty"`
}

// Validate validates this registration g c p account req obj v2
func (m *RegistrationGCPAccountReqObjV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccountConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationGCPAccountReqObjV2) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationGCPAccountReqObjV2) validateServiceAccountConditions(formats strfmt.Registry) error {
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

// ContextValidate validate this registration g c p account req obj v2 based on the context it is used
func (m *RegistrationGCPAccountReqObjV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceAccountConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationGCPAccountReqObjV2) contextValidateServiceAccountConditions(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RegistrationGCPAccountReqObjV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationGCPAccountReqObjV2) UnmarshalBinary(b []byte) error {
	var res RegistrationGCPAccountReqObjV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
