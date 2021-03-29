GC=go
CLI_GO_FILE=publish_cli
RELEASE_DIR=release

build :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${CLI_GO_FILE} ${CLI_GO_FILE}.go

release :
	make build
	upx release/*

clean :
	@if [ -d ${RELEASE_DIR} ] ; then rm -rf ${RELEASE_DIR}; fi