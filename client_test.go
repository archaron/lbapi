package lbapi

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"
)

func TestRPCClient_CallResult_NilReturnsErrNotConnected(t *testing.T) {
	var client *rpcClient // nil wrapper

	err := client.CallResult(context.Background(), "test", nil, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNotConnected) {
		t.Fatalf("expected ErrNotConnected, got %v", err)
	}
}

func TestRPCClient_CallResult_NilInnerReturnsErrNotConnected(t *testing.T) {
	client := &rpcClient{c: nil}

	err := client.CallResult(context.Background(), "test", nil, nil)
	if !errors.Is(err, ErrNotConnected) {
		t.Fatalf("expected ErrNotConnected, got %v", err)
	}
}

func TestRPCClient_Close_NilSafe(t *testing.T) {
	var client *rpcClient
	if err := client.Close(); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	client = &rpcClient{c: nil}
	if err := client.Close(); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestClient_Call_ReturnsErrNotConnected(t *testing.T) {
	api := &Client{
		cfg: ClientConfig{Timeout: time.Second},
		log: slog.Default(),
	}
	// client field is nil — not connected
	err := api.Call(context.Background(), testRequest{}, new(int))
	if !errors.Is(err, ErrNotConnected) {
		t.Fatalf("expected ErrNotConnected, got %v", err)
	}
}

func TestClient_Call_AfterClose(t *testing.T) {
	api := &Client{
		cfg: ClientConfig{Timeout: time.Second},
		log: slog.Default(),
		client: &rpcClient{c: nil}, // wrapped nil inner client
	}
	err := api.Call(context.Background(), testRequest{}, new(int))
	if !errors.Is(err, ErrNotConnected) {
		t.Fatalf("expected ErrNotConnected after close, got %v", err)
	}
}

func TestClient_Close_CleansUp(t *testing.T) {
	api := &Client{
		cfg:        ClientConfig{Timeout: time.Second},
		log:        slog.Default(),
		client:     &rpcClient{c: nil},
		connection: nil,
	}
	if err := api.Close(); err != nil {
		t.Fatalf("close should not fail: %v", err)
	}
	// After close, Call should still return ErrNotConnected (no panic)
	err := api.Call(context.Background(), testRequest{}, new(int))
	if !errors.Is(err, ErrNotConnected) {
		t.Fatalf("expected ErrNotConnected after close, got %v", err)
	}
}

func TestNewClient_BackoffDefaults(t *testing.T) {
	api := NewClient(ClientConfig{
		Address: "127.0.0.1:0",
	}, slog.Default())
	if api.currentReconnectPeriod != 10*time.Second {
		t.Fatalf("expected default backoff min 10s, got %v", api.currentReconnectPeriod)
	}
	if api.cfg.ReconnectBackoffMax != 5*time.Minute {
		t.Fatalf("expected default backoff max 5m, got %v", api.cfg.ReconnectBackoffMax)
	}
}

func TestNewClient_CustomBackoff(t *testing.T) {
	api := NewClient(ClientConfig{
		Address:                "127.0.0.1:0",
		ReconnectBackoffMin:    5 * time.Second,
		ReconnectBackoffMax:    30 * time.Second,
		ReconnectBackoffFactor: 3,
	}, slog.Default())
	if api.currentReconnectPeriod != 5*time.Second {
		t.Fatalf("expected backoff min 5s, got %v", api.currentReconnectPeriod)
	}
	if api.cfg.ReconnectBackoffMax != 30*time.Second {
		t.Fatalf("expected backoff max 30s, got %v", api.cfg.ReconnectBackoffMax)
	}
	if api.cfg.ReconnectBackoffFactor != 3 {
		t.Fatalf("expected backoff factor 3, got %v", api.cfg.ReconnectBackoffFactor)
	}
}

// testRequest implements types.LanbillingRequest for test purposes
type testRequest struct{}

func (testRequest) Method() string { return "test" }