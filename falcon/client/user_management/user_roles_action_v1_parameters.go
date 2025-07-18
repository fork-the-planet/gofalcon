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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUserRolesActionV1Params creates a new UserRolesActionV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserRolesActionV1Params() *UserRolesActionV1Params {
	return &UserRolesActionV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserRolesActionV1ParamsWithTimeout creates a new UserRolesActionV1Params object
// with the ability to set a timeout on a request.
func NewUserRolesActionV1ParamsWithTimeout(timeout time.Duration) *UserRolesActionV1Params {
	return &UserRolesActionV1Params{
		timeout: timeout,
	}
}

// NewUserRolesActionV1ParamsWithContext creates a new UserRolesActionV1Params object
// with the ability to set a context for a request.
func NewUserRolesActionV1ParamsWithContext(ctx context.Context) *UserRolesActionV1Params {
	return &UserRolesActionV1Params{
		Context: ctx,
	}
}

// NewUserRolesActionV1ParamsWithHTTPClient creates a new UserRolesActionV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUserRolesActionV1ParamsWithHTTPClient(client *http.Client) *UserRolesActionV1Params {
	return &UserRolesActionV1Params{
		HTTPClient: client,
	}
}

/*
UserRolesActionV1Params contains all the parameters to send to the API endpoint

	for the user roles action v1 operation.

	Typically these are written to a http.Request.
*/
type UserRolesActionV1Params struct {

	/* Body.

	   CID, RoleID(s), User UUID and Action are required. Allowed values for Action param include 'grant' and 'revoke'.
	*/
	Body *models.FlightcontrolapiGrantInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user roles action v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserRolesActionV1Params) WithDefaults() *UserRolesActionV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user roles action v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserRolesActionV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user roles action v1 params
func (o *UserRolesActionV1Params) WithTimeout(timeout time.Duration) *UserRolesActionV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user roles action v1 params
func (o *UserRolesActionV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user roles action v1 params
func (o *UserRolesActionV1Params) WithContext(ctx context.Context) *UserRolesActionV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user roles action v1 params
func (o *UserRolesActionV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user roles action v1 params
func (o *UserRolesActionV1Params) WithHTTPClient(client *http.Client) *UserRolesActionV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user roles action v1 params
func (o *UserRolesActionV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user roles action v1 params
func (o *UserRolesActionV1Params) WithBody(body *models.FlightcontrolapiGrantInput) *UserRolesActionV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user roles action v1 params
func (o *UserRolesActionV1Params) SetBody(body *models.FlightcontrolapiGrantInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UserRolesActionV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
