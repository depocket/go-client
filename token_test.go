package client

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestTokensService(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		query := r.URL.Query()

		if len(query.Get("addresses")) > 0 || len(query.Get("symbols")) > 0 {
			_, _ = fmt.Fprint(w, `{ "data":[ { "name":"GovWorld" }, { "name":"Dino Land" } ], "error_code": 0 }`)
		} else {
			_, _ = fmt.Fprint(w, `{ "data":[ { "name":"DePo" } ], "error_code": 0 }`)
		}
	})

	t.Run("ListAll", func(t *testing.T) {
		opt := &TokenListOptions{
			ListOptions: ListOptions{
				Chain: "bsc",
			},
		}
		ctx := context.Background()
		tokens, _, err := client.Tokens.List(ctx, opt)
		if err != nil {
			t.Errorf("Tokens.List returned error: %v", err)
		}

		want := []*Token{{Name: "DePo"}}
		if !cmp.Equal(tokens, want) {
			t.Errorf("Token.List returned %+v, want %+v", tokens, want)
		}
	})

	t.Run("ListByAddresses", func(t *testing.T) {
		opt := &TokenListOptions{
			ListOptions: ListOptions{
				Chain: "bsc",
			},
			Addresses: "0xcc7a91413769891de2e9ebbfc96d2eb1874b5760,0x6b9481fb5465ef9ab9347b332058d894ab09b2f6",
		}
		ctx := context.Background()
		tokens, _, err := client.Tokens.List(ctx, opt)
		if err != nil {
			t.Errorf("Tokens.List returned error: %v", err)
		}

		want := []*Token{{Name: "GovWorld"}, {Name: "Dino Land"}}
		if !cmp.Equal(tokens, want) {
			t.Errorf("Token.List returned %+v, want %+v", tokens, want)
		}
	})

	t.Run("ListBySymbols", func(t *testing.T) {
		opt := &TokenListOptions{
			ListOptions: ListOptions{
				Chain: "bsc",
			},
			Symbols: "GOV,DNL",
		}
		ctx := context.Background()
		tokens, _, err := client.Tokens.List(ctx, opt)
		if err != nil {
			t.Errorf("Tokens.List returned error: %v", err)
		}

		want := []*Token{{Name: "GovWorld"}, {Name: "Dino Land"}}
		if !cmp.Equal(tokens, want) {
			t.Errorf("Token.List returned %+v, want %+v", tokens, want)
		}
	})
}
