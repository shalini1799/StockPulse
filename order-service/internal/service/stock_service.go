package service

import (
	"stockpulse/internal/model"
	"stockpulse/internal/repository"
)

type StockService struct {
	repo *repository.StockRepository
}

func NewStockService(repo *repository.StockRepository) *StockService {
	return &StockService{repo: repo}
}

func (s *StockService) AddStock(stock model.Stock) {
	s.repo.Save(stock)
}

func (s *StockService) ListStocks() []model.Stock {
	return s.repo.GetAll()
}