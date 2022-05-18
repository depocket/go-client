package client

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestTokensService_List_All(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "data":[ { "name":"DePo" } ], "error_code": 0 }`)
	})

	opt := &TokenListOptions{
		ListOptions{Chain: "bsc"},
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
}
