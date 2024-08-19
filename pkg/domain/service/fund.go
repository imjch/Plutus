package service

import (
	"context"
	"fmt"

	"plutus.io/plutus/pkg/domain/entity"
	"plutus.io/plutus/pkg/domain/repository"
)

type Service struct {
	repo repository.FundRepository
}

func InitService(ctx context.Context, repo repository.FundRepository) (*Service, error) {
	return &Service{repo: repo}, nil
}

func (s *Service) GetFunds(ctx context.Context, ids []string) (map[string]*entity.Fund, error) {
	fundM, err := s.repo.Get(ids)
	if err != nil {
		return nil, fmt.Errorf("errServiceFund:%w", err)
	}
	return fundM, nil
}
