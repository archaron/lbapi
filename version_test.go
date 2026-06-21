package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var versionResponse = json.RawMessage(`{"build":"51.0","release_date":"2025-05-06","revision":"9d2369b1","version":"2.0","version_supported":false}`)

func Test_Version(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call version success",
				out:  mustUnmarshal[VersionInfo](t, versionResponse),
				method: CallMethod[VersionRequest, VersionInfo](ApiCall[VersionRequest, VersionInfo]{
					req:    VersionRequest{},
					res:    mustUnmarshal[VersionInfo](t, versionResponse),
					mux:    mux,
					method: func(ctx context.Context, _ VersionRequest) (VersionInfo, error) {
						v, err := api.Version(ctx)
						if err != nil { return VersionInfo{}, err }
						return *v, nil
					},
				}),
			},
			{
				name: "call version with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[VersionRequest, VersionInfo](ApiCall[VersionRequest, VersionInfo]{
					req:    VersionRequest{},
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, _ VersionRequest) (VersionInfo, error) {
						v, err := api.Version(ctx)
						if err != nil { return VersionInfo{}, err }
						return *v, nil
					},
				}),
			},
		}
	})
}
