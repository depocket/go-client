package client

import (
	"context"
	"fmt"
	"net/http"
)

type TokenService service

type Token struct {
	Id       int64   `json:"id"`
	Address  string  `json:"address"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	SiteUrl  string  `json:"icon_url"`
	IconUrl  string  `json:"site_url"`
	Type     string  `json:"type"`
	Decimals int     `json:"decimals"`
	Chain    string  `json:"chain"`
	Price    float64 `json:"price"`
}

type DataResponse[T any] struct {
	Data      T   `json:"data"`
	ErrorCode int `json:"error_code"`
}

type TokenUpdate struct {
	Name         string `json:"name,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	Type         string `json:"type,omitempty"`
	Decimals     uint64 `json:"decimals,omitempty"`
	GenesisBlock uint64 `json:"genesis_block,omitempty"`
}

func (i Token) String() string {
	return Stringify(i)
}

type TokenListOptions struct {
	ListOptions
	Addresses []string `url:"addresses,omitempty"`
	Symbols   []string `url:"symbols,omitempty"`
	Projects  []string `url:"projects,omitempty"`
}

func (s *TokenService) List(ctx context.Context, opts *TokenListOptions) ([]*Token, *Response, error) {
	var u = "tokens"
	return s.listTokens(ctx, u, opts)
}

func (s *TokenService) Genesis(ctx context.Context, opts *TokenListOptions) (map[string]uint64, *Response, error) {
	var u = "genesis/tokens"

	opts.Addresses = ConvertArrayOptsToApiParam(opts.Addresses)
	opts.Symbols = ConvertArrayOptsToApiParam(opts.Symbols)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var response *DataResponse[map[string]uint64]
	resp, err := s.client.Do(ctx, req, &response)
	if err != nil {
		return nil, resp, err
	}
	if response.ErrorCode != 0 {
		return nil, resp, fmt.Errorf("error code %d", response.ErrorCode)
	}

	return response.Data, resp, nil
}

func (s *TokenService) Update(ctx context.Context, chain, address string, tokenUpdate TokenUpdate) error {
	url := fmt.Sprintf("token/%s?chain=%s", address, chain)
	req, err := s.client.NewRequest(http.MethodPatch, url, tokenUpdate)
	if err != nil {
		return err
	}

	var response *struct {
		Message string `json:"message"`
	}
	resp, err := s.client.Do(ctx, req, &response)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusInternalServerError {
			return fmt.Errorf("error with message %s", response.Message)
		}
		return fmt.Errorf("error response code %d", resp.StatusCode)
	}
	return nil
}

func (s *TokenService) listTokens(ctx context.Context, u string, opts *TokenListOptions) ([]*Token, *Response, error) {
	opts.Addresses = ConvertArrayOptsToApiParam(opts.Addresses)
	opts.Symbols = ConvertArrayOptsToApiParam(opts.Symbols)
	opts.Projects = ConvertArrayOptsToApiParam(opts.Projects)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var tokensResponse *DataResponse[[]*Token]
	resp, err := s.client.Do(ctx, req, &tokensResponse)
	if err != nil {
		return nil, resp, err
	}
	if tokensResponse.ErrorCode != 0 {
		return nil, resp, fmt.Errorf("error code %d", tokensResponse.ErrorCode)
	}

	return tokensResponse.Data, resp, nil
}
