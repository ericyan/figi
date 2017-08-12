package figi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type MappingRequest struct {
	IDType       string `json:"idType"`
	IDValue      string `json:"idValue"`
	ExchangeCode string `json:"exchCode,omitempty"`
	MIC          string `json:"micCode,omitempty"`
	Currency     string `json:"currency,omitempty"`
	MarketSector string `json:"marketSecDes,omitempty"`
}

type MappingResponse struct {
	Error string
	Data  []struct {
		FIGI                 string `json:"figi"`
		SecurityType         string `json:"securityType"`
		MarketSector         string `json:"marketSector"`
		Ticker               string `json:"ticker"`
		Name                 string `json:"name"`
		UniqueID             string `json:"uniqueID"`
		ExchangeCode         string `json:"exchCode"`
		ShareClassFIGI       string `json:"shareClassFIGI"`
		CompositeFIGI        string `json:"compositeFIGI"`
		SecurityType2        string `json:"securityType2"`
		SecurityDescription  string `json:"securityDescription"`
		UniqueIDFutureOption string `json:"uniqueIDFutOpt"`
	}
}

type Client struct {
	APIKey string
	*http.Client
}

func NewClient() *Client {
	return &Client{"", http.DefaultClient}
}

func (c *Client) SetAPIKey(key string) {
	c.APIKey = key
}

func (c *Client) Query(q []MappingRequest) ([]MappingResponse, error) {
	body, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openfigi.com/v1/mapping", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "text/json")
	if c.APIKey != "" {
		req.Header.Add("X-OPENFIGI-APIKEY", c.APIKey)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var results []MappingResponse
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return nil, err
	}

	return results, err
}
