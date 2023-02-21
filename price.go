package client

import (
	"context"
	"fmt"
)

type PriceService service

type Price struct {
	Usd          float64 `json:"usd"`
	Usd24hChange float64 `json:"usd_24h_change"`
}

type PriceResponse struct {
	Data      map[string]*Price `json:"data"`
	ErrorCode int               `json:"error_code"`
}

func (i PriceResponse) String() string {
	return Stringify(i)
}

type PriceListOptions struct {
	ListOptions
	Addresses []string `url:"addresses,omitempty"`
}

func (s *PriceService) List(ctx context.Context, opts *PriceListOptions) (map[string]*Price, *Response, error) {
	var u = "prices"
	return s.listTokensMapPrice(ctx, u, opts)
}

func (s *PriceService) listTokensMapPrice(ctx context.Context, u string, opts *PriceListOptions) (map[string]*Price, *Response, error) {
	opts.Addresses = ConvertArrayOptsToApiParam(opts.Addresses)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var pricesResponse *PriceResponse
	resp, err := s.client.Do(ctx, req, &pricesResponse)
	if err != nil {
		return nil, resp, err
	}
	if pricesResponse.ErrorCode != 0 {
		return nil, resp, fmt.Errorf("error code %d", pricesResponse.ErrorCode)
	}

	return pricesResponse.Data, resp, nil
}
