package fund

import (
	"context"
	"fmt"
	"sync"

	"plutus.io/plutus/pkg/domain/repository"
)

type View struct {
}

type APP interface {
	Detail(ctx context.Context, ids []string) (*View, error)
}

type appImpl struct {
	ctx  context.Context
	repo repository.FundRepository
}

var app APP
var once sync.Once

func NewFundAPP(ctx context.Context, repo repository.FundRepository) APP {
	once.Do(func() {
		app = &appImpl{
			ctx:  ctx,
			repo: repo,
		}
	})
	return app
}

func (app *appImpl) Detail(ctx context.Context, ids []string) (*View, error) {
	_, err := app.repo.Get(ids)
	if err != nil {
		return nil, fmt.Errorf("errAPPFuncDetail:%w", err)
	}
	v := &View{}
	return v, nil
}
