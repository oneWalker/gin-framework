1.5之后不需要指定GOPRIVATE
#export GO111MODULE ?= on
#export GOPROXY ?= https://goproxy.cn
#export GOPRIVATE ?= `${gitAddress}`

#一般将启动相关的放进cmd文件夹

CMD_DIR := ./cobra

default: start

start:
	go run $(CMD_DIR)/main.go

start-dev:
	go run $(CMD_DIR)/main.go -c ./viperconf/config.dev.yaml

build:
	go build -mod=vendor $(CMD_DIR)/main.go


.PHONY: default build start