#!/bin/bash

go mod vendor
go mod tidy

go build
