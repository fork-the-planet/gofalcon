// Code generated by go-swagger; DO NOT EDIT.

package a_s_p_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetIntegrationTasksAdminParams creates a new GetIntegrationTasksAdminParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationTasksAdminParams() *GetIntegrationTasksAdminParams {
	return &GetIntegrationTasksAdminParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationTasksAdminParamsWithTimeout creates a new GetIntegrationTasksAdminParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationTasksAdminParamsWithTimeout(timeout time.Duration) *GetIntegrationTasksAdminParams {
	return &GetIntegrationTasksAdminParams{
		timeout: timeout,
	}
}

// NewGetIntegrationTasksAdminParamsWithContext creates a new GetIntegrationTasksAdminParams object
// with the ability to set a context for a request.
func NewGetIntegrationTasksAdminParamsWithContext(ctx context.Context) *GetIntegrationTasksAdminParams {
	return &GetIntegrationTasksAdminParams{
		Context: ctx,
	}
}

// NewGetIntegrationTasksAdminParamsWithHTTPClient creates a new GetIntegrationTasksAdminParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationTasksAdminParamsWithHTTPClient(client *http.Client) *GetIntegrationTasksAdminParams {
	return &GetIntegrationTasksAdminParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationTasksAdminParams contains all the parameters to send to the API endpoint

	for the get integration tasks admin operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationTasksAdminParams struct {

	// Category.
	Category *string

	// Direction.
	Direction *string

	// Ids.
	Ids *int64

	// IntegrationTaskType.
	IntegrationTaskType *int64

	// IntegrationTaskTypes.
	IntegrationTaskTypes *int64

	// Limit.
	Limit *int64

	// Names.
	Names *string

	// Offset.
	Offset *int64

	// OrderBy.
	OrderBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integration tasks admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationTasksAdminParams) WithDefaults() *GetIntegrationTasksAdminParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integration tasks admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationTasksAdminParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithTimeout(timeout time.Duration) *GetIntegrationTasksAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithContext(ctx context.Context) *GetIntegrationTasksAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithHTTPClient(client *http.Client) *GetIntegrationTasksAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithCategory(category *string) *GetIntegrationTasksAdminParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetCategory(category *string) {
	o.Category = category
}

// WithDirection adds the direction to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithDirection(direction *string) *GetIntegrationTasksAdminParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithIds adds the ids to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithIds(ids *int64) *GetIntegrationTasksAdminParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetIds(ids *int64) {
	o.Ids = ids
}

// WithIntegrationTaskType adds the integrationTaskType to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithIntegrationTaskType(integrationTaskType *int64) *GetIntegrationTasksAdminParams {
	o.SetIntegrationTaskType(integrationTaskType)
	return o
}

// SetIntegrationTaskType adds the integrationTaskType to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetIntegrationTaskType(integrationTaskType *int64) {
	o.IntegrationTaskType = integrationTaskType
}

// WithIntegrationTaskTypes adds the integrationTaskTypes to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithIntegrationTaskTypes(integrationTaskTypes *int64) *GetIntegrationTasksAdminParams {
	o.SetIntegrationTaskTypes(integrationTaskTypes)
	return o
}

// SetIntegrationTaskTypes adds the integrationTaskTypes to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetIntegrationTaskTypes(integrationTaskTypes *int64) {
	o.IntegrationTaskTypes = integrationTaskTypes
}

// WithLimit adds the limit to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithLimit(limit *int64) *GetIntegrationTasksAdminParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNames adds the names to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithNames(names *string) *GetIntegrationTasksAdminParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetNames(names *string) {
	o.Names = names
}

// WithOffset adds the offset to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithOffset(offset *int64) *GetIntegrationTasksAdminParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) WithOrderBy(orderBy *string) *GetIntegrationTasksAdminParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get integration tasks admin params
func (o *GetIntegrationTasksAdminParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationTasksAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// query param category
		var qrCategory string

		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {

			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string

		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {

			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}
	}

	if o.Ids != nil {

		// query param ids
		var qrIds int64

		if o.Ids != nil {
			qrIds = *o.Ids
		}
		qIds := swag.FormatInt64(qrIds)
		if qIds != "" {

			if err := r.SetQueryParam("ids", qIds); err != nil {
				return err
			}
		}
	}

	if o.IntegrationTaskType != nil {

		// query param integration_task_type
		var qrIntegrationTaskType int64

		if o.IntegrationTaskType != nil {
			qrIntegrationTaskType = *o.IntegrationTaskType
		}
		qIntegrationTaskType := swag.FormatInt64(qrIntegrationTaskType)
		if qIntegrationTaskType != "" {

			if err := r.SetQueryParam("integration_task_type", qIntegrationTaskType); err != nil {
				return err
			}
		}
	}

	if o.IntegrationTaskTypes != nil {

		// query param integration_task_types
		var qrIntegrationTaskTypes int64

		if o.IntegrationTaskTypes != nil {
			qrIntegrationTaskTypes = *o.IntegrationTaskTypes
		}
		qIntegrationTaskTypes := swag.FormatInt64(qrIntegrationTaskTypes)
		if qIntegrationTaskTypes != "" {

			if err := r.SetQueryParam("integration_task_types", qIntegrationTaskTypes); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Names != nil {

		// query param names
		var qrNames string

		if o.Names != nil {
			qrNames = *o.Names
		}
		qNames := qrNames
		if qNames != "" {

			if err := r.SetQueryParam("names", qNames); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
