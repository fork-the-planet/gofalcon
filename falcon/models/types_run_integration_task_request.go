// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesRunIntegrationTaskRequest types run integration task request
//
// swagger:model types.RunIntegrationTaskRequest
type TypesRunIntegrationTaskRequest struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// override
	Override bool `json:"override,omitempty"`

	// scheduled
	Scheduled bool `json:"scheduled,omitempty"`

	// task id
	TaskID int64 `json:"task_id,omitempty"`
}

// Validate validates this types run integration task request
func (m *TypesRunIntegrationTaskRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this types run integration task request based on context it is used
func (m *TypesRunIntegrationTaskRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesRunIntegrationTaskRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesRunIntegrationTaskRequest) UnmarshalBinary(b []byte) error {
	var res TypesRunIntegrationTaskRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
