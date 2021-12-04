alter table monitoring.agents add column if not exists state int;

alter table monitoring.agents comment column state '1:off 2:on 3:error';

alter table monitoring.agents update state = 2 where state = 0;