create table if not exists monitoring.stats
(
    Id              UUID,
    AgentId         UUID,
    Datetime        DateTime,
    VmTotal         Int64,
    VmFree          Int64,
    VmUsedPercent   Float64,
    DiskTotal       Int64,
    DiskFree        Int64,
    DiskUsed        Int64,
    DiskUsedPercent Float64,
    CpuPercentage   Float64,
    HostProcs       Int64

) engine = MergeTree
      primary key (Id);