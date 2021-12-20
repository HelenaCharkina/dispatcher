package repository

import (
	"dispatcher/types"
	"errors"
	"github.com/jmoiron/sqlx"
	uuid "github.com/nu7hatch/gouuid"
	"time"
)

type StatisticsRepo struct {
	db *sqlx.DB
}

func NewStatisticsRepo(db *sqlx.DB) *StatisticsRepo {
	return &StatisticsRepo{
		db: db,
	}
}

func (s *StatisticsRepo) Add(stats *types.Statistics) error {
	var (
		tx, _   = s.db.Begin()
		stmt, _ = tx.Prepare("INSERT INTO monitoring.stats (Id, AgentId, Datetime, VmTotal, VmFree, VmUsedPercent, DiskTotal, DiskFree, DiskUsed, DiskUsedPercent,CpuPercentage, HostProcs , OS, Platform, PlatformVersion, Model, Cores) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	)
	defer stmt.Close()

	id, _ := uuid.NewV4()
	if _, err := stmt.Exec(
		id.String(),
		stats.AgentId,
		time.Now(),
		stats.VmStat.Total,
		stats.VmStat.Free,
		stats.VmStat.UsedPercent,
		stats.Disk.Total,
		stats.Disk.Free,
		stats.Disk.Used,
		stats.Disk.UsedPercent,
		stats.Cpu.Percentage[0],
		stats.Host.Procs,
		stats.Host.OS,
		stats.Host.Platform,
		stats.Host.PlatformVersion,
		stats.Cpu.Model,
		stats.Cpu.Cores,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *StatisticsRepo) GetByAgentId(params *types.StatisticsRequest) (*[]types.DBStatistics, error) {
	var statistics []types.DBStatistics

	interval := s.getInterval(params.Interval)
	if interval == "" {
		return nil, errors.New("Не задан интервал ")
	}

	if err := s.db.Select(&statistics,
		`SELECT max(Id) as Id,
       max(AgentId) as Agent, ` +
		interval +`(Datetime ) as Datetime,
       avg(VmUsedPercent) as VmUsedPercent,
       avg(DiskUsedPercent) as DiskUsedPercent,
       avg(CpuPercentage) as CpuPercentage,
       avg(HostProcs) as HostProcs,
		max(DiskTotal) as DiskTotal,
       max(DiskFree) as DiskFree,
       max(DiskUsed) as DiskUsed,
		max(OS) as OS,
       max(Platform) as Platform,
       max(PlatformVersion) as PlatformVersion,
	   max(Model) as Model,
       max(Cores) as Cores
		FROM monitoring.stats
		where AgentId = ? and Datetime between ? and ?
		group by Datetime
		order by Datetime`, params.AgentId, params.StartTime, params.EndTime); err != nil {
		return nil, err
	}

	return &statistics, nil
}

func (s *StatisticsRepo) getInterval(interval types.StatisticsInterval) string {
	switch interval {
	case types.EVERY_MINUTE:
		return "toStartOfMinute"
	case types.EVERY_FIVE_MINUTE:
		return "toStartOfFiveMinute"
	case types.EVERY_TEN_MINUTE:
		return "toStartOfTenMinutes"
	case types.EVERY_FIFTEEN_MINUTE:
		return "toStartOfFifteenMinutes"
	case types.EVERY_HOUR:
		return "toStartOfHour"
	case types.EVERY_DAY:
		return "toStartOfDay"
	default:
		return ""
	}
}