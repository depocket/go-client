package client

import (
	"context"
	"fmt"
)

type PoolService service

type PoolComponent struct {
	TokenAddress string `json:"token_address"`
	Type         string `json:"type"`
	Token        Token  `json:"token"`
}

type Pool struct {
	Address         string          `json:"address"`
	Type            string          `json:"type"`
	PoolIndex       int64           `json:"pool_index"`
	ProjectCode     string          `json:"project_code"`
	Symbol          string          `json:"symbol"`
	StakingStrategy string          `json:"staking_strategy"`
	FarmingAddress  string          `json:"farming_address"`
	PoolComponents  []PoolComponent `json:"pool_components"`
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

type PoolListByAddressesOptions struct {
	ListOptions
	Addresses []string `url:"addresses,omitempty"`
}

func (s *PoolService) ListByProjectCode(ctx context.Context, projectCode string, opts *PoolListOptions) ([]*Pool, *Response, error) {
	var u = fmt.Sprintf("%s/pools", projectCode)
	return s.listPoolByProjectCodes(ctx, u, opts)
}

func (s *PoolService) ListByAddresses(ctx context.Context, opts *PoolListByAddressesOptions) ([]*Pool, *Response, error) {
	var u = "pools"
	return s.listPoolByAddresses(ctx, u, opts)
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

func (s *PoolService) listPoolByAddresses(ctx context.Context, u string, opts *PoolListByAddressesOptions) ([]*Pool, *Response, error) {
	opts.Addresses = ConvertArrayOptsToApiParam(opts.Addresses)
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
