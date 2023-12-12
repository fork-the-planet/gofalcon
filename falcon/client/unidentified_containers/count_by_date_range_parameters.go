// Code generated by go-swagger; DO NOT EDIT.

package unidentified_containers

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
)

// NewCountByDateRangeParams creates a new CountByDateRangeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCountByDateRangeParams() *CountByDateRangeParams {
	return &CountByDateRangeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCountByDateRangeParamsWithTimeout creates a new CountByDateRangeParams object
// with the ability to set a timeout on a request.
func NewCountByDateRangeParamsWithTimeout(timeout time.Duration) *CountByDateRangeParams {
	return &CountByDateRangeParams{
		timeout: timeout,
	}
}

// NewCountByDateRangeParamsWithContext creates a new CountByDateRangeParams object
// with the ability to set a context for a request.
func NewCountByDateRangeParamsWithContext(ctx context.Context) *CountByDateRangeParams {
	return &CountByDateRangeParams{
		Context: ctx,
	}
}

// NewCountByDateRangeParamsWithHTTPClient creates a new CountByDateRangeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCountByDateRangeParamsWithHTTPClient(client *http.Client) *CountByDateRangeParams {
	return &CountByDateRangeParams{
		HTTPClient: client,
	}
}

/*
CountByDateRangeParams contains all the parameters to send to the API endpoint

	for the count by date range operation.

	Typically these are written to a http.Request.
*/
type CountByDateRangeParams struct {

	/* Filter.

	   Filter Unidentified Containers using a query in Falcon Query Language (FQL). Supported filters:  assessed_images_count,cid,cluster_name,containers_impacted_count,detections_count,image_assessment_detections_count,last_seen,namespace,node_name,severity,unassessed_images_count,visible_to_k8s
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the count by date range params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CountByDateRangeParams) WithDefaults() *CountByDateRangeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the count by date range params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CountByDateRangeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the count by date range params
func (o *CountByDateRangeParams) WithTimeout(timeout time.Duration) *CountByDateRangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the count by date range params
func (o *CountByDateRangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the count by date range params
func (o *CountByDateRangeParams) WithContext(ctx context.Context) *CountByDateRangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the count by date range params
func (o *CountByDateRangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the count by date range params
func (o *CountByDateRangeParams) WithHTTPClient(client *http.Client) *CountByDateRangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the count by date range params
func (o *CountByDateRangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the count by date range params
func (o *CountByDateRangeParams) WithFilter(filter *string) *CountByDateRangeParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the count by date range params
func (o *CountByDateRangeParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *CountByDateRangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}