create table IF NOT EXISTS table_schedule
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

alter table table_schedule
    owner to postgres;
