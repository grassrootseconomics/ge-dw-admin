BIN := ge-dw-admin

.PHONY: build
build:
	CGO_ENABLED=1 GOOS=linux go build -o ${BIN} -ldflags="-s -w" cmd/*.go