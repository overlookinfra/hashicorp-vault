// A minimal UI for simple testing via a UI without Vault
package jwtauth

import (
	"context"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathUI(b *jwtAuthBackend) *framework.Path {
	return &framework.Path{
		Pattern: `ui$`,

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation: b.pathUI,
		},
	}
}

func (b *jwtAuthBackend) pathUI(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	resp := &logical.Response{
		Data: map[string]interface{}{
			logical.HTTPStatusCode:  200,
			logical.HTTPRawBody:     testUIHTMLstr,
			logical.HTTPContentType: "text/html",
		},
	}

	return resp, nil
}
