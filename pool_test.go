package client

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestPoolsService_List_By_Project_Code(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/project_code/pools", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": [{ "address":"0x0000", "pool_components":[{"name":"DePo"}]}], "error_code": 0 }`)
	})

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
}
