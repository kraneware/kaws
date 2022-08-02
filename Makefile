SHELL := /bin/bash

TEST_PACKAGES = $(shell find . -name "*_test.go" | sort | rev | cut -d'/' -f2- | rev | uniq)
CURDIR = $(shell pwd)

branchId := $(shell echo ${BRANCH_ID})
ifeq ($(branchId),)
	FMT_GITHUB_REF = $(shell git status | grep "On branch " | cut -d' ' -f3 | sed 's/[\/]/\\\//g')
	FMT_GITHUB_REPO = github.com\/kraneware\/$(shell basename `git rev-parse --show-toplevel`)
else
	FMT_GITHUB_REF = $(shell echo $$BRANCH_ID | sed 's/[\/]/\\\//g')
	FMT_GITHUB_REPO = github.com\/$(shell echo $$TRAVIS_REPO_SLUG | sed 's/[\/]/\\\//g')
endif

.DEFAULT_GOAL := test

.PHONY: clean init displavars test coverage coverage-checks build buildOnly

clean:
	@rm -Rf target
	@rm -f kriptoe
	@rm -f kriptoe.so
	@rm -Rf vendor

init: clean
	@echo Running in ${CURDIR}
	@mkdir target
	@mkdir target/testing
	@mkdir target/bin
	@mkdir target/deploy
	@mkdir target/tools

deps: init
	go env GOPRIVATE=github.com/kraneware/* GOOS=darwin build -v ./...

displayvars:
	@for package in $(TEST_PACKAGES); do \
		echo $${package:2}; \
	done

cleanup:
	gofmt -w .
	$(GOPATH)/bin/goimports -w .

linter: deps cleanup
	@target/tools/golangci-lint run --timeout 1h --enable-all --disable=typecheck

test: init
	@for package in $(TEST_PACKAGES); do \
	  echo Testing package $$package ; \
	  cd $(CURDIR)/$$package ; \
	  mkdir ${CURDIR}/target/testing/$$package ; \
	  go test -v -race -covermode=atomic -coverprofile ${CURDIR}/target/testing/$$package/coverage.out | tee ${CURDIR}/target/testing/$$package/target.txt ; \
	  if [ "$${PIPESTATUS[0]}" -ne "0" ]; then exit 1; fi; \
	  grep "FAIL!" ${CURDIR}/target/testing/$$package/target.txt ; \
	  if [ "$$?" -ne "1" ]; then exit 1; fi; \
	  cat ${CURDIR}/target/testing/$$package/coverage.out >> ${CURDIR}/target/coverage_profile.out ; \
	done

coverage: test
	@for package in ${TEST_PACKAGES}; do \
	  export MIN_COVERAGE=95 ; \
	  echo Generating coverage report for $$package ; \
	  cd $(CURDIR)/$$package ; \
	  if [ -f test.config ]; then source test.config; fi; \
	  go tool cover -html=${CURDIR}/target/testing/$$package/coverage.out -o ${CURDIR}/target/testing/$$package/coverage.html ; \
	done

coverage-checks: coverage
	@for package in ${TEST_PACKAGES}; do \
	  export MIN_COVERAGE=100 ; \
	  cd $(CURDIR)/$$package ; \
	  if [ -f test.config ]; then source ./test.config; fi; \
	  echo Checking coverage for $$package at $$MIN_COVERAGE% ; \
	  export COVERAGE_PCT=`grep "coverage: " ${CURDIR}/target/testing/$$package/target.txt | cut -d' ' -f2` ; \
	  export COVERAGE=`echo $$COVERAGE_PCT | cut -d'.' -f1` ; \
	  if [ "$$COVERAGE" -lt "$$MIN_COVERAGE" ]; then echo - Coverage not met at $$COVERAGE_PCT. ; exit 1; fi ; \
	  echo "  Coverage passed with $$COVERAGE_PCT" ; \
	done

build: coverage-checks buildOnly

buildOnly:
	set -m # Enable Job Control | build all inner packages in loop | build lib at very end in root dir
	@for f in $(TEST_PACKAGES); do \
  		echo "processing dir $${f}"; \
  		dir=$${f}; \
		cd $${dir} && env GOPRIVATE=github.com/kraneware/* GOOS=darwin go build && cd ..; \
		if [ $$? -ne 0 ]; then \
  		  	echo "exiting on go build error: $$? "; \
  			exit 1; \
  		fi; \
	done; \
	wait < <(jobs -p); \
	echo "processing root dir (.) for executable build"; \
	env GOPRIVATE=github.com/kraneware/* GOOS=darwin go build; \
	echo "Executable build successful ;)";