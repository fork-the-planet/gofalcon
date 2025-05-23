// Code generated by go-swagger; DO NOT EDIT.

package foundry_logscale

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewExecuteDynamicParams creates a new ExecuteDynamicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteDynamicParams() *ExecuteDynamicParams {
	return &ExecuteDynamicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteDynamicParamsWithTimeout creates a new ExecuteDynamicParams object
// with the ability to set a timeout on a request.
func NewExecuteDynamicParamsWithTimeout(timeout time.Duration) *ExecuteDynamicParams {
	return &ExecuteDynamicParams{
		timeout: timeout,
	}
}

// NewExecuteDynamicParamsWithContext creates a new ExecuteDynamicParams object
// with the ability to set a context for a request.
func NewExecuteDynamicParamsWithContext(ctx context.Context) *ExecuteDynamicParams {
	return &ExecuteDynamicParams{
		Context: ctx,
	}
}

// NewExecuteDynamicParamsWithHTTPClient creates a new ExecuteDynamicParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteDynamicParamsWithHTTPClient(client *http.Client) *ExecuteDynamicParams {
	return &ExecuteDynamicParams{
		HTTPClient: client,
	}
}

/*
ExecuteDynamicParams contains all the parameters to send to the API endpoint

	for the execute dynamic operation.

	Typically these are written to a http.Request.
*/
type ExecuteDynamicParams struct {

	/* XCSUSERUUID.

	   Requester UUID.
	*/
	XCSUSERUUID *string

	/* AppID.

	   Application ID.
	*/
	AppID *string

	// Body.
	Body *models.ApidomainDynamicExecuteSearchRequestV1

	/* IncludeSchemaGeneration.

	   Include generated schemas in the response
	*/
	IncludeSchemaGeneration *bool

	/* IncludeTestData.

	   Include test data when executing searches
	*/
	IncludeTestData *bool

	/* InferJSONTypes.

	   Whether to try to infer data types in json event response instead of returning map[string]string
	*/
	InferJSONTypes *bool

	/* MatchResponseSchema.

	   Whether to validate search results against their schema
	*/
	MatchResponseSchema *bool

	/* Metadata.

	   Whether to include metadata in the response
	*/
	Metadata *bool

	/* Mode.

	   Mode to execute the query under.
	*/
	Mode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute dynamic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteDynamicParams) WithDefaults() *ExecuteDynamicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute dynamic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteDynamicParams) SetDefaults() {
	var (
		includeSchemaGenerationDefault = bool(false)

		inferJSONTypesDefault = bool(false)

		matchResponseSchemaDefault = bool(false)

		metadataDefault = bool(false)
	)

	val := ExecuteDynamicParams{
		IncludeSchemaGeneration: &includeSchemaGenerationDefault,
		InferJSONTypes:          &inferJSONTypesDefault,
		MatchResponseSchema:     &matchResponseSchemaDefault,
		Metadata:                &metadataDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the execute dynamic params
func (o *ExecuteDynamicParams) WithTimeout(timeout time.Duration) *ExecuteDynamicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute dynamic params
func (o *ExecuteDynamicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute dynamic params
func (o *ExecuteDynamicParams) WithContext(ctx context.Context) *ExecuteDynamicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute dynamic params
func (o *ExecuteDynamicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute dynamic params
func (o *ExecuteDynamicParams) WithHTTPClient(client *http.Client) *ExecuteDynamicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute dynamic params
func (o *ExecuteDynamicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the execute dynamic params
func (o *ExecuteDynamicParams) WithXCSUSERUUID(xCSUSERUUID *string) *ExecuteDynamicParams {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the execute dynamic params
func (o *ExecuteDynamicParams) SetXCSUSERUUID(xCSUSERUUID *string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithAppID adds the appID to the execute dynamic params
func (o *ExecuteDynamicParams) WithAppID(appID *string) *ExecuteDynamicParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the execute dynamic params
func (o *ExecuteDynamicParams) SetAppID(appID *string) {
	o.AppID = appID
}

// WithBody adds the body to the execute dynamic params
func (o *ExecuteDynamicParams) WithBody(body *models.ApidomainDynamicExecuteSearchRequestV1) *ExecuteDynamicParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute dynamic params
func (o *ExecuteDynamicParams) SetBody(body *models.ApidomainDynamicExecuteSearchRequestV1) {
	o.Body = body
}

// WithIncludeSchemaGeneration adds the includeSchemaGeneration to the execute dynamic params
func (o *ExecuteDynamicParams) WithIncludeSchemaGeneration(includeSchemaGeneration *bool) *ExecuteDynamicParams {
	o.SetIncludeSchemaGeneration(includeSchemaGeneration)
	return o
}

// SetIncludeSchemaGeneration adds the includeSchemaGeneration to the execute dynamic params
func (o *ExecuteDynamicParams) SetIncludeSchemaGeneration(includeSchemaGeneration *bool) {
	o.IncludeSchemaGeneration = includeSchemaGeneration
}

// WithIncludeTestData adds the includeTestData to the execute dynamic params
func (o *ExecuteDynamicParams) WithIncludeTestData(includeTestData *bool) *ExecuteDynamicParams {
	o.SetIncludeTestData(includeTestData)
	return o
}

// SetIncludeTestData adds the includeTestData to the execute dynamic params
func (o *ExecuteDynamicParams) SetIncludeTestData(includeTestData *bool) {
	o.IncludeTestData = includeTestData
}

// WithInferJSONTypes adds the inferJSONTypes to the execute dynamic params
func (o *ExecuteDynamicParams) WithInferJSONTypes(inferJSONTypes *bool) *ExecuteDynamicParams {
	o.SetInferJSONTypes(inferJSONTypes)
	return o
}

// SetInferJSONTypes adds the inferJsonTypes to the execute dynamic params
func (o *ExecuteDynamicParams) SetInferJSONTypes(inferJSONTypes *bool) {
	o.InferJSONTypes = inferJSONTypes
}

// WithMatchResponseSchema adds the matchResponseSchema to the execute dynamic params
func (o *ExecuteDynamicParams) WithMatchResponseSchema(matchResponseSchema *bool) *ExecuteDynamicParams {
	o.SetMatchResponseSchema(matchResponseSchema)
	return o
}

// SetMatchResponseSchema adds the matchResponseSchema to the execute dynamic params
func (o *ExecuteDynamicParams) SetMatchResponseSchema(matchResponseSchema *bool) {
	o.MatchResponseSchema = matchResponseSchema
}

// WithMetadata adds the metadata to the execute dynamic params
func (o *ExecuteDynamicParams) WithMetadata(metadata *bool) *ExecuteDynamicParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the execute dynamic params
func (o *ExecuteDynamicParams) SetMetadata(metadata *bool) {
	o.Metadata = metadata
}

// WithMode adds the mode to the execute dynamic params
func (o *ExecuteDynamicParams) WithMode(mode *string) *ExecuteDynamicParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the execute dynamic params
func (o *ExecuteDynamicParams) SetMode(mode *string) {
	o.Mode = mode
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteDynamicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XCSUSERUUID != nil {

		// header param X-CS-USERUUID
		if err := r.SetHeaderParam("X-CS-USERUUID", *o.XCSUSERUUID); err != nil {
			return err
		}
	}

	if o.AppID != nil {

		// query param app_id
		var qrAppID string

		if o.AppID != nil {
			qrAppID = *o.AppID
		}
		qAppID := qrAppID
		if qAppID != "" {

			if err := r.SetQueryParam("app_id", qAppID); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.IncludeSchemaGeneration != nil {

		// query param include_schema_generation
		var qrIncludeSchemaGeneration bool

		if o.IncludeSchemaGeneration != nil {
			qrIncludeSchemaGeneration = *o.IncludeSchemaGeneration
		}
		qIncludeSchemaGeneration := swag.FormatBool(qrIncludeSchemaGeneration)
		if qIncludeSchemaGeneration != "" {

			if err := r.SetQueryParam("include_schema_generation", qIncludeSchemaGeneration); err != nil {
				return err
			}
		}
	}

	if o.IncludeTestData != nil {

		// query param include_test_data
		var qrIncludeTestData bool

		if o.IncludeTestData != nil {
			qrIncludeTestData = *o.IncludeTestData
		}
		qIncludeTestData := swag.FormatBool(qrIncludeTestData)
		if qIncludeTestData != "" {

			if err := r.SetQueryParam("include_test_data", qIncludeTestData); err != nil {
				return err
			}
		}
	}

	if o.InferJSONTypes != nil {

		// query param infer_json_types
		var qrInferJSONTypes bool

		if o.InferJSONTypes != nil {
			qrInferJSONTypes = *o.InferJSONTypes
		}
		qInferJSONTypes := swag.FormatBool(qrInferJSONTypes)
		if qInferJSONTypes != "" {

			if err := r.SetQueryParam("infer_json_types", qInferJSONTypes); err != nil {
				return err
			}
		}
	}

	if o.MatchResponseSchema != nil {

		// query param match_response_schema
		var qrMatchResponseSchema bool

		if o.MatchResponseSchema != nil {
			qrMatchResponseSchema = *o.MatchResponseSchema
		}
		qMatchResponseSchema := swag.FormatBool(qrMatchResponseSchema)
		if qMatchResponseSchema != "" {

			if err := r.SetQueryParam("match_response_schema", qMatchResponseSchema); err != nil {
				return err
			}
		}
	}

	if o.Metadata != nil {

		// query param metadata
		var qrMetadata bool

		if o.Metadata != nil {
			qrMetadata = *o.Metadata
		}
		qMetadata := swag.FormatBool(qrMetadata)
		if qMetadata != "" {

			if err := r.SetQueryParam("metadata", qMetadata); err != nil {
				return err
			}
		}
	}

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
