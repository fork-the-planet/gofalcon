// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// NewClustersByStatusCountParams creates a new ClustersByStatusCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClustersByStatusCountParams() *ClustersByStatusCountParams {
	return &ClustersByStatusCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClustersByStatusCountParamsWithTimeout creates a new ClustersByStatusCountParams object
// with the ability to set a timeout on a request.
func NewClustersByStatusCountParamsWithTimeout(timeout time.Duration) *ClustersByStatusCountParams {
	return &ClustersByStatusCountParams{
		timeout: timeout,
	}
}

// NewClustersByStatusCountParamsWithContext creates a new ClustersByStatusCountParams object
// with the ability to set a context for a request.
func NewClustersByStatusCountParamsWithContext(ctx context.Context) *ClustersByStatusCountParams {
	return &ClustersByStatusCountParams{
		Context: ctx,
	}
}

// NewClustersByStatusCountParamsWithHTTPClient creates a new ClustersByStatusCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewClustersByStatusCountParamsWithHTTPClient(client *http.Client) *ClustersByStatusCountParams {
	return &ClustersByStatusCountParams{
		HTTPClient: client,
	}
}

/*
ClustersByStatusCountParams contains all the parameters to send to the API endpoint

	for the clusters by status count operation.

	Typically these are written to a http.Request.
*/
type ClustersByStatusCountParams struct {

	/* Filter.

	     Retrieve count of Kubernetes clusters that match a query in Falcon Query Language (FQL). Supported filter fields:
	- `access`
	- `agent_id`
	- `agent_status`
	- `agent_type`
	- `cid`
	- `cloud_account_id`
	- `cloud_name`
	- `cloud_region`
	- `cloud_service`
	- `cluster_id`
	- `cluster_name`
	- `cluster_status`
	- `container_count`
	- `iar_coverage`
	- `kac_agent_id`
	- `kubernetes_version`
	- `last_seen`
	- `management_status`
	- `node_count`
	- `pod_count`
	- `tags`
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clusters by status count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClustersByStatusCountParams) WithDefaults() *ClustersByStatusCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clusters by status count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClustersByStatusCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clusters by status count params
func (o *ClustersByStatusCountParams) WithTimeout(timeout time.Duration) *ClustersByStatusCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clusters by status count params
func (o *ClustersByStatusCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clusters by status count params
func (o *ClustersByStatusCountParams) WithContext(ctx context.Context) *ClustersByStatusCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clusters by status count params
func (o *ClustersByStatusCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clusters by status count params
func (o *ClustersByStatusCountParams) WithHTTPClient(client *http.Client) *ClustersByStatusCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clusters by status count params
func (o *ClustersByStatusCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the clusters by status count params
func (o *ClustersByStatusCountParams) WithFilter(filter *string) *ClustersByStatusCountParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the clusters by status count params
func (o *ClustersByStatusCountParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *ClustersByStatusCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
