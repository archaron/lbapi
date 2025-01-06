package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// getPorts
type (
	GetPortsRequest struct {
		DeviceId int `json:"device_id,omitempty"`
		PortId   int `json:"port_id,omitempty"`
		VgId     int `json:"vg_id,omitempty"`
	}

	PortInfo struct {
		Comment     string         `json:"comment,omitempty"`
		DeviceId    int            `json:"device_id,omitempty"`
		Flush       int            `json:"flush,omitempty"`
		InnerVlan   int            `json:"inner_vlan,omitempty"`
		OuterVlan   int            `json:"outer_vlan,omitempty"`
		Login       string         `json:"login,omitempty"`
		Media       string         `json:"media,omitempty"`
		Name        int            `json:"name,omitempty"`
		PolicyId    *int           `json:"policy_id,omitempty"`
		PrototypeId *int           `json:"prototype_id,omitempty"`
		PortId      int            `json:"port_id,omitempty"`
		Speed       string         `json:"speed,omitempty"`
		Status      int            `json:"status"`
		Tpl         int            `json:"tpl,omitempty"`
		VlanId      int            `json:"vlan_id,omitempty"`
		VgId        int            `json:"vg_id,omitempty"`
		TimeMark    *string        `json:"time_mark"`
		Timestamp   int            `json:"timestamp"`
		TreeItem    []PortTreeItem `json:"tree_item"`
	}

	PortTreeItem struct {
		DeviceId         int    `json:"device_id,omitempty"`
		DeviceName       string `json:"device_name"`
		ParentDeviceId   int    `json:"parent_device_id"`
		ParentDeviceName string `json:"parent_device_name"`
		ParentPortId     int    `json:"parent_port_id"`
		ParentPortName   int    `json:"parent_port_name"`
		PortId           int    `json:"port_id"`
		PortName         int    `json:"port_name"`
	}
)

func (GetPortsRequest) Method() string {
	return "getPorts"
}

func (api *Client) GetPorts(ctx context.Context, searchParams GetPortsRequest) ([]PortInfo, error) {

	var result []PortInfo
	if err := api.client.CallResult(ctx, "getPorts", &searchParams, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
