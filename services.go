package sdkdemo

import (
	"context"
	"fmt"
	"net/http"
)

type CreateOrderRequest struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type OrderResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// CreateOrder 必须接受 context.Context 以支持超时和撤消
func (c *Client) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*OrderResponse, error) {
	url := fmt.Sprintf("%s/v1/orders", c.config.Endpoint)
	var res OrderResponse
	err := c.do(ctx, http.MethodPost, url, req, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) do(ctx context.Context, method, url string, body interface{}, result interface{}) error {
	return nil
}
