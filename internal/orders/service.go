package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/Ravish052/goEcon/internal/adapters/postgres/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product out of stock")
)

type svc struct {
	repo repo.Querier
	db   *pgx.Conn
}

func NewService(repo repo.Querier, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {

	//validate payload

	if tempOrder.CustomerId == 0 {
		return repo.Order{}, fmt.Errorf("invalid customer id")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("order must contain at least one item")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}

	defer tx.Rollback(ctx)
	qtx := repo.New(tx)

	//create an order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerId)

	if err != nil {
		return repo.Order{}, err
	}
	// loojk for product if exists

	for _, item := range tempOrder.Items {
		product, err := s.repo.FindProductByID(ctx, item.ProductId)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNoStock
		}

		// create order item

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductId,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCents,
		})
		if err != nil {
			return repo.Order{}, err
		}
	}
	tx.Commit(ctx)
	return order, nil
}
