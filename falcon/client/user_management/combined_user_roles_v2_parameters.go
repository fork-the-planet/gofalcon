// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// NewCombinedUserRolesV2Params creates a new CombinedUserRolesV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCombinedUserRolesV2Params() *CombinedUserRolesV2Params {
	return &CombinedUserRolesV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCombinedUserRolesV2ParamsWithTimeout creates a new CombinedUserRolesV2Params object
// with the ability to set a timeout on a request.
func NewCombinedUserRolesV2ParamsWithTimeout(timeout time.Duration) *CombinedUserRolesV2Params {
	return &CombinedUserRolesV2Params{
		timeout: timeout,
	}
}

// NewCombinedUserRolesV2ParamsWithContext creates a new CombinedUserRolesV2Params object
// with the ability to set a context for a request.
func NewCombinedUserRolesV2ParamsWithContext(ctx context.Context) *CombinedUserRolesV2Params {
	return &CombinedUserRolesV2Params{
		Context: ctx,
	}
}

// NewCombinedUserRolesV2ParamsWithHTTPClient creates a new CombinedUserRolesV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCombinedUserRolesV2ParamsWithHTTPClient(client *http.Client) *CombinedUserRolesV2Params {
	return &CombinedUserRolesV2Params{
		HTTPClient: client,
	}
}

/*
CombinedUserRolesV2Params contains all the parameters to send to the API endpoint

	for the combined user roles v2 operation.

	Typically these are written to a http.Request.
*/
type CombinedUserRolesV2Params struct {

	/* Cid.

	   Customer ID to get grants for. Empty CID would result in Role IDs for user against current CID in view.
	*/
	Cid *string

	/* DirectOnly.

	   Specifies if to request direct Only role grants or all role grants between user and CID (specified in query params)
	*/
	DirectOnly *bool

	/* Filter.

	   Filter using a query in Falcon Query Language (FQL). Supported filters: expires_at, role_id, role_name
	*/
	Filter *string

	/* Limit.

	   The maximum records to return. [1-500]

	   Default: 100
	*/
	Limit *int64

	/* Offset.

	   The offset to start retrieving records from
	*/
	Offset *int64

	/* Sort.

	   The property to sort by

	   Default: "role_name|asc"
	*/
	Sort *string

	/* UserUUID.

	   User UUID to get available roles for.
	*/
	UserUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the combined user roles v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedUserRolesV2Params) WithDefaults() *CombinedUserRolesV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the combined user roles v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedUserRolesV2Params) SetDefaults() {
	var (
		directOnlyDefault = bool(false)

		limitDefault = int64(100)

		offsetDefault = int64(0)

		sortDefault = string("role_name|asc")
	)

	val := CombinedUserRolesV2Params{
		DirectOnly: &directOnlyDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		Sort:       &sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithTimeout(timeout time.Duration) *CombinedUserRolesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithContext(ctx context.Context) *CombinedUserRolesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithHTTPClient(client *http.Client) *CombinedUserRolesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithCid(cid *string) *CombinedUserRolesV2Params {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetCid(cid *string) {
	o.Cid = cid
}

// WithDirectOnly adds the directOnly to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithDirectOnly(directOnly *bool) *CombinedUserRolesV2Params {
	o.SetDirectOnly(directOnly)
	return o
}

// SetDirectOnly adds the directOnly to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetDirectOnly(directOnly *bool) {
	o.DirectOnly = directOnly
}

// WithFilter adds the filter to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithFilter(filter *string) *CombinedUserRolesV2Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithLimit(limit *int64) *CombinedUserRolesV2Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithOffset(offset *int64) *CombinedUserRolesV2Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithSort(sort *string) *CombinedUserRolesV2Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetSort(sort *string) {
	o.Sort = sort
}

// WithUserUUID adds the userUUID to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) WithUserUUID(userUUID string) *CombinedUserRolesV2Params {
	o.SetUserUUID(userUUID)
	return o
}

// SetUserUUID adds the userUuid to the combined user roles v2 params
func (o *CombinedUserRolesV2Params) SetUserUUID(userUUID string) {
	o.UserUUID = userUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CombinedUserRolesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cid != nil {

		// query param cid
		var qrCid string

		if o.Cid != nil {
			qrCid = *o.Cid
		}
		qCid := qrCid
		if qCid != "" {

			if err := r.SetQueryParam("cid", qCid); err != nil {
				return err
			}
		}
	}

	if o.DirectOnly != nil {

		// query param direct_only
		var qrDirectOnly bool

		if o.DirectOnly != nil {
			qrDirectOnly = *o.DirectOnly
		}
		qDirectOnly := swag.FormatBool(qrDirectOnly)
		if qDirectOnly != "" {

			if err := r.SetQueryParam("direct_only", qDirectOnly); err != nil {
				return err
			}
		}
	}

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

	// query param user_uuid
	qrUserUUID := o.UserUUID
	qUserUUID := qrUserUUID
	if qUserUUID != "" {

		if err := r.SetQueryParam("user_uuid", qUserUUID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
