package repository

import (
	"stockpulse/internal/model"
	"sync"
)

type StockRepository struct {
	stocks map[string]model.Stock
	mu     sync.RWMutex
}

func NewStockRepository() *StockRepository {
	return &StockRepository{
		stocks: make(map[string]model.Stock),
	}
}

func (r *StockRepository) Save(stock model.Stock) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.stocks[stock.Symbol] = stock
}

func (r *StockRepository) GetAll() []model.Stock {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []model.Stock
	for _, stock := range r.stocks {
		result = append(result, stock)
	}
	return result
}