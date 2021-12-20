package types

import "time"

type Statistics struct {
	AgentId  string
	Datetime time.Time
	VmStat   VmStat
	Disk     Disk
	Cpu      Cpu
	Host     Host
}
type Host struct {
	Procs           int64
	OS              string
	PlatformVersion string
	Platform        string
}
type Cpu struct {
	Percentage []float64
	Model      string
	Cores      int
}
type Disk struct {
	Total       int64
	Free        int64
	Used        int64
	UsedPercent float64
}

type VmStat struct {
	Total       int64
	Free        int64
	UsedPercent float64
}
