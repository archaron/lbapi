package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// Login
type (
	LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

func (LoginRequest) Method() string {
	return "Login"
}

func (api *Client) Login(ctx context.Context, request LoginRequest) (bool, error) {

	var result bool
	if err := api.client.CallResult(ctx, "Login", request, &result); err != nil {
		return false, types.NewWrappedError(types.ErrClient, err)
	}

	return result, nil
}
