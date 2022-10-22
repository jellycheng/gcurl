.PHONY: all build run clean help go_build

# 定义变量
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)

fmt:
	@gofmt -l -w $(GO_FILES)

#显示命令帮助，如 make help
help:
	@echo "make fmt - 代码格式化，统一代码风格"

