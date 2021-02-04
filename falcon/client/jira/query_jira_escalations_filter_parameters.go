// Code generated by go-swagger; DO NOT EDIT.

package jira

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

// NewQueryJiraEscalationsFilterParams creates a new QueryJiraEscalationsFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryJiraEscalationsFilterParams() *QueryJiraEscalationsFilterParams {
	return &QueryJiraEscalationsFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryJiraEscalationsFilterParamsWithTimeout creates a new QueryJiraEscalationsFilterParams object
// with the ability to set a timeout on a request.
func NewQueryJiraEscalationsFilterParamsWithTimeout(timeout time.Duration) *QueryJiraEscalationsFilterParams {
	return &QueryJiraEscalationsFilterParams{
		timeout: timeout,
	}
}

// NewQueryJiraEscalationsFilterParamsWithContext creates a new QueryJiraEscalationsFilterParams object
// with the ability to set a context for a request.
func NewQueryJiraEscalationsFilterParamsWithContext(ctx context.Context) *QueryJiraEscalationsFilterParams {
	return &QueryJiraEscalationsFilterParams{
		Context: ctx,
	}
}

// NewQueryJiraEscalationsFilterParamsWithHTTPClient creates a new QueryJiraEscalationsFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryJiraEscalationsFilterParamsWithHTTPClient(client *http.Client) *QueryJiraEscalationsFilterParams {
	return &QueryJiraEscalationsFilterParams{
		HTTPClient: client,
	}
}

/* QueryJiraEscalationsFilterParams contains all the parameters to send to the API endpoint
   for the query jira escalations filter operation.

   Typically these are written to a http.Request.
*/
type QueryJiraEscalationsFilterParams struct {

	/* Filter.

	   Optional filter and sort criteria in the form of an FQL query. For more information about FQL queries, see [our FQL documentation in Falcon](https://falcon.crowdstrike.com/support/documentation/45/falcon-query-language-feature-guide).
	*/
	Filter *string

	/* Limit.

	   The maximum records to return. [1-500]
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return ids.
	*/
	Offset *string

	/* Sort.

	   The property to sort on, followed by a dot (.), followed by the sort direction, either "asc" or "desc".
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query jira escalations filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryJiraEscalationsFilterParams) WithDefaults() *QueryJiraEscalationsFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query jira escalations filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryJiraEscalationsFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) WithTimeout(timeout time.Duration) *QueryJiraEscalationsFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) WithContext(ctx context.Context) *QueryJiraEscalationsFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) WithHTTPClient(client *http.Client) *QueryJiraEscalationsFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) WithFilter(filter *string) *QueryJiraEscalationsFilterParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) WithLimit(limit *int64) *QueryJiraEscalationsFilterParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) WithOffset(offset *string) *QueryJiraEscalationsFilterParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) WithSort(sort *string) *QueryJiraEscalationsFilterParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query jira escalations filter params
func (o *QueryJiraEscalationsFilterParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryJiraEscalationsFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
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

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
