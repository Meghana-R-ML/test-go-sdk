SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
# .DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

GO_BUILD_FLAGS :=

GO_MODULE := github.com/meghana-rajashekar/test-go-sdk

intersight-go-example: $(shell find . -name \*.go -type f) go.mod Makefile
> go build -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)


clean:
> rm intersight-go-example
.PHONY: clean

