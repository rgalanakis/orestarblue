BIN := ./orestarblue

build:
	@go build

itest: build
	@FILER_ID=24090 ${BIN} tmp/testinput.csv
