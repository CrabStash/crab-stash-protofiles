BIN_DIR = bin
PROTO_DIR = proto

ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
	OS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
	CHECK_DIR_CMD = if (!(Test-Path $@)) { $$e = [char]27; Write-Error "$$e[31mDirectory $@ doesn't exist$${e}[0m" }
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
	RM_RF_CMD = ${RM_F_CMD} -Recurse
else
	SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS = macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		OS = linux ${VERSION_AND_ARCH}
	else
		$(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
endif

.DEFAULT_GOAL := help
.PHONY: auth warehouse user core
project := auth warehouse user core

all: $(project) ## Generate Pbs and build

auth: $@ ## Generate Pbs and build for auth

warehouse: $@ ## Generate Pbs and build for warehouse

user: $@ ## Generate Pbs and build for user

core: $@ ## Generate Pbs and build for core

$(project):
	@${CHECK_DIR_CMD}
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto

clean_auth: ## Clean generated files for auth
	${RM_F_CMD} auth/${PROTO_DIR}/*.pb.go

clean_warehouse: ## Clean generated files for warehouse
	${RM_F_CMD} warehouse/${PROTO_DIR}/*.pb.go

clean_user: ## Clean generated files for user
	${RM_F_CMD} user/${PROTO_DIR}/*.pb.go

clean_user: ## Clean generated files for core
	${RM_F_CMD} core/${PROTO_DIR}/*.pb.go