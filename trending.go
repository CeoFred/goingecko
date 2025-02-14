package goingecko

import (
	"encoding/json"

	"github.com/CeoFred/goingecko/trending"
)

func (c *Client) Trending() (*trending.Trending, error) {
	rUrl := trendingURL
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *trending.Trending
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
