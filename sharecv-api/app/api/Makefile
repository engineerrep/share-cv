.PHONY: all build run gotool install release swagger image clean
# 构建二进制名称
BIN_NAME=sharecvapi
# 构建根目录
BIN_FILE=$(pwd)../../deploy
GO_FILE:=core.go
API_FILE:=core.api
YAML_FILE:=core-api.yaml
YAML_FILE_PRE:=core-pre.yaml

all: install swagger

# 构建
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BIN_NAME} ${GO_FILE}

# 运行
run:
	@go run ./

# 代码格式化
format:
	go fmt ./

# 代码检测
check:
	go vet ./

# 测试环境
install:
	make build
	mv ${BIN_NAME} ${BIN_FILE}
	cp etc/${YAML_FILE_PRE} ${BIN_FILE}/etc/${YAML_FILE}

# 发布环境
release:
	make build
	mv ${BIN_NAME} ${BIN_FILE}
	cp etc/${YAML_FILE} ${BIN_FILE}/etc/${YAML_FILE}

# 生成文档
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename core.json -host http://sharecvapi.xxxxxx.com" -api ${API_FILE} -dir ${BIN_FILE}/doc/api

# 生成api
api:
	goctl api go -api ${API_FILE} -dir . -style goZero

# 生成模型
model:
	goctl model mysql ddl -src="${BIN_FILE}/doc/sql/*.sql" -dir="../model" -c -style goZero

# 清除二进制
clean:
	@if [ -f ${BIN_NAME} ] ; then rm ${BIN_NAME} ; fi

# 构建Docker镜像
image:
	docker build -t lab/${BIN_NAME}:latest .