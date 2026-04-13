package orders

import "context"

type Service interface {
	PlaceOrder(ctx context.Context) error
}
