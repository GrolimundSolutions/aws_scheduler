#!/usr/bin/env bash


# insert into table_schedule (dbid, type, "day", "hour", "action", id)
# values ();


for day in {1..7}
do
  for hour in {0..23}
  do
    g=$[$hour%2==0];
    if [ $g -eq 1 ]; then
      # echo "gerade"
      echo "insert into table_schedule (dbid, type, day, hour, action) values ('TEST-Cluster', 'cluster', $day, $hour, 'stop');"
      echo "insert into table_schedule (dbid, type, day, hour, action) values ('TEST-DB', 'db', $day, $hour, 'start');"
    else
      # echo "ungerade"
      echo "insert into table_schedule (dbid, type, day, hour, action) values ('TEST-Cluster', 'cluster', $day, $hour, 'start');"
      echo "insert into table_schedule (dbid, type, day, hour, action) values ('TEST-DB', 'db', $day, $hour, 'stop');"
    fi
  done
done