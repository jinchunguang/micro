#!/usr/bin/env bash

# 代理
GOPROXY=http://goproxy.io && go mod tidy -v
GOPROXY='' && go mod tidy
echo '验证依赖包'
go mod verify
