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

// RestAccountProductUpsertRequestExtV1 rest account product upsert request ext v1
//
// swagger:model rest.AccountProductUpsertRequestExtV1
type RestAccountProductUpsertRequestExtV1 struct {

	// features
	// Required: true
	Features []string `json:"features"`

	// product
	// Required: true
	Product *string `json:"product"`
}

// Validate validates this rest account product upsert request ext v1
func (m *RestAccountProductUpsertRequestExtV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAccountProductUpsertRequestExtV1) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *RestAccountProductUpsertRequestExtV1) validateProduct(formats strfmt.Registry) error {

	if err := validate.Required("product", "body", m.Product); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rest account product upsert request ext v1 based on context it is used
func (m *RestAccountProductUpsertRequestExtV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestAccountProductUpsertRequestExtV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestAccountProductUpsertRequestExtV1) UnmarshalBinary(b []byte) error {
	var res RestAccountProductUpsertRequestExtV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}