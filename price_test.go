package client

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestPricesService(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/prices", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, `{ "data":{ "0xcc7a91413769891de2e9ebbfc96d2eb1874b5760": { "usd": 0.5 }, "0x6b9481fb5465ef9ab9347b332058d894ab09b2f6": { "usd": 0.8 } }, "error_code": 0 }`)
	})

	t.Run("ListByAddresses", func(t *testing.T) {
		opt := &PriceListOptions{
			ListOptions: ListOptions{
				Chain: "bsc",
			},
			Addresses: []string{"0xcc7a91413769891de2e9ebbfc96d2eb1874b5760", "0x6b9481fb5465ef9ab9347b332058d894ab09b2f6"},
		}
		ctx := context.Background()
		tokens, _, err := client.Prices.List(ctx, opt)
		if err != nil {
			t.Errorf("Prices.List returned error: %v", err)
		}

		want := map[string]*Price{
			"0xcc7a91413769891de2e9ebbfc96d2eb1874b5760": {
				Usd:          0.5,
				Usd24hChange: 0,
			},
			"0x6b9481fb5465ef9ab9347b332058d894ab09b2f6": {
				Usd:          0.8,
				Usd24hChange: 0,
			},
		}
		if !cmp.Equal(tokens, want) {
			t.Errorf("Prices.List returned %+v, want %+v", tokens, want)
		}
	})
}
