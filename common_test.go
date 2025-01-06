package lbapi

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/handler"
	"github.com/creachadair/jrpc2/server"
	"github.com/stretchr/testify/require"

	"github.com/archaron/lbapi/types"
)

type (
	MethodCall func(context.Context) (any, error)

	APICallOneFunc[Req types.LanbillingRequest, Res any] func(context.Context) (Res, error)

	APIMux struct {
		items handler.Map
		sync.Mutex
	}

	ApiCall[Req types.LanbillingRequest, Res any] struct {
		req Req
		res Res
		err error
		mux *APIMux

		method func(ctx context.Context, req Req) (Res, error)
	}

	APITestCase struct {
		name string
		err  error
		out  any

		method MethodCall
	}
)

func (a APICallOneFunc[Req, Res]) call(ctx context.Context, _ Req) (Res, error) {
	return a(ctx)
}

func (a *APIMux) Assign(ctx context.Context, method string) jrpc2.Handler {
	a.Lock()
	defer a.Unlock()
	return a.items.Assign(ctx, method)
}

func (a *APIMux) Add(method string, handle jrpc2.Handler) {
	a.Lock()
	defer a.Unlock()
	if a.items == nil {
		a.items = make(handler.Map)
	}
	a.items[method] = handle
}

func (a *APIMux) Del(method string) {
	a.Lock()
	defer a.Unlock()
	delete(a.items, method)
}

func CallMethod[Req types.LanbillingRequest, Res any](opts ApiCall[Req, Res]) MethodCall {

	return func(ctx context.Context) (any, error) {

		opts.mux.Add(opts.req.Method(), func(context.Context, *jrpc2.Request) (interface{}, error) {
			return opts.res, opts.err
		})

		defer opts.mux.Del(opts.req.Method())

		res, err := opts.method(ctx, opts.req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

}

func mustUnmarshal[T any](t *testing.T, raw json.RawMessage) T {
	var result T

	require.NoError(t, json.Unmarshal(raw, &result))

	return result
}

func runAPITests(t *testing.T, casesGenerator func(*Client, *APIMux) []APITestCase) {
	t.Helper()
	mux := new(APIMux)
	api := &Client{client: server.NewLocal(mux, nil).Client}
	cases := casesGenerator(api, mux)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.method(context.Background())
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.out, actual)
		})
	}
}
