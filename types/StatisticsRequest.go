package types

import "time"

type StatisticsRequest struct {
	AgentId   string             `json:"agent_id"`
	StartTime time.Time          `json:"start_time"`
	EndTime   time.Time          `json:"end_time"`
	Interval  StatisticsInterval `json:"interval"`
}

type StatisticsInterval int

const (
	EVERY_MINUTE         StatisticsInterval = 0
	EVERY_FIVE_MINUTE    StatisticsInterval = 1
	EVERY_TEN_MINUTE     StatisticsInterval = 2
	EVERY_FIFTEEN_MINUTE StatisticsInterval = 3
	EVERY_HOUR           StatisticsInterval = 4
	EVERY_DAY            StatisticsInterval = 5
)
