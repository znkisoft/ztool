# TODO find a better way to install bin file
BIN := $(GOPATH)/bin
originBin := $(BIN)/cmd
changedBin := $(BIN)/ztool

build:
	@echo BIN = $(GOPATH)/bin
	cd cmd && go install ./...
	mv $(originBin) $(changedBin)