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

// DataclassificationsDataClassificationFinding dataclassifications data classification finding
//
// swagger:model dataclassifications.DataClassificationFinding
type DataclassificationsDataClassificationFinding struct {

	// label Id
	// Required: true
	LabelID *string `json:"labelId"`

	// tag Id
	// Required: true
	TagID *string `json:"tagId"`

	// timestamp
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this dataclassifications data classification finding
func (m *DataclassificationsDataClassificationFinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataclassificationsDataClassificationFinding) validateLabelID(formats strfmt.Registry) error {

	if err := validate.Required("labelId", "body", m.LabelID); err != nil {
		return err
	}

	return nil
}

func (m *DataclassificationsDataClassificationFinding) validateTagID(formats strfmt.Registry) error {

	if err := validate.Required("tagId", "body", m.TagID); err != nil {
		return err
	}

	return nil
}

func (m *DataclassificationsDataClassificationFinding) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dataclassifications data classification finding based on context it is used
func (m *DataclassificationsDataClassificationFinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataclassificationsDataClassificationFinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataclassificationsDataClassificationFinding) UnmarshalBinary(b []byte) error {
	var res DataclassificationsDataClassificationFinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
