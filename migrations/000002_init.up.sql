create table if not exists monitoring.agents
(
    id       UUID,
    ip String,
    port String,
    name String,
    schedule String
) engine = MergeTree
    primary key (id);