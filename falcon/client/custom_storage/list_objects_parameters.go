// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// NewListObjectsParams creates a new ListObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListObjectsParams() *ListObjectsParams {
	return &ListObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListObjectsParamsWithTimeout creates a new ListObjectsParams object
// with the ability to set a timeout on a request.
func NewListObjectsParamsWithTimeout(timeout time.Duration) *ListObjectsParams {
	return &ListObjectsParams{
		timeout: timeout,
	}
}

// NewListObjectsParamsWithContext creates a new ListObjectsParams object
// with the ability to set a context for a request.
func NewListObjectsParamsWithContext(ctx context.Context) *ListObjectsParams {
	return &ListObjectsParams{
		Context: ctx,
	}
}

// NewListObjectsParamsWithHTTPClient creates a new ListObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListObjectsParamsWithHTTPClient(client *http.Client) *ListObjectsParams {
	return &ListObjectsParams{
		HTTPClient: client,
	}
}

/*
ListObjectsParams contains all the parameters to send to the API endpoint

	for the list objects operation.

	Typically these are written to a http.Request.
*/
type ListObjectsParams struct {

	/* CollectionName.

	   The name of the collection
	*/
	CollectionName string

	/* End.

	   The end key to end listing to
	*/
	End string

	/* Limit.

	   The limit of results to return
	*/
	Limit int64

	/* Start.

	   The start key to start listing from
	*/
	Start string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListObjectsParams) WithDefaults() *ListObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list objects params
func (o *ListObjectsParams) WithTimeout(timeout time.Duration) *ListObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list objects params
func (o *ListObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list objects params
func (o *ListObjectsParams) WithContext(ctx context.Context) *ListObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list objects params
func (o *ListObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list objects params
func (o *ListObjectsParams) WithHTTPClient(client *http.Client) *ListObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list objects params
func (o *ListObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionName adds the collectionName to the list objects params
func (o *ListObjectsParams) WithCollectionName(collectionName string) *ListObjectsParams {
	o.SetCollectionName(collectionName)
	return o
}

// SetCollectionName adds the collectionName to the list objects params
func (o *ListObjectsParams) SetCollectionName(collectionName string) {
	o.CollectionName = collectionName
}

// WithEnd adds the end to the list objects params
func (o *ListObjectsParams) WithEnd(end string) *ListObjectsParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the list objects params
func (o *ListObjectsParams) SetEnd(end string) {
	o.End = end
}

// WithLimit adds the limit to the list objects params
func (o *ListObjectsParams) WithLimit(limit int64) *ListObjectsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list objects params
func (o *ListObjectsParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithStart adds the start to the list objects params
func (o *ListObjectsParams) WithStart(start string) *ListObjectsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list objects params
func (o *ListObjectsParams) SetStart(start string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection_name
	if err := r.SetPathParam("collection_name", o.CollectionName); err != nil {
		return err
	}

	// query param end
	qrEnd := o.End
	qEnd := qrEnd

	if err := r.SetQueryParam("end", qEnd); err != nil {
		return err
	}

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)

	if err := r.SetQueryParam("limit", qLimit); err != nil {
		return err
	}

	// query param start
	qrStart := o.Start
	qStart := qrStart

	if err := r.SetQueryParam("start", qStart); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
