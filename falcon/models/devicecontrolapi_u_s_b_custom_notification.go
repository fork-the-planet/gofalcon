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

// DevicecontrolapiUSBCustomNotification devicecontrolapi u s b custom notification
//
// swagger:model devicecontrolapi.USBCustomNotification
type DevicecontrolapiUSBCustomNotification struct {

	// custom message
	// Required: true
	CustomMessage *string `json:"custom_message"`

	// use custom
	// Required: true
	UseCustom *bool `json:"use_custom"`
}

// Validate validates this devicecontrolapi u s b custom notification
func (m *DevicecontrolapiUSBCustomNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseCustom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevicecontrolapiUSBCustomNotification) validateCustomMessage(formats strfmt.Registry) error {

	if err := validate.Required("custom_message", "body", m.CustomMessage); err != nil {
		return err
	}

	return nil
}

func (m *DevicecontrolapiUSBCustomNotification) validateUseCustom(formats strfmt.Registry) error {

	if err := validate.Required("use_custom", "body", m.UseCustom); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this devicecontrolapi u s b custom notification based on context it is used
func (m *DevicecontrolapiUSBCustomNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DevicecontrolapiUSBCustomNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevicecontrolapiUSBCustomNotification) UnmarshalBinary(b []byte) error {
	var res DevicecontrolapiUSBCustomNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
