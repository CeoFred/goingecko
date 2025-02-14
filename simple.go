package goingecko

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/CeoFred/goingecko/simple"
)

func (c *Client) SimplePrice(ids, vsCurrencies string, includeMarketCap, includeDayVolume, includeDayChange, includeLastTimeUpdated bool) (simple.Price, error) {
	params := url.Values{}

	params.Add("ids", ids)
	params.Add("vs_currencies", vsCurrencies)
	params.Add("include_market_cap", strconv.FormatBool(includeMarketCap))
	params.Add("include_24hr_vol", strconv.FormatBool(includeDayVolume))
	params.Add("include_24hr_change", strconv.FormatBool(includeDayChange))
	params.Add("include_last_updated_at", strconv.FormatBool(includeLastTimeUpdated))

	rUrl := fmt.Sprintf("%s/%s?%s", simpleURL, "price", params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data simple.Price
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) SimpleTokenPrice(id, contractAddresses, vsCurrencies string, includeMarketCap, includeDayVolume, includeDayChange, includeLastTimeUpdated bool) (simple.TokenPrice, error) {
	params := url.Values{}

	params.Add("contract_addresses", contractAddresses)
	params.Add("vs_currencies", vsCurrencies)
	params.Add("include_market_cap", strconv.FormatBool(includeMarketCap))
	params.Add("include_24hr_vol", strconv.FormatBool(includeDayVolume))
	params.Add("include_24hr_change", strconv.FormatBool(includeDayChange))
	params.Add("include_last_updated_at", strconv.FormatBool(includeLastTimeUpdated))

	rUrl := fmt.Sprintf("%s/%s/%s?%s", simpleURL, "token_price", id, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data simple.TokenPrice
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) SimpleSupportedVsCurrency() (*simple.SupportedVsCurrency, error) {
	rUrl := fmt.Sprintf("%s/%s", simpleURL, "supported_vs_currencies")
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *simple.SupportedVsCurrency
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
