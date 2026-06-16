package lbapi

import (
	"context"
	"testing"
)

func TestRPCClient_CallResult_NilReturnsErrNotConnected(t *testing.T) {
	var client *rpcClient // nil wrapper

	err := client.CallResult(context.Background(), "test", nil, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != ErrNotConnected {
		t.Fatalf("expected ErrNotConnected, got %v", err)
	}
}

func TestRPCClient_CallResult_NilInnerReturnsErrNotConnected(t *testing.T) {
	client := &rpcClient{c: nil}

	err := client.CallResult(context.Background(), "test", nil, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != ErrNotConnected {
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