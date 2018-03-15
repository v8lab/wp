SRC = $(wildcard *.go)

export GOPATH=/home/shijiaotong:$(shell pwd)

export GOBIN=$(shell pwd)/bin

ws:$(SRC)
	rm -rf ./pkg/*
	go install -x ./src/main $^
	mv ./bin/main ./bin/redis 
 
	@echo
	@echo "请确认，当前 gopath 是：" $(GOPATH)
	@echo

 
