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

// ModelsDeliverySettingsRequest models delivery settings request
//
// swagger:model models.DeliverySettingsRequest
type ModelsDeliverySettingsRequest struct {

	// delivery settings
	// Required: true
	DeliverySettings []*ModelsDeliverySettingsInput `json:"delivery_settings"`
}

// Validate validates this models delivery settings request
func (m *ModelsDeliverySettingsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeliverySettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsDeliverySettingsRequest) validateDeliverySettings(formats strfmt.Registry) error {

	if err := validate.Required("delivery_settings", "body", m.DeliverySettings); err != nil {
		return err
	}

	for i := 0; i < len(m.DeliverySettings); i++ {
		if swag.IsZero(m.DeliverySettings[i]) { // not required
			continue
		}

		if m.DeliverySettings[i] != nil {
			if err := m.DeliverySettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delivery_settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delivery_settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this models delivery settings request based on the context it is used
func (m *ModelsDeliverySettingsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeliverySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsDeliverySettingsRequest) contextValidateDeliverySettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeliverySettings); i++ {

		if m.DeliverySettings[i] != nil {

			if swag.IsZero(m.DeliverySettings[i]) { // not required
				return nil
			}

			if err := m.DeliverySettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delivery_settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delivery_settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsDeliverySettingsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsDeliverySettingsRequest) UnmarshalBinary(b []byte) error {
	var res ModelsDeliverySettingsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}