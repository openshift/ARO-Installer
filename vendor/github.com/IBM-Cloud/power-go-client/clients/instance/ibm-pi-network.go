package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPINetworkClient ...
type IBMPINetworkClient struct {
	IBMPIClient
}

// NewIBMPINetworkClient ...
func NewIBMPINetworkClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPINetworkClient {
	return &IBMPINetworkClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get ...
func (f *IBMPINetworkClient) Get(id string) (*models.Network, error) {
	params := p_cloud_networks.NewPcloudNetworksGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetNetworkOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network %s", id)
	}
	return resp.Payload, nil
}

// Get All
func (f *IBMPINetworkClient) GetAll() (*models.Networks, error) {
	params := p_cloud_networks.NewPcloudNetworksGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create ...
func (f *IBMPINetworkClient) Create(body *models.NetworkCreate) (*models.Network, error) {
	params := p_cloud_networks.NewPcloudNetworksPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, postcreated, err := f.session.Power.PCloudNetworks.PcloudNetworksPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateNetworkOperationFailed, body.Name, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	if postcreated != nil && postcreated.Payload != nil {
		return postcreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to perform Create Network Operation for Network %s", body.Name)
}

// Update ...
func (f *IBMPINetworkClient) Update(id string, body *models.NetworkUpdate) (*models.Network, error) {
	params := p_cloud_networks.NewPcloudNetworksPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s with error %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s", id)
	}
	return resp.Payload, nil
}

// GetPublic ...
func (f *IBMPINetworkClient) GetAllPublic() (*models.Networks, error) {
	filterQuery := "type=\"pub-vlan\""
	params := p_cloud_networks.NewPcloudNetworksGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithFilter(&filterQuery)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Delete ...
func (f *IBMPINetworkClient) Delete(id string) error {
	params := p_cloud_networks.NewPcloudNetworksDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id)
	_, err := f.session.Power.PCloudNetworks.PcloudNetworksDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PI Network %s: %w", id, err)
	}
	return nil
}

//GetAllPorts ...
func (f *IBMPINetworkClient) GetAllPorts(id string) (*models.NetworkPorts, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s", id)
	}
	return resp.Payload, nil
}

// GetPort ...
func (f *IBMPINetworkClient) GetPort(id string, networkPortID string) (*models.NetworkPort, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithPortID(networkPortID)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsGet(params, f.session.AuthInfo(f.cloudInstanceID))

	if err != nil {
		return nil, fmt.Errorf("failed to Get PI Network Port %s for Network %s: %w", networkPortID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PI Network Port %s for Network %s", networkPortID, id)
	}
	return resp.Payload, nil
}

// CreatePort ...
func (f *IBMPINetworkClient) CreatePort(id string, body *models.NetworkPortCreate) (*models.NetworkPort, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateNetworkPortOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Create Network Port Operation for Network %s", id)
	}
	return resp.Payload, nil
}

// DeletePort ...
func (f *IBMPINetworkClient) DeletePort(id string, networkPortID string) error {
	params := p_cloud_networks.NewPcloudNetworksPortsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithPortID(networkPortID)
	_, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to delete the network port %s for network %s with error %w", networkPortID, id, err)
	}
	return nil
}

//UpdatePort with the PVM Instance
func (f *IBMPINetworkClient) UpdatePort(id, networkPortID string, body *models.NetworkPortUpdate) (*models.NetworkPort, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithPortID(networkPortID).WithBody(body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s with error %w", networkPortID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s", networkPortID, id)
	}
	return resp.Payload, nil
}
