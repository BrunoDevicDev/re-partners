package service

import (
	"re/internal/algorithm"
	"sync"
)

// Service is gonna be used to handle requests, and store currently used packs in memory
type Service struct {
	packSizes []int
	mutex     sync.Mutex
}

func NewService() *Service {
	return &Service{
		packSizes: []int{250, 500, 1000, 2000, 5000},
	}
}

func (s *Service) EditPackSizes(packSizes []int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.packSizes = packSizes
}

func (s *Service) GetPackSizes() []int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.packSizes
}

func (s *Service) SolveAlgorithm(target int) map[int]int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ret := algorithm.Solve(target, s.packSizes)
	return ret.M
}
