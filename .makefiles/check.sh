#!/usr/bin/env bash

# 检查Go环境是否正常
if ! go version > /dev/null 2>&1; then
    echo "没有检测Golang环境"
fi

