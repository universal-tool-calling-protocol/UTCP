package mcp

import (
	"context"
	"testing"

	. "github.com/universal-tool-calling-protocol/go-utcp/src/providers/mcp"
)

func TestMCPClientTransport_RegisterAndCall(t *testing.T) {
	tr := NewMCPTransport(nil)

	prov := NewMCPProvider("python3", []string{"python3", "../../../scripts/server.py"})
	ctx := context.Background()
	tools, err := tr.RegisterToolProvider(ctx, prov)
	if err != nil {
		t.Fatalf("register error: %v", err)
	}
	if tools == nil {
		t.Fatalf("expected non-nil tools")
	}

	if res, err := tr.CallTool(ctx, "hello", nil, prov, nil); err != nil {
		t.Fatalf("call error: %v", err)
	} else if res == nil {
		t.Fatalf("expected non-nil result")
	}

	if err := tr.DeregisterToolProvider(ctx, prov); err != nil {
		t.Fatalf("deregister error: %v", err)
	}
}
