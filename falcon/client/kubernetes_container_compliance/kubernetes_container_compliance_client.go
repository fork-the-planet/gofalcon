// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_container_compliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new kubernetes container compliance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kubernetes container compliance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AggregateAssessmentsGroupedByClustersV2(params *AggregateAssessmentsGroupedByClustersV2Params, opts ...ClientOption) (*AggregateAssessmentsGroupedByClustersV2OK, error)

	AggregateAssessmentsGroupedByRulesV2(params *AggregateAssessmentsGroupedByRulesV2Params, opts ...ClientOption) (*AggregateAssessmentsGroupedByRulesV2OK, error)

	AggregateComplianceByAssetType(params *AggregateComplianceByAssetTypeParams, opts ...ClientOption) (*AggregateComplianceByAssetTypeOK, error)

	AggregateComplianceByClusterType(params *AggregateComplianceByClusterTypeParams, opts ...ClientOption) (*AggregateComplianceByClusterTypeOK, error)

	AggregateComplianceByFramework(params *AggregateComplianceByFrameworkParams, opts ...ClientOption) (*AggregateComplianceByFrameworkOK, error)

	AggregateFailedRulesByClustersV3(params *AggregateFailedRulesByClustersV3Params, opts ...ClientOption) (*AggregateFailedRulesByClustersV3OK, error)

	AggregateTopFailedImages(params *AggregateTopFailedImagesParams, opts ...ClientOption) (*AggregateTopFailedImagesOK, error)

	CombinedImagesFindings(params *CombinedImagesFindingsParams, opts ...ClientOption) (*CombinedImagesFindingsOK, error)

	CombinedNodesFindings(params *CombinedNodesFindingsParams, opts ...ClientOption) (*CombinedNodesFindingsOK, error)

	GetRulesMetadataByID(params *GetRulesMetadataByIDParams, opts ...ClientOption) (*GetRulesMetadataByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AggregateAssessmentsGroupedByClustersV2 returns cluster details along with aggregated assessment results organized by cluster including pass fail assessment counts for various asset types
*/
func (a *Client) AggregateAssessmentsGroupedByClustersV2(params *AggregateAssessmentsGroupedByClustersV2Params, opts ...ClientOption) (*AggregateAssessmentsGroupedByClustersV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateAssessmentsGroupedByClustersV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateAssessmentsGroupedByClustersV2",
		Method:             "GET",
		PathPattern:        "/container-compliance/aggregates/clusters/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateAssessmentsGroupedByClustersV2Reader{formats: a.formats},
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
	success, ok := result.(*AggregateAssessmentsGroupedByClustersV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AggregateAssessmentsGroupedByClustersV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AggregateAssessmentsGroupedByRulesV2 returns rule details along with aggregated assessment results organized by compliance rule including pass fail assessment counts
*/
func (a *Client) AggregateAssessmentsGroupedByRulesV2(params *AggregateAssessmentsGroupedByRulesV2Params, opts ...ClientOption) (*AggregateAssessmentsGroupedByRulesV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateAssessmentsGroupedByRulesV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateAssessmentsGroupedByRulesV2",
		Method:             "GET",
		PathPattern:        "/container-compliance/aggregates/rules/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateAssessmentsGroupedByRulesV2Reader{formats: a.formats},
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
	success, ok := result.(*AggregateAssessmentsGroupedByRulesV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AggregateAssessmentsGroupedByRulesV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AggregateComplianceByAssetType provides aggregated compliance assessment metrics and rule status information organized by asset type
*/
func (a *Client) AggregateComplianceByAssetType(params *AggregateComplianceByAssetTypeParams, opts ...ClientOption) (*AggregateComplianceByAssetTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateComplianceByAssetTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateComplianceByAssetType",
		Method:             "GET",
		PathPattern:        "/container-compliance/aggregates/compliance-by-asset-type/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateComplianceByAssetTypeReader{formats: a.formats},
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
	success, ok := result.(*AggregateComplianceByAssetTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AggregateComplianceByAssetType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AggregateComplianceByClusterType provides aggregated compliance assessment metrics and rule status information organized by kubernetes cluster type
*/
func (a *Client) AggregateComplianceByClusterType(params *AggregateComplianceByClusterTypeParams, opts ...ClientOption) (*AggregateComplianceByClusterTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateComplianceByClusterTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateComplianceByClusterType",
		Method:             "GET",
		PathPattern:        "/container-compliance/aggregates/compliance-by-cluster-type/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateComplianceByClusterTypeReader{formats: a.formats},
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
	success, ok := result.(*AggregateComplianceByClusterTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AggregateComplianceByClusterType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AggregateComplianceByFramework provides aggregated compliance assessment metrics and rule status information organized by compliance framework
*/
func (a *Client) AggregateComplianceByFramework(params *AggregateComplianceByFrameworkParams, opts ...ClientOption) (*AggregateComplianceByFrameworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateComplianceByFrameworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateComplianceByFramework",
		Method:             "GET",
		PathPattern:        "/container-compliance/aggregates/compliance-by-framework/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateComplianceByFrameworkReader{formats: a.formats},
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
	success, ok := result.(*AggregateComplianceByFrameworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AggregateComplianceByFramework: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AggregateFailedRulesByClustersV3 retrieves the most non compliant clusters ranked in descending order based on the number of failed compliance rules across severity levels critical high medium and low
*/
func (a *Client) AggregateFailedRulesByClustersV3(params *AggregateFailedRulesByClustersV3Params, opts ...ClientOption) (*AggregateFailedRulesByClustersV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateFailedRulesByClustersV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateFailedRulesByClustersV3",
		Method:             "GET",
		PathPattern:        "/container-compliance/aggregates/failed-rules-by-clusters/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateFailedRulesByClustersV3Reader{formats: a.formats},
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
	success, ok := result.(*AggregateFailedRulesByClustersV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AggregateFailedRulesByClustersV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AggregateTopFailedImages retrieves the most non compliant container images ranked in descending order based on the number of failed assessments across severity levels critical high medium and low
*/
func (a *Client) AggregateTopFailedImages(params *AggregateTopFailedImagesParams, opts ...ClientOption) (*AggregateTopFailedImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateTopFailedImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateTopFailedImages",
		Method:             "GET",
		PathPattern:        "/container-compliance/aggregates/top-failed-images/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateTopFailedImagesReader{formats: a.formats},
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
	success, ok := result.(*AggregateTopFailedImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AggregateTopFailedImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CombinedImagesFindings returns detailed compliance assessment results for container images providing the information needed to identify compliance violations
*/
func (a *Client) CombinedImagesFindings(params *CombinedImagesFindingsParams, opts ...ClientOption) (*CombinedImagesFindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombinedImagesFindingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CombinedImagesFindings",
		Method:             "GET",
		PathPattern:        "/container-compliance/combined/findings-by-images/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombinedImagesFindingsReader{formats: a.formats},
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
	success, ok := result.(*CombinedImagesFindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CombinedImagesFindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CombinedNodesFindings returns detailed compliance assessment results for kubernetes nodes providing the information needed to identify compliance violations
*/
func (a *Client) CombinedNodesFindings(params *CombinedNodesFindingsParams, opts ...ClientOption) (*CombinedNodesFindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombinedNodesFindingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CombinedNodesFindings",
		Method:             "GET",
		PathPattern:        "/container-compliance/combined/findings-by-nodes/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombinedNodesFindingsReader{formats: a.formats},
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
	success, ok := result.(*CombinedNodesFindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CombinedNodesFindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRulesMetadataByID retrieves detailed compliance rule information including descriptions remediation steps and audit procedures by specifying rule identifiers
*/
func (a *Client) GetRulesMetadataByID(params *GetRulesMetadataByIDParams, opts ...ClientOption) (*GetRulesMetadataByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRulesMetadataByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRulesMetadataByID",
		Method:             "GET",
		PathPattern:        "/container-compliance/combined/rule-details-by-rule-ids/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRulesMetadataByIDReader{formats: a.formats},
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
	success, ok := result.(*GetRulesMetadataByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRulesMetadataByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
