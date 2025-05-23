// Code generated by go-swagger; DO NOT EDIT.

package cloud_oci_registration

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

// NewCloudSecurityRegistrationOciCreateAccountParams creates a new CloudSecurityRegistrationOciCreateAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudSecurityRegistrationOciCreateAccountParams() *CloudSecurityRegistrationOciCreateAccountParams {
	return &CloudSecurityRegistrationOciCreateAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudSecurityRegistrationOciCreateAccountParamsWithTimeout creates a new CloudSecurityRegistrationOciCreateAccountParams object
// with the ability to set a timeout on a request.
func NewCloudSecurityRegistrationOciCreateAccountParamsWithTimeout(timeout time.Duration) *CloudSecurityRegistrationOciCreateAccountParams {
	return &CloudSecurityRegistrationOciCreateAccountParams{
		timeout: timeout,
	}
}

// NewCloudSecurityRegistrationOciCreateAccountParamsWithContext creates a new CloudSecurityRegistrationOciCreateAccountParams object
// with the ability to set a context for a request.
func NewCloudSecurityRegistrationOciCreateAccountParamsWithContext(ctx context.Context) *CloudSecurityRegistrationOciCreateAccountParams {
	return &CloudSecurityRegistrationOciCreateAccountParams{
		Context: ctx,
	}
}

// NewCloudSecurityRegistrationOciCreateAccountParamsWithHTTPClient creates a new CloudSecurityRegistrationOciCreateAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudSecurityRegistrationOciCreateAccountParamsWithHTTPClient(client *http.Client) *CloudSecurityRegistrationOciCreateAccountParams {
	return &CloudSecurityRegistrationOciCreateAccountParams{
		HTTPClient: client,
	}
}

/*
CloudSecurityRegistrationOciCreateAccountParams contains all the parameters to send to the API endpoint

	for the cloud security registration oci create account operation.

	Typically these are written to a http.Request.
*/
type CloudSecurityRegistrationOciCreateAccountParams struct {

	// Body.
	Body *models.DomainOCITenancyCreateRequestExtV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud security registration oci create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudSecurityRegistrationOciCreateAccountParams) WithDefaults() *CloudSecurityRegistrationOciCreateAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud security registration oci create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudSecurityRegistrationOciCreateAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) WithTimeout(timeout time.Duration) *CloudSecurityRegistrationOciCreateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) WithContext(ctx context.Context) *CloudSecurityRegistrationOciCreateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) WithHTTPClient(client *http.Client) *CloudSecurityRegistrationOciCreateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) WithBody(body *models.DomainOCITenancyCreateRequestExtV1) *CloudSecurityRegistrationOciCreateAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cloud security registration oci create account params
func (o *CloudSecurityRegistrationOciCreateAccountParams) SetBody(body *models.DomainOCITenancyCreateRequestExtV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CloudSecurityRegistrationOciCreateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
