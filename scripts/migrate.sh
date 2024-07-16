#!/bin/bash

goose -dir ./migrations postgres "postgres://postgres:password@localhost:5432/wb?sslmode=disable" status

goose -dir ./migrations postgres "postgres://postgres:password@localhost:5432/wb?sslmode=disable" up
