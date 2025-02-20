// Code generated by go-swagger; DO NOT EDIT.

package delivery_settings

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

// NewPostDeliverySettingsParams creates a new PostDeliverySettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDeliverySettingsParams() *PostDeliverySettingsParams {
	return &PostDeliverySettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDeliverySettingsParamsWithTimeout creates a new PostDeliverySettingsParams object
// with the ability to set a timeout on a request.
func NewPostDeliverySettingsParamsWithTimeout(timeout time.Duration) *PostDeliverySettingsParams {
	return &PostDeliverySettingsParams{
		timeout: timeout,
	}
}

// NewPostDeliverySettingsParamsWithContext creates a new PostDeliverySettingsParams object
// with the ability to set a context for a request.
func NewPostDeliverySettingsParamsWithContext(ctx context.Context) *PostDeliverySettingsParams {
	return &PostDeliverySettingsParams{
		Context: ctx,
	}
}

// NewPostDeliverySettingsParamsWithHTTPClient creates a new PostDeliverySettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDeliverySettingsParamsWithHTTPClient(client *http.Client) *PostDeliverySettingsParams {
	return &PostDeliverySettingsParams{
		HTTPClient: client,
	}
}

/*
PostDeliverySettingsParams contains all the parameters to send to the API endpoint

	for the post delivery settings operation.

	Typically these are written to a http.Request.
*/
type PostDeliverySettingsParams struct {

	// Body.
	Body *models.ModelsDeliverySettingsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post delivery settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDeliverySettingsParams) WithDefaults() *PostDeliverySettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post delivery settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDeliverySettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post delivery settings params
func (o *PostDeliverySettingsParams) WithTimeout(timeout time.Duration) *PostDeliverySettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post delivery settings params
func (o *PostDeliverySettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post delivery settings params
func (o *PostDeliverySettingsParams) WithContext(ctx context.Context) *PostDeliverySettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post delivery settings params
func (o *PostDeliverySettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post delivery settings params
func (o *PostDeliverySettingsParams) WithHTTPClient(client *http.Client) *PostDeliverySettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post delivery settings params
func (o *PostDeliverySettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post delivery settings params
func (o *PostDeliverySettingsParams) WithBody(body *models.ModelsDeliverySettingsRequest) *PostDeliverySettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post delivery settings params
func (o *PostDeliverySettingsParams) SetBody(body *models.ModelsDeliverySettingsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDeliverySettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
