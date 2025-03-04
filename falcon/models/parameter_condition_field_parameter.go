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

// ParameterConditionFieldParameter parameter condition field parameter
//
// swagger:model parameter.ConditionFieldParameter
type ParameterConditionFieldParameter struct {

	// Optional default operator to be used as part of the condition
	DefaultOperator string `json:"default_operator,omitempty"`

	// Optional default value used for the condition, type is dynamic depending on the underlying field.
	DefaultValue ParameterConditionFieldParameterDefaultValue `json:"default_value,omitempty"`

	// Optional text/description which can be used to provide differentiation for parameterized fields during app installation.
	HelperText string `json:"helperText,omitempty"`

	// Whether the field can be specified multiple times as provisioning parameter. When true, all values or combined via an OR operator.
	// Required: true
	Multiple *bool `json:"multiple"`

	// The default operator that should be applied for this field.
	Operator string `json:"operator,omitempty"`

	// Indicates whether the field must be specified as a parameter at provision time.
	// Required: true
	Required *bool `json:"required"`
}

// Validate validates this parameter condition field parameter
func (m *ParameterConditionFieldParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMultiple(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterConditionFieldParameter) validateMultiple(formats strfmt.Registry) error {

	if err := validate.Required("multiple", "body", m.Multiple); err != nil {
		return err
	}

	return nil
}

func (m *ParameterConditionFieldParameter) validateRequired(formats strfmt.Registry) error {

	if err := validate.Required("required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this parameter condition field parameter based on context it is used
func (m *ParameterConditionFieldParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterConditionFieldParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterConditionFieldParameter) UnmarshalBinary(b []byte) error {
	var res ParameterConditionFieldParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
