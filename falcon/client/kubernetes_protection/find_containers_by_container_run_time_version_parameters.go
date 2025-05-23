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
	"github.com/go-openapi/swag"
)

// NewFindContainersByContainerRunTimeVersionParams creates a new FindContainersByContainerRunTimeVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindContainersByContainerRunTimeVersionParams() *FindContainersByContainerRunTimeVersionParams {
	return &FindContainersByContainerRunTimeVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindContainersByContainerRunTimeVersionParamsWithTimeout creates a new FindContainersByContainerRunTimeVersionParams object
// with the ability to set a timeout on a request.
func NewFindContainersByContainerRunTimeVersionParamsWithTimeout(timeout time.Duration) *FindContainersByContainerRunTimeVersionParams {
	return &FindContainersByContainerRunTimeVersionParams{
		timeout: timeout,
	}
}

// NewFindContainersByContainerRunTimeVersionParamsWithContext creates a new FindContainersByContainerRunTimeVersionParams object
// with the ability to set a context for a request.
func NewFindContainersByContainerRunTimeVersionParamsWithContext(ctx context.Context) *FindContainersByContainerRunTimeVersionParams {
	return &FindContainersByContainerRunTimeVersionParams{
		Context: ctx,
	}
}

// NewFindContainersByContainerRunTimeVersionParamsWithHTTPClient creates a new FindContainersByContainerRunTimeVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindContainersByContainerRunTimeVersionParamsWithHTTPClient(client *http.Client) *FindContainersByContainerRunTimeVersionParams {
	return &FindContainersByContainerRunTimeVersionParams{
		HTTPClient: client,
	}
}

/*
FindContainersByContainerRunTimeVersionParams contains all the parameters to send to the API endpoint

	for the find containers by container run time version operation.

	Typically these are written to a http.Request.
*/
type FindContainersByContainerRunTimeVersionParams struct {

	/* Filter.

	     Retrieve count of Kubernetes containers that match a query in Falcon Query Language (FQL). Supported filter fields:
	- `agent_id`
	- `agent_type`
	- `ai_related`
	- `allow_privilege_escalation`
	- `app_name`
	- `cid`
	- `cloud_account_id`
	- `cloud_instance_id`
	- `cloud_name`
	- `cloud_region`
	- `cloud_service`
	- `cluster_id`
	- `cluster_name`
	- `container_id`
	- `container_image_id`
	- `container_name`
	- `cve_id`
	- `detection_name`
	- `first_seen`
	- `image_detection_count`
	- `image_digest`
	- `image_has_been_assessed`
	- `image_id`
	- `image_registry`
	- `image_repository`
	- `image_tag`
	- `image_vulnerability_count`
	- `insecure_mount_source`
	- `insecure_mount_type`
	- `insecure_propagation_mode`
	- `interactive_mode`
	- `ipv4`
	- `ipv6`
	- `kac_agent_id`
	- `labels`
	- `last_seen`
	- `namespace`
	- `node_name`
	- `node_uid`
	- `package_name_version`
	- `pod_id`
	- `pod_name`
	- `port`
	- `privileged`
	- `root_write_access`
	- `run_as_root_group`
	- `run_as_root_user`
	- `running_status`
	*/
	Filter *string

	/* Limit.

	   The upper-bound on the number of records to retrieve.

	   Default: 200
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* Sort.

	   The fields to sort the records on.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find containers by container run time version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindContainersByContainerRunTimeVersionParams) WithDefaults() *FindContainersByContainerRunTimeVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find containers by container run time version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindContainersByContainerRunTimeVersionParams) SetDefaults() {
	var (
		limitDefault = int64(200)
	)

	val := FindContainersByContainerRunTimeVersionParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) WithTimeout(timeout time.Duration) *FindContainersByContainerRunTimeVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) WithContext(ctx context.Context) *FindContainersByContainerRunTimeVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) WithHTTPClient(client *http.Client) *FindContainersByContainerRunTimeVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) WithFilter(filter *string) *FindContainersByContainerRunTimeVersionParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) WithLimit(limit *int64) *FindContainersByContainerRunTimeVersionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) WithOffset(offset *int64) *FindContainersByContainerRunTimeVersionParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) WithSort(sort *string) *FindContainersByContainerRunTimeVersionParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the find containers by container run time version params
func (o *FindContainersByContainerRunTimeVersionParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *FindContainersByContainerRunTimeVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
