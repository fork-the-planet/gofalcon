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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUpsertTagsParams creates a new UpsertTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpsertTagsParams() *UpsertTagsParams {
	return &UpsertTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpsertTagsParamsWithTimeout creates a new UpsertTagsParams object
// with the ability to set a timeout on a request.
func NewUpsertTagsParamsWithTimeout(timeout time.Duration) *UpsertTagsParams {
	return &UpsertTagsParams{
		timeout: timeout,
	}
}

// NewUpsertTagsParamsWithContext creates a new UpsertTagsParams object
// with the ability to set a context for a request.
func NewUpsertTagsParamsWithContext(ctx context.Context) *UpsertTagsParams {
	return &UpsertTagsParams{
		Context: ctx,
	}
}

// NewUpsertTagsParamsWithHTTPClient creates a new UpsertTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpsertTagsParamsWithHTTPClient(client *http.Client) *UpsertTagsParams {
	return &UpsertTagsParams{
		HTTPClient: client,
	}
}

/*
UpsertTagsParams contains all the parameters to send to the API endpoint

	for the upsert tags operation.

	Typically these are written to a http.Request.
*/
type UpsertTagsParams struct {

	// Body.
	Body *models.TypesEditUniqueTagRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upsert tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertTagsParams) WithDefaults() *UpsertTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upsert tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upsert tags params
func (o *UpsertTagsParams) WithTimeout(timeout time.Duration) *UpsertTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upsert tags params
func (o *UpsertTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upsert tags params
func (o *UpsertTagsParams) WithContext(ctx context.Context) *UpsertTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upsert tags params
func (o *UpsertTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upsert tags params
func (o *UpsertTagsParams) WithHTTPClient(client *http.Client) *UpsertTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upsert tags params
func (o *UpsertTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upsert tags params
func (o *UpsertTagsParams) WithBody(body *models.TypesEditUniqueTagRequest) *UpsertTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upsert tags params
func (o *UpsertTagsParams) SetBody(body *models.TypesEditUniqueTagRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
