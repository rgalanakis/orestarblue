BIN := ./orestarblue

build:
	@go build

fmt:
	@go fmt ./...

check:
	@go vet ./...

itest: build
	@FILER_ID=24090 ${BIN} temp/testinput.csv
