package client

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestPoolsService(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/project_code/pools", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{ "address":"0x0000", "pool_components":[{"name":"DePo"}]}], "error_code": 0 }`)
	})
	mux.HandleFunc("/pools", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{ "address":"0x0001", "pool_components":[{"name":"DePo"}]}], "error_code": 0 }`)
	})

	t.Run("ListByProjectCode", func(t *testing.T) {
		opt := &PoolListOptions{
			ListOptions{Chain: "bsc"},
		}
		ctx := context.Background()
		pools, _, err := client.Pools.ListByProjectCode(ctx, "project_code", opt)
		if err != nil {
			t.Errorf("Pools.ListByProjectCode returned error: %v", err)
		}
		fmt.Println(pools)

		want := []*Pool{{Address: "0x0000", Type: "", PoolIndex: 0, ProjectCode: "", PoolComponents: []PoolComponent{{TokenAddress: "", Type: ""}}}}
		if !cmp.Equal(pools, want) {
			t.Errorf("Pool.List returned %+v, want %+v", pools, want)
		}
	})

	t.Run("ListByAddresses", func(t *testing.T) {
		opt := &PoolListByAddressesOptions{
			ListOptions: ListOptions{Chain: "bsc"},
			Addresses: []string{
				"0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
				"0x0ed7e52944161450477ee417de9cd3a859b14fd0",
			},
		}
		ctx := context.Background()
		pools, _, err := client.Pools.ListByAddresses(ctx, opt)
		if err != nil {
			t.Errorf("Pools.ListByAddresses returned error: %v", err)
		}
		fmt.Println(pools)

		want := []*Pool{{Address: "0x0001", Type: "", PoolIndex: 0, ProjectCode: "", PoolComponents: []PoolComponent{{TokenAddress: "", Type: ""}}}}
		if !cmp.Equal(pools, want) {
			t.Errorf("Pool.List returned %+v, want %+v", pools, want)
		}
	})
}
