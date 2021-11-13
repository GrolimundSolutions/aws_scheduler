#!/usr/bin/env bash

docker run --rm -it --name PD_DEV_DB \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=scheduler_db \
  -v /${PWD}/internal/pg_data:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres:alpine