// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// NewQueryNotificationsExposedDataRecordsV1Params creates a new QueryNotificationsExposedDataRecordsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryNotificationsExposedDataRecordsV1Params() *QueryNotificationsExposedDataRecordsV1Params {
	return &QueryNotificationsExposedDataRecordsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryNotificationsExposedDataRecordsV1ParamsWithTimeout creates a new QueryNotificationsExposedDataRecordsV1Params object
// with the ability to set a timeout on a request.
func NewQueryNotificationsExposedDataRecordsV1ParamsWithTimeout(timeout time.Duration) *QueryNotificationsExposedDataRecordsV1Params {
	return &QueryNotificationsExposedDataRecordsV1Params{
		timeout: timeout,
	}
}

// NewQueryNotificationsExposedDataRecordsV1ParamsWithContext creates a new QueryNotificationsExposedDataRecordsV1Params object
// with the ability to set a context for a request.
func NewQueryNotificationsExposedDataRecordsV1ParamsWithContext(ctx context.Context) *QueryNotificationsExposedDataRecordsV1Params {
	return &QueryNotificationsExposedDataRecordsV1Params{
		Context: ctx,
	}
}

// NewQueryNotificationsExposedDataRecordsV1ParamsWithHTTPClient creates a new QueryNotificationsExposedDataRecordsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewQueryNotificationsExposedDataRecordsV1ParamsWithHTTPClient(client *http.Client) *QueryNotificationsExposedDataRecordsV1Params {
	return &QueryNotificationsExposedDataRecordsV1Params{
		HTTPClient: client,
	}
}

/*
QueryNotificationsExposedDataRecordsV1Params contains all the parameters to send to the API endpoint

	for the query notifications exposed data records v1 operation.

	Typically these are written to a http.Request.
*/
type QueryNotificationsExposedDataRecordsV1Params struct {

	/* Filter.

	   FQL query to filter notifications by. Possible filter properties are: [id cid user_uuid created_date exposure_date rule.id rule.name rule.topic notification_id source_category site site_id author author_id user_id user_name impacted_url impacted_domain impacted_ip email email_domain hash_type display_name full_name user_ip phone_number company job_position file.name file.complete_data_set file.download_urls location.postal_code location.city location.state location.federal_district location.federal_admin_region location.country_code social.twitter_id social.facebook_id social.vk_id social.vk_token social.aim_id social.icq_id social.msn_id social.instagram_id social.skype_id financial.credit_card financial.bank_account financial.crypto_currency_addresses login_id _all]
	*/
	Filter *string

	/* Limit.

	   Number of IDs to return. Offset + limit should NOT be above 10K.
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return ids.
	*/
	Offset *int64

	/* Q.

	   Free text search across all indexed fields.
	*/
	Q *string

	/* Sort.

	   Possible order by fields: created_date, updated_date. Ex: 'updated_date|desc'.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query notifications exposed data records v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryNotificationsExposedDataRecordsV1Params) WithDefaults() *QueryNotificationsExposedDataRecordsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query notifications exposed data records v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryNotificationsExposedDataRecordsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithTimeout(timeout time.Duration) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithContext(ctx context.Context) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithHTTPClient(client *http.Client) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithFilter(filter *string) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithLimit(limit *int64) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithOffset(offset *int64) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithQ(q *string) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) WithSort(sort *string) *QueryNotificationsExposedDataRecordsV1Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query notifications exposed data records v1 params
func (o *QueryNotificationsExposedDataRecordsV1Params) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryNotificationsExposedDataRecordsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
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