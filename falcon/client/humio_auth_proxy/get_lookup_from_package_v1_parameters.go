// Code generated by go-swagger; DO NOT EDIT.

package humio_auth_proxy

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
)

// NewGetLookupFromPackageV1Params creates a new GetLookupFromPackageV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLookupFromPackageV1Params() *GetLookupFromPackageV1Params {
	return &GetLookupFromPackageV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLookupFromPackageV1ParamsWithTimeout creates a new GetLookupFromPackageV1Params object
// with the ability to set a timeout on a request.
func NewGetLookupFromPackageV1ParamsWithTimeout(timeout time.Duration) *GetLookupFromPackageV1Params {
	return &GetLookupFromPackageV1Params{
		timeout: timeout,
	}
}

// NewGetLookupFromPackageV1ParamsWithContext creates a new GetLookupFromPackageV1Params object
// with the ability to set a context for a request.
func NewGetLookupFromPackageV1ParamsWithContext(ctx context.Context) *GetLookupFromPackageV1Params {
	return &GetLookupFromPackageV1Params{
		Context: ctx,
	}
}

// NewGetLookupFromPackageV1ParamsWithHTTPClient creates a new GetLookupFromPackageV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetLookupFromPackageV1ParamsWithHTTPClient(client *http.Client) *GetLookupFromPackageV1Params {
	return &GetLookupFromPackageV1Params{
		HTTPClient: client,
	}
}

/*
GetLookupFromPackageV1Params contains all the parameters to send to the API endpoint

	for the get lookup from package v1 operation.

	Typically these are written to a http.Request.
*/
type GetLookupFromPackageV1Params struct {

	/* Filename.

	   name of lookup file
	*/
	Filename string

	/* Package.

	   name of package
	*/
	Package string

	/* Repository.

	   name of repository
	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get lookup from package v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLookupFromPackageV1Params) WithDefaults() *GetLookupFromPackageV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get lookup from package v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLookupFromPackageV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) WithTimeout(timeout time.Duration) *GetLookupFromPackageV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) WithContext(ctx context.Context) *GetLookupFromPackageV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) WithHTTPClient(client *http.Client) *GetLookupFromPackageV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilename adds the filename to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) WithFilename(filename string) *GetLookupFromPackageV1Params {
	o.SetFilename(filename)
	return o
}

// SetFilename adds the filename to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) SetFilename(filename string) {
	o.Filename = filename
}

// WithPackage adds the packageVar to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) WithPackage(packageVar string) *GetLookupFromPackageV1Params {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WithRepository adds the repository to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) WithRepository(repository string) *GetLookupFromPackageV1Params {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get lookup from package v1 params
func (o *GetLookupFromPackageV1Params) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *GetLookupFromPackageV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filename
	if err := r.SetPathParam("filename", o.Filename); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}