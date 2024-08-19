package persistence

import (
	"context"
	"sync"

	"plutus.io/plutus/pkg/domain/entity"
	"plutus.io/plutus/pkg/infrastructure/utils"
	"plutus.io/plutus/pkg/library/resource"
)

type FundImpl struct {
	fechFunc utils.FetcherFunc
}

func NewFuncImpl(ctx context.Context, f utils.FetcherFunc) (*FundImpl, error) {
	return &FundImpl{
		fechFunc: f,
	}, nil
}

// Get - 根据基金ID获取基金详情
func (f *FundImpl) Get(ids []string) (map[string]*entity.Fund, error) {

	wg := &sync.WaitGroup{}

	syncM := sync.Map{}
	for _, item := range ids {
		wg.Add(1)
		go func(id string) {
			defer func() {
				if e := recover(); e != nil {
					//resource.Slog.Warn("errFundGetPanic:%w", e)

				}
				wg.Done()
			}()
			res := &entity.Fund{}
			err := f.fechFunc(id, &res)
			if err != nil {
				resource.Slog.Warn("errFundGetFechFunc:%w", err)
				return
			}
			syncM.Store(id, res)
		}(item)
	}
	wg.Wait()
	res := make(map[string]*entity.Fund)
	syncM.Range(func(key, value any) bool {
		k, ok := key.(string)
		if !ok {
			resource.Slog.Warn("errFundGetConvKey:%+v", key)
			return true
		}
		f, ok := value.(*entity.Fund)
		if !ok {
			resource.Slog.Warn("errFundGetConvValue:%+v", value)
			return true
		}
		res[k] = f
		return true
	})
	return res, nil
}
