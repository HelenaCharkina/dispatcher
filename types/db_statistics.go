package types

import "time"

type DBStatistics struct {
	Id              string    `json:"id" db:"Id"`
	AgentId         string    `json:"agent_id" db:"Agent"`
	Datetime        time.Time `json:"datetime" db:"Datetime"`
	VmTotal         uint64    `json:"vm_total" db:"VmTotal"`
	VmFree          uint64    `json:"vm_free" db:"VmFree"`
	VmUsedPercent   float64   `json:"vm_used_percent" db:"VmUsedPercent"`
	DiskTotal       uint64    `json:"disk_total" db:"DiskTotal"`
	DiskFree        uint64    `json:"disk_free" db:"DiskFree"`
	DiskUsed        uint64    `json:"disk_used" db:"DiskUsed"`
	DiskUsedPercent float64   `json:"disk_used_percent" db:"DiskUsedPercent"`
	CpuPercentage   float64   `json:"cpu_percentage" db:"CpuPercentage"`
	HostProcs       float64   `json:"host_procs" db:"HostProcs"`
	OS              string    `json:"os" db:"OS"`
	PlatformVersion string    `json:"platform_version" db:"PlatformVersion"`
	Platform        string    `json:"platform" db:"Platform"`
	Model           string    `json:"model" db:"Model"`
	Cores           int       `json:"cores" db:"Cores"`
}
