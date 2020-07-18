# https://danishpraka.sh/2019/12/07/using-makefiles-for-go.html
GO111MODULES=on

APP=monkey_repl
SRCS = main.go

.PHONY: build clean run help
## build: build the application
build:
	go build -o ${APP} ${SRCS}

## run: runs go run main.go
run:
	go run -race ${SRCS}

## clean: cleans the binary
clean:
	@go clean
	@if sh -c "ls *~ > /dev/null 2>&1"; then rm -f *~; fi

## test: run unit test
test:
	go test ./lexer

## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor


## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
