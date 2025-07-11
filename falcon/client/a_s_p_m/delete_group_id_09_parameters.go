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

// NewDeleteGroupID09Params creates a new DeleteGroupID09Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGroupID09Params() *DeleteGroupID09Params {
	return &DeleteGroupID09Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupID09ParamsWithTimeout creates a new DeleteGroupID09Params object
// with the ability to set a timeout on a request.
func NewDeleteGroupID09ParamsWithTimeout(timeout time.Duration) *DeleteGroupID09Params {
	return &DeleteGroupID09Params{
		timeout: timeout,
	}
}

// NewDeleteGroupID09ParamsWithContext creates a new DeleteGroupID09Params object
// with the ability to set a context for a request.
func NewDeleteGroupID09ParamsWithContext(ctx context.Context) *DeleteGroupID09Params {
	return &DeleteGroupID09Params{
		Context: ctx,
	}
}

// NewDeleteGroupID09ParamsWithHTTPClient creates a new DeleteGroupID09Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGroupID09ParamsWithHTTPClient(client *http.Client) *DeleteGroupID09Params {
	return &DeleteGroupID09Params{
		HTTPClient: client,
	}
}

/*
DeleteGroupID09Params contains all the parameters to send to the API endpoint

	for the delete group ID 0 9 operation.

	Typically these are written to a http.Request.
*/
type DeleteGroupID09Params struct {

	/* ID.

	   Group ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete group ID 0 9 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupID09Params) WithDefaults() *DeleteGroupID09Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete group ID 0 9 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupID09Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) WithTimeout(timeout time.Duration) *DeleteGroupID09Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) WithContext(ctx context.Context) *DeleteGroupID09Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) WithHTTPClient(client *http.Client) *DeleteGroupID09Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) WithID(id int64) *DeleteGroupID09Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete group ID 0 9 params
func (o *DeleteGroupID09Params) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupID09Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
