package lbapi

// {"id":1,"handler":"WebSessionHandler","params":{"event":"open","id":"vk5gd5uZLVlFydT98G6AKue8umHnLHP8ztTM9ZaUlNq4F1wumZ1ndFqBcB0TuLqD","time_stamp":0,"data":""}}
//{ "id": 1, "jsonrpc": "2.0", "result": { "authorized": true, "data": null, "id": "vk5gd5uZLVlFydT98G6AKue8umHnLHP8ztTM9ZaUlNq4F1wumZ1ndFqBcB0TuLqD", "success": true, "time_stamp": 1647043912 } }

// Cannot be implemented with current JSON RPC library

//
//import (
//	"context"
//
//	"github.com/archaron/lbapi/types"
//)
//
//type WebSessionHandlerRequest struct {
//	Event     string `json:"event"`
//	ID        string `json:"id"`
//	TimeStamp int64  `json:"timestamp"`
//	Data      string `json:"data"`
//}
//
//func (WebSessionHandlerRequest) Method() string { return "WebSessionHandler" }
//
//func (api *Client) WebSessionHandler(ctx context.Context, request WebSessionHandlerRequest) (interface{}, error) {
//	var response interface{}
//	if err := api.client.CallResult(ctx, "WebSessionHandler", &request, &response); err != nil {
//		return nil, types.ParseJSONRPCError(err)
//	}
//
//	return response, nil
//}
