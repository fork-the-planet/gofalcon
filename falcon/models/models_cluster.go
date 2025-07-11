// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsCluster models cluster
//
// swagger:model models.Cluster
type ModelsCluster struct {

	// access
	// Required: true
	Access *string `json:"access"`

	// agent status
	// Required: true
	AgentStatus *string `json:"agent_status"`

	// agents
	// Required: true
	Agents []ModelsClusterAgents `json:"agents"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// cloud account id
	// Required: true
	CloudAccountID *string `json:"cloud_account_id"`

	// cloud name
	// Required: true
	CloudName *string `json:"cloud_name"`

	// cloud region
	// Required: true
	CloudRegion *string `json:"cloud_region"`

	// cloud service
	// Required: true
	CloudService *string `json:"cloud_service"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// cluster name
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// cluster security group
	// Required: true
	ClusterSecurityGroup *string `json:"cluster_security_group"`

	// cluster status
	// Required: true
	ClusterStatus *string `json:"cluster_status"`

	// container count
	// Required: true
	ContainerCount *int32 `json:"container_count"`

	// first seen
	// Required: true
	FirstSeen *string `json:"first_seen"`

	// iar coverage
	// Required: true
	IarCoverage *bool `json:"iar_coverage"`

	// kac agent active
	// Required: true
	KacAgentActive *bool `json:"kac_agent_active"`

	// kac agent id
	// Required: true
	KacAgentID *string `json:"kac_agent_id"`

	// kubernetes version
	// Required: true
	KubernetesVersion *string `json:"kubernetes_version"`

	// labels list
	// Required: true
	LabelsList []string `json:"labels_list"`

	// last seen
	// Required: true
	LastSeen *string `json:"last_seen"`

	// management status
	// Required: true
	ManagementStatus *string `json:"management_status"`

	// node count
	// Required: true
	NodeCount *int32 `json:"node_count"`

	// pod count
	// Required: true
	PodCount *int32 `json:"pod_count"`

	// security group
	// Required: true
	SecurityGroup *string `json:"security_group"`

	// tags
	// Required: true
	Tags map[string]string `json:"tags"`

	// virtual network
	// Required: true
	VirtualNetwork *string `json:"virtual_network"`
}

// Validate validates this models cluster
func (m *ModelsCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIarCoverage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKacAgentActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKacAgentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCluster) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateAgentStatus(formats strfmt.Registry) error {

	if err := validate.Required("agent_status", "body", m.AgentStatus); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateAgents(formats strfmt.Registry) error {

	if err := validate.Required("agents", "body", m.Agents); err != nil {
		return err
	}

	for i := 0; i < len(m.Agents); i++ {

		if m.Agents[i] != nil {
			if err := m.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsCluster) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateCloudAccountID(formats strfmt.Registry) error {

	if err := validate.Required("cloud_account_id", "body", m.CloudAccountID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateCloudName(formats strfmt.Registry) error {

	if err := validate.Required("cloud_name", "body", m.CloudName); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateCloudRegion(formats strfmt.Registry) error {

	if err := validate.Required("cloud_region", "body", m.CloudRegion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateCloudService(formats strfmt.Registry) error {

	if err := validate.Required("cloud_service", "body", m.CloudService); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("cluster_name", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateClusterSecurityGroup(formats strfmt.Registry) error {

	if err := validate.Required("cluster_security_group", "body", m.ClusterSecurityGroup); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateClusterStatus(formats strfmt.Registry) error {

	if err := validate.Required("cluster_status", "body", m.ClusterStatus); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateContainerCount(formats strfmt.Registry) error {

	if err := validate.Required("container_count", "body", m.ContainerCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateFirstSeen(formats strfmt.Registry) error {

	if err := validate.Required("first_seen", "body", m.FirstSeen); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateIarCoverage(formats strfmt.Registry) error {

	if err := validate.Required("iar_coverage", "body", m.IarCoverage); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateKacAgentActive(formats strfmt.Registry) error {

	if err := validate.Required("kac_agent_active", "body", m.KacAgentActive); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateKacAgentID(formats strfmt.Registry) error {

	if err := validate.Required("kac_agent_id", "body", m.KacAgentID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateKubernetesVersion(formats strfmt.Registry) error {

	if err := validate.Required("kubernetes_version", "body", m.KubernetesVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateLabelsList(formats strfmt.Registry) error {

	if err := validate.Required("labels_list", "body", m.LabelsList); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateLastSeen(formats strfmt.Registry) error {

	if err := validate.Required("last_seen", "body", m.LastSeen); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateManagementStatus(formats strfmt.Registry) error {

	if err := validate.Required("management_status", "body", m.ManagementStatus); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("node_count", "body", m.NodeCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validatePodCount(formats strfmt.Registry) error {

	if err := validate.Required("pod_count", "body", m.PodCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateSecurityGroup(formats strfmt.Registry) error {

	if err := validate.Required("security_group", "body", m.SecurityGroup); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCluster) validateVirtualNetwork(formats strfmt.Registry) error {

	if err := validate.Required("virtual_network", "body", m.VirtualNetwork); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this models cluster based on the context it is used
func (m *ModelsCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCluster) contextValidateAgents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Agents); i++ {

		if swag.IsZero(m.Agents[i]) { // not required
			return nil
		}

		if err := m.Agents[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agents" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agents" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsCluster) UnmarshalBinary(b []byte) error {
	var res ModelsCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
