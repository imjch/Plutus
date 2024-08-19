package fund

import (
	"context"
	"fmt"

	"plutus.io/plutus/pkg/application/fund"
)

type Handler interface {
	Detail(context.Context, []string) (any, error)
}

type handler struct {
	app fund.APP
}

func NewHandler(ctx context.Context, app fund.APP) Handler {
	return &handler{
		app: app,
	}
}

func (h *handler) Detail(ctx context.Context, ids []string) (any, error) {
	c, err := h.app.Detail(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("errHandlerDetail:%w", err)
	}
	return c, nil
}
