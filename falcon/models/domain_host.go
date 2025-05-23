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

// DomainHost domain host
//
// swagger:model domain.Host
type DomainHost struct {

	// cloud ioas
	CloudIoas *DomainIOACounts `json:"cloud_ioas,omitempty"`

	// cloud ioms
	CloudIoms *DomainIOMCounts `json:"cloud_ioms,omitempty"`

	// configuration assessments
	ConfigurationAssessments *DomainSCAMisconfigsCounts `json:"configuration_assessments,omitempty"`

	// entity p k
	EntityPK string `json:"entityPK,omitempty"`

	// extra info
	ExtraInfo *DomainXLR8Info `json:"extra_info,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// sensitive data
	// Required: true
	SensitiveData *bool `json:"sensitive_data"`

	// total count
	// Required: true
	TotalCount *float64 `json:"total_count"`

	// vulnerabilities
	Vulnerabilities *DomainVulnerabilitiesCount `json:"vulnerabilities,omitempty"`
}

// Validate validates this domain host
func (m *DomainHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudIoas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudIoms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationAssessments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensitiveData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainHost) validateCloudIoas(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudIoas) { // not required
		return nil
	}

	if m.CloudIoas != nil {
		if err := m.CloudIoas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_ioas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_ioas")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) validateCloudIoms(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudIoms) { // not required
		return nil
	}

	if m.CloudIoms != nil {
		if err := m.CloudIoms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_ioms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_ioms")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) validateConfigurationAssessments(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigurationAssessments) { // not required
		return nil
	}

	if m.ConfigurationAssessments != nil {
		if err := m.ConfigurationAssessments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration_assessments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration_assessments")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) validateExtraInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraInfo) { // not required
		return nil
	}

	if m.ExtraInfo != nil {
		if err := m.ExtraInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extra_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extra_info")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainHost) validateSensitiveData(formats strfmt.Registry) error {

	if err := validate.Required("sensitive_data", "body", m.SensitiveData); err != nil {
		return err
	}

	return nil
}

func (m *DomainHost) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

func (m *DomainHost) validateVulnerabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	if m.Vulnerabilities != nil {
		if err := m.Vulnerabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vulnerabilities")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain host based on the context it is used
func (m *DomainHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudIoas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudIoms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigurationAssessments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtraInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVulnerabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainHost) contextValidateCloudIoas(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudIoas != nil {

		if swag.IsZero(m.CloudIoas) { // not required
			return nil
		}

		if err := m.CloudIoas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_ioas")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_ioas")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) contextValidateCloudIoms(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudIoms != nil {

		if swag.IsZero(m.CloudIoms) { // not required
			return nil
		}

		if err := m.CloudIoms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_ioms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_ioms")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) contextValidateConfigurationAssessments(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigurationAssessments != nil {

		if swag.IsZero(m.ConfigurationAssessments) { // not required
			return nil
		}

		if err := m.ConfigurationAssessments.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration_assessments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration_assessments")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) contextValidateExtraInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtraInfo != nil {

		if swag.IsZero(m.ExtraInfo) { // not required
			return nil
		}

		if err := m.ExtraInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extra_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extra_info")
			}
			return err
		}
	}

	return nil
}

func (m *DomainHost) contextValidateVulnerabilities(ctx context.Context, formats strfmt.Registry) error {

	if m.Vulnerabilities != nil {

		if swag.IsZero(m.Vulnerabilities) { // not required
			return nil
		}

		if err := m.Vulnerabilities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vulnerabilities")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainHost) UnmarshalBinary(b []byte) error {
	var res DomainHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
