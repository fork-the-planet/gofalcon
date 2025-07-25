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

// APIRuleCreateRequestV1 api rule create request v1
//
// swagger:model api.RuleCreateRequestV1
type APIRuleCreateRequestV1 struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// description
	Description string `json:"description,omitempty"`

	// mitre attack
	MitreAttack []*ModelMitreAttackMapping `json:"mitre_attack"`

	// name
	// Required: true
	Name *string `json:"name"`

	// notifications
	Notifications []*APICreateRuleNotifications `json:"notifications"`

	// operation
	// Required: true
	Operation *APICreateRuleOperationV1 `json:"operation"`

	// search
	// Required: true
	Search *APIRuleSearchV1 `json:"search"`

	// severity
	// Required: true
	Severity *int32 `json:"severity"`

	// status
	// Required: true
	Status *string `json:"status"`

	// tactic
	Tactic string `json:"tactic,omitempty"`

	// technique
	Technique string `json:"technique,omitempty"`

	// template id
	// Required: true
	TemplateID *string `json:"template_id"`

	// trigger on create
	TriggerOnCreate bool `json:"trigger_on_create,omitempty"`
}

// Validate validates this api rule create request v1
func (m *APIRuleCreateRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMitreAttack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRuleCreateRequestV1) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateMitreAttack(formats strfmt.Registry) error {
	if swag.IsZero(m.MitreAttack) { // not required
		return nil
	}

	for i := 0; i < len(m.MitreAttack); i++ {
		if swag.IsZero(m.MitreAttack[i]) { // not required
			continue
		}

		if m.MitreAttack[i] != nil {
			if err := m.MitreAttack[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitre_attack" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitre_attack" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	for i := 0; i < len(m.Notifications); i++ {
		if swag.IsZero(m.Notifications[i]) { // not required
			continue
		}

		if m.Notifications[i] != nil {
			if err := m.Notifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateSearch(formats strfmt.Registry) error {

	if err := validate.Required("search", "body", m.Search); err != nil {
		return err
	}

	if m.Search != nil {
		if err := m.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleCreateRequestV1) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("template_id", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api rule create request v1 based on the context it is used
func (m *APIRuleCreateRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMitreAttack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRuleCreateRequestV1) contextValidateMitreAttack(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MitreAttack); i++ {

		if m.MitreAttack[i] != nil {

			if swag.IsZero(m.MitreAttack[i]) { // not required
				return nil
			}

			if err := m.MitreAttack[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitre_attack" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitre_attack" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleCreateRequestV1) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notifications); i++ {

		if m.Notifications[i] != nil {

			if swag.IsZero(m.Notifications[i]) { // not required
				return nil
			}

			if err := m.Notifications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleCreateRequestV1) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {

		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *APIRuleCreateRequestV1) contextValidateSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.Search != nil {

		if err := m.Search.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIRuleCreateRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRuleCreateRequestV1) UnmarshalBinary(b []byte) error {
	var res APIRuleCreateRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
