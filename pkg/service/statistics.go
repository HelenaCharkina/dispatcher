package service

import (
	"dispatcher/pkg/repository"
	"dispatcher/types"
)

type StatisticsService struct {
	repo repository.Statistics
}

func NewStatisticsService(repo repository.Statistics) *StatisticsService {
	return &StatisticsService{
		repo: repo,
	}
}

func (s *StatisticsService) Add(stats *types.Statistics) error {
	return s.repo.Add(stats)
}

func (s *StatisticsService) GetByAgentId(params *types.StatisticsRequest) (*[]types.DBStatistics, error) {
	stats, err := s.repo.GetByAgentId(params)
	if err != nil {
		return nil, err
	}
	for idx, stat := range *stats {
		(*stats)[idx].DiskTotal = stat.DiskTotal / 1000 / 1000
		(*stats)[idx].DiskFree = stat.DiskFree / 1000 / 1000
		(*stats)[idx].DiskUsed = stat.DiskUsed / 1000 / 1000
	}
	return stats, nil
}
