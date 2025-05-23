// Code generated by go-swagger; DO NOT EDIT.

package ngsiem

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ngsiem API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ngsiem API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLookupFromPackageV1(params *GetLookupFromPackageV1Params, opts ...ClientOption) (*GetLookupFromPackageV1OK, error)

	GetLookupFromPackageWithNamespaceV1(params *GetLookupFromPackageWithNamespaceV1Params, opts ...ClientOption) (*GetLookupFromPackageWithNamespaceV1OK, error)

	GetLookupV1(params *GetLookupV1Params, opts ...ClientOption) (*GetLookupV1OK, error)

	GetSearchStatusV1(params *GetSearchStatusV1Params, opts ...ClientOption) (*GetSearchStatusV1OK, error)

	StartSearchV1(params *StartSearchV1Params, opts ...ClientOption) (*StartSearchV1OK, error)

	StopSearchV1(params *StopSearchV1Params, opts ...ClientOption) (*StopSearchV1OK, error)

	UploadLookupV1(params *UploadLookupV1Params, opts ...ClientOption) (*UploadLookupV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetLookupFromPackageV1 downloads lookup file in package from n g s i e m
*/
func (a *Client) GetLookupFromPackageV1(params *GetLookupFromPackageV1Params, opts ...ClientOption) (*GetLookupFromPackageV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLookupFromPackageV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLookupFromPackageV1",
		Method:             "GET",
		PathPattern:        "/humio/api/v1/repositories/{repository}/files/{package}/{filename}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLookupFromPackageV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLookupFromPackageV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLookupFromPackageV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLookupFromPackageWithNamespaceV1 downloads lookup file in namespaced package from n g s i e m
*/
func (a *Client) GetLookupFromPackageWithNamespaceV1(params *GetLookupFromPackageWithNamespaceV1Params, opts ...ClientOption) (*GetLookupFromPackageWithNamespaceV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLookupFromPackageWithNamespaceV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLookupFromPackageWithNamespaceV1",
		Method:             "GET",
		PathPattern:        "/humio/api/v1/repositories/{repository}/files/{namespace}/{package}/{filename}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLookupFromPackageWithNamespaceV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLookupFromPackageWithNamespaceV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLookupFromPackageWithNamespaceV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLookupV1 downloads lookup file from n g s i e m
*/
func (a *Client) GetLookupV1(params *GetLookupV1Params, opts ...ClientOption) (*GetLookupV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLookupV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLookupV1",
		Method:             "GET",
		PathPattern:        "/humio/api/v1/repositories/{repository}/files/{filename}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLookupV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLookupV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLookupV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSearchStatusV1 gets status of search
*/
func (a *Client) GetSearchStatusV1(params *GetSearchStatusV1Params, opts ...ClientOption) (*GetSearchStatusV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchStatusV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSearchStatusV1",
		Method:             "GET",
		PathPattern:        "/humio/api/v1/repositories/{repository}/queryjobs/{id}",
		ProducesMediaTypes: []string{"application/json", "application/x-ndjson", "text/html", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchStatusV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSearchStatusV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSearchStatusV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartSearchV1 initiates search
*/
func (a *Client) StartSearchV1(params *StartSearchV1Params, opts ...ClientOption) (*StartSearchV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartSearchV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartSearchV1",
		Method:             "POST",
		PathPattern:        "/humio/api/v1/repositories/{repository}/queryjobs",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartSearchV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartSearchV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StartSearchV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopSearchV1 stops search
*/
func (a *Client) StopSearchV1(params *StopSearchV1Params, opts ...ClientOption) (*StopSearchV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopSearchV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "StopSearchV1",
		Method:             "DELETE",
		PathPattern:        "/humio/api/v1/repositories/{repository}/queryjobs/{id}",
		ProducesMediaTypes: []string{"*/*", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopSearchV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopSearchV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StopSearchV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UploadLookupV1 uploads file to n g s i e m
*/
func (a *Client) UploadLookupV1(params *UploadLookupV1Params, opts ...ClientOption) (*UploadLookupV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadLookupV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "UploadLookupV1",
		Method:             "POST",
		PathPattern:        "/humio/api/v1/repositories/{repository}/files",
		ProducesMediaTypes: []string{"*/*", "application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadLookupV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadLookupV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UploadLookupV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
