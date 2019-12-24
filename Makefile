.PHONY: update test check build run upgrade

all: check test build

update:
	@echo "更新"

test:
	@echo "测试"

check:
	@echo "检查"
	@.makefiles/check.sh

build: check deps
	@echo "编译文件"
	@.makefiles/build.sh

deps:
	@echo "下载依赖"
	@.makefiles/deps.sh

run: check deps
	@echo "运行程序"
	@.makefiles/run.sh

upgrade:
	@echo "升级"
	@.makefiles/upgrade.sh