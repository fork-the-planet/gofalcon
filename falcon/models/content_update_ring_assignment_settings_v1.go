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

// ContentUpdateRingAssignmentSettingsV1 content update ring assignment settings v1
//
// swagger:model content_update.RingAssignmentSettingsV1
type ContentUpdateRingAssignmentSettingsV1 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// ring assignment
	// Required: true
	RingAssignment *string `json:"ring_assignment"`
}

// Validate validates this content update ring assignment settings v1
func (m *ContentUpdateRingAssignmentSettingsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRingAssignment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentUpdateRingAssignmentSettingsV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ContentUpdateRingAssignmentSettingsV1) validateRingAssignment(formats strfmt.Registry) error {

	if err := validate.Required("ring_assignment", "body", m.RingAssignment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this content update ring assignment settings v1 based on context it is used
func (m *ContentUpdateRingAssignmentSettingsV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContentUpdateRingAssignmentSettingsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentUpdateRingAssignmentSettingsV1) UnmarshalBinary(b []byte) error {
	var res ContentUpdateRingAssignmentSettingsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
