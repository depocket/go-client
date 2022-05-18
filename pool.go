package client

import (
	"context"
	"fmt"
)

type PoolService service

type PoolComponent struct {
	TokenAddress string `json:"token_address"`
	Type         string `json:"type"`
}

type Pool struct {
	Address        string          `json:"address"`
	Type           string          `json:"type"`
	PoolIndex      int64           `json:"pool_index"`
	ProjectCode    string          `json:"project_code"`
	PoolComponents []PoolComponent `json:"pool_components"`
}

type PoolsResponse struct {
	Data      []*Pool `json:"data"`
	ErrorCode int     `json:"error_code"`
}

func (i Pool) String() string {
	return Stringify(i)
}

type PoolListOptions struct {
	ListOptions
}

func (s *PoolService) ListByProjectCode(ctx context.Context, projectCode string, opts *PoolListOptions) ([]*Pool, *Response, error) {
	var u = fmt.Sprintf("%s/pools", projectCode)
	return s.listPoolByProjectCodes(ctx, u, opts)
}

func (s *PoolService) listPoolByProjectCodes(ctx context.Context, u string, opts *PoolListOptions) ([]*Pool, *Response, error) {
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var poolsResponse *PoolsResponse
	resp, err := s.client.Do(ctx, req, &poolsResponse)
	if err != nil {
		return nil, resp, err
	}

	if poolsResponse.ErrorCode != 0 {
		return nil, resp, fmt.Errorf("error code %d", poolsResponse.ErrorCode)
	}

	return poolsResponse.Data, resp, nil
}
