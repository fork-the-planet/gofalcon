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

// DomainActorNewsDocument domain actor news document
//
// swagger:model domain.ActorNewsDocument
type DomainActorNewsDocument struct {

	// Date of the news document creation, unix timestamp
	// Required: true
	CreatedDate *int64 `json:"created_date"`

	// Integer ID of the News document
	// Required: true
	ID *int64 `json:"id"`

	// News title
	// Required: true
	Name *string `json:"name"`

	// News title in a url friendly way, which is title in lowercase and special characters including space replaced with dash
	// Required: true
	Slug *string `json:"slug"`
}

// Validate validates this domain actor news document
func (m *DomainActorNewsDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainActorNewsDocument) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("created_date", "body", m.CreatedDate); err != nil {
		return err
	}

	return nil
}

func (m *DomainActorNewsDocument) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainActorNewsDocument) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainActorNewsDocument) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain actor news document based on context it is used
func (m *DomainActorNewsDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainActorNewsDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainActorNewsDocument) UnmarshalBinary(b []byte) error {
	var res DomainActorNewsDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}