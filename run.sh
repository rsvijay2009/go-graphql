#!/usr/bin/env bash
export PORT=8080
export MYSQL_URL="root:root@12345@tcp(127.0.0.1:3306)/graphql?parseTime=true&sql_mode=ansi"
go run ./server/server.go