package lbapi

// ????
//import (
//	"context"
//
//	"github.com/archaron/lbapi/types"
//)
//
//type GetStaffRequest struct {
//	AgentID   int                    `json:"agentid"`
//	VgID      int                    `json:"vgid"`
//	ParentID  int                    `json:"parentid"`
//	IP        string                 `json:"ip"`
//	Type      types.StaffAddressType `json:"type"`
//	BlockType types.StaffBlockType   `json:"blocktype"`
//}
//
//func (GetStaffRequest) Method() string {
//	return "getStaff"
//}
//
//func (api *Client) GetStaff(ctx context.Context, request GetStaffRequest) ([]types.VGroupRecord, error) {
//	var result []types.VGroupRecord
//	if err := api.client.CallResult(ctx, "getStaff", &request, &result); err != nil {
//		return nil, types.ParseJSONRPCError(err)
//	}
//	return result, nil
//}
