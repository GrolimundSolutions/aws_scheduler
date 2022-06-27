BEGIN;
CREATE SCHEMA IF NOT EXISTS rds_scheduler;
create table IF NOT EXISTS rds_scheduler.table_schedule
(
    id     serial
        constraint table_schedule_pk
            primary key,
    dbid   varchar,
    type   varchar,
    day    integer,
    hour   integer,
    action varchar
);
COMMIT;