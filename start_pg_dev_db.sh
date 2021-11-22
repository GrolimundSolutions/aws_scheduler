#!/usr/bin/env bash

docker run --rm -it -d --name PD_DEV_DB \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=scheduler_db \
  -p 5432:5432 \
  postgres:alpine

docker run --rm -it -d --name PG_ADMIN \
  -p 8080:80 \
  -e 'PGADMIN_DEFAULT_EMAIL=user@user.com' \
  -e 'PGADMIN_DEFAULT_PASSWORD=user' \
  dpage/pgadmin4

