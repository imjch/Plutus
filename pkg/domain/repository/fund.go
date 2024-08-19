package repository

import "plutus.io/plutus/pkg/domain/entity"

type FundRepository interface {
	Get([]string) (map[string]*entity.Fund, error)
}

var fr FundRepository

func GetFundRepository() FundRepository {
	return fr
}

func InitFundRepository(f FundRepository) {
	fr = f
}
