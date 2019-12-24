#!/usr/bin/env bash

# 代理
GOPROXY=https://goproxy.io,direct
go mod tidy -v
echo '验证依赖包'
go mod verify
go mod vendor