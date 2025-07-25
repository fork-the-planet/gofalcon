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

// DomainAPIAggregateResponseComplianceByClusterTypeV1Resources domain API aggregate response compliance by cluster type v1 resources
//
// swagger:model domain.APIAggregateResponseComplianceByClusterTypeV1.resources
type DomainAPIAggregateResponseComplianceByClusterTypeV1Resources struct {

	// buckets
	// Required: true
	Buckets []*DomainAPIAggregateComplianceByClusterTypeV1 `json:"buckets"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this domain API aggregate response compliance by cluster type v1 resources
func (m *DomainAPIAggregateResponseComplianceByClusterTypeV1Resources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuckets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIAggregateResponseComplianceByClusterTypeV1Resources) validateBuckets(formats strfmt.Registry) error {

	if err := validate.Required("buckets", "body", m.Buckets); err != nil {
		return err
	}

	for i := 0; i < len(m.Buckets); i++ {
		if swag.IsZero(m.Buckets[i]) { // not required
			continue
		}

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainAPIAggregateResponseComplianceByClusterTypeV1Resources) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain API aggregate response compliance by cluster type v1 resources based on the context it is used
func (m *DomainAPIAggregateResponseComplianceByClusterTypeV1Resources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuckets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIAggregateResponseComplianceByClusterTypeV1Resources) contextValidateBuckets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Buckets); i++ {

		if m.Buckets[i] != nil {

			if swag.IsZero(m.Buckets[i]) { // not required
				return nil
			}

			if err := m.Buckets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIAggregateResponseComplianceByClusterTypeV1Resources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIAggregateResponseComplianceByClusterTypeV1Resources) UnmarshalBinary(b []byte) error {
	var res DomainAPIAggregateResponseComplianceByClusterTypeV1Resources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
