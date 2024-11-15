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

// NewGetExecutorNodesParams creates a new GetExecutorNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExecutorNodesParams() *GetExecutorNodesParams {
	return &GetExecutorNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExecutorNodesParamsWithTimeout creates a new GetExecutorNodesParams object
// with the ability to set a timeout on a request.
func NewGetExecutorNodesParamsWithTimeout(timeout time.Duration) *GetExecutorNodesParams {
	return &GetExecutorNodesParams{
		timeout: timeout,
	}
}

// NewGetExecutorNodesParamsWithContext creates a new GetExecutorNodesParams object
// with the ability to set a context for a request.
func NewGetExecutorNodesParamsWithContext(ctx context.Context) *GetExecutorNodesParams {
	return &GetExecutorNodesParams{
		Context: ctx,
	}
}

// NewGetExecutorNodesParamsWithHTTPClient creates a new GetExecutorNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExecutorNodesParamsWithHTTPClient(client *http.Client) *GetExecutorNodesParams {
	return &GetExecutorNodesParams{
		HTTPClient: client,
	}
}

/*
GetExecutorNodesParams contains all the parameters to send to the API endpoint

	for the get executor nodes operation.

	Typically these are written to a http.Request.
*/
type GetExecutorNodesParams struct {

	// IntegrationType.
	IntegrationType *int64

	// NodeType.
	NodeType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get executor nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExecutorNodesParams) WithDefaults() *GetExecutorNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get executor nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExecutorNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get executor nodes params
func (o *GetExecutorNodesParams) WithTimeout(timeout time.Duration) *GetExecutorNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get executor nodes params
func (o *GetExecutorNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get executor nodes params
func (o *GetExecutorNodesParams) WithContext(ctx context.Context) *GetExecutorNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get executor nodes params
func (o *GetExecutorNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get executor nodes params
func (o *GetExecutorNodesParams) WithHTTPClient(client *http.Client) *GetExecutorNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get executor nodes params
func (o *GetExecutorNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationType adds the integrationType to the get executor nodes params
func (o *GetExecutorNodesParams) WithIntegrationType(integrationType *int64) *GetExecutorNodesParams {
	o.SetIntegrationType(integrationType)
	return o
}

// SetIntegrationType adds the integrationType to the get executor nodes params
func (o *GetExecutorNodesParams) SetIntegrationType(integrationType *int64) {
	o.IntegrationType = integrationType
}

// WithNodeType adds the nodeType to the get executor nodes params
func (o *GetExecutorNodesParams) WithNodeType(nodeType string) *GetExecutorNodesParams {
	o.SetNodeType(nodeType)
	return o
}

// SetNodeType adds the nodeType to the get executor nodes params
func (o *GetExecutorNodesParams) SetNodeType(nodeType string) {
	o.NodeType = nodeType
}

// WriteToRequest writes these params to a swagger request
func (o *GetExecutorNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IntegrationType != nil {

		// query param integration_type
		var qrIntegrationType int64

		if o.IntegrationType != nil {
			qrIntegrationType = *o.IntegrationType
		}
		qIntegrationType := swag.FormatInt64(qrIntegrationType)
		if qIntegrationType != "" {

			if err := r.SetQueryParam("integration_type", qIntegrationType); err != nil {
				return err
			}
		}
	}

	// query param node_type
	qrNodeType := o.NodeType
	qNodeType := qrNodeType
	if qNodeType != "" {

		if err := r.SetQueryParam("node_type", qNodeType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
