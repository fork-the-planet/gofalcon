// Code generated by go-swagger; DO NOT EDIT.

package container_images

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

// NewAggregateImageCountByStateParams creates a new AggregateImageCountByStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateImageCountByStateParams() *AggregateImageCountByStateParams {
	return &AggregateImageCountByStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateImageCountByStateParamsWithTimeout creates a new AggregateImageCountByStateParams object
// with the ability to set a timeout on a request.
func NewAggregateImageCountByStateParamsWithTimeout(timeout time.Duration) *AggregateImageCountByStateParams {
	return &AggregateImageCountByStateParams{
		timeout: timeout,
	}
}

// NewAggregateImageCountByStateParamsWithContext creates a new AggregateImageCountByStateParams object
// with the ability to set a context for a request.
func NewAggregateImageCountByStateParamsWithContext(ctx context.Context) *AggregateImageCountByStateParams {
	return &AggregateImageCountByStateParams{
		Context: ctx,
	}
}

// NewAggregateImageCountByStateParamsWithHTTPClient creates a new AggregateImageCountByStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateImageCountByStateParamsWithHTTPClient(client *http.Client) *AggregateImageCountByStateParams {
	return &AggregateImageCountByStateParams{
		HTTPClient: client,
	}
}

/*
AggregateImageCountByStateParams contains all the parameters to send to the API endpoint

	for the aggregate image count by state operation.

	Typically these are written to a http.Request.
*/
type AggregateImageCountByStateParams struct {

	/* Filter.

	     Filter images using a query in Falcon Query Language (FQL). Supported filter fields:
	- `arch`
	- `base_os`
	- `cid`
	- `first_seen`
	- `image_digest`
	- `image_id`
	- `registry`
	- `repository`
	- `source`
	- `tag`
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate image count by state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateImageCountByStateParams) WithDefaults() *AggregateImageCountByStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate image count by state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateImageCountByStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) WithTimeout(timeout time.Duration) *AggregateImageCountByStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) WithContext(ctx context.Context) *AggregateImageCountByStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) WithHTTPClient(client *http.Client) *AggregateImageCountByStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) WithFilter(filter *string) *AggregateImageCountByStateParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the aggregate image count by state params
func (o *AggregateImageCountByStateParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateImageCountByStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
