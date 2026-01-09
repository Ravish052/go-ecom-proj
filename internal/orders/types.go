package orders

import (
	"context"

	repo "github.com/Ravish052/goEcon/internal/adapters/postgres/sqlc"
)

type orderItem struct {
	ProductId int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerId int64       `json:"customer_id"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
