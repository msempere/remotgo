.PHONY: version all run dist clean
APP_NAME := remotgo
VERSION := 0.0.1
DIR=.

# Go setup
GO=go
GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

DEPENDENCIES := github.com/aws/aws-sdk-go \
		github.com/codegangsta/cli \
		golang.org/x/crypto/ssh


# Sources and Targets
EXECUTABLES :=bin/$(APP_NAME)
# Package target
PACKAGE :=$(DIR)/dist/$(APP_NAME)-$(VERSION).tar.gz

# Custom go path for project
VENDOR_DIR :=_vendor/src
DEPENDENCIES_DIR := $(addprefix $(VENDOR_DIR),$(DEPENDENCIES))

.DEFAULT: all

all: | $(EXECUTABLES)

# print the version
version:
	@echo $(VERSION)
# print the name of the app
name:
	@echo $(APP_NAME)

# print the package path
package:
	@echo $(PACKAGE)

# We have to set GOPATH to just the _vendor
# directory to ensure that `go get` doesn't
# update packages in our primary GOPATH instead.
# This will happen if you already have the package
# installed in GOPATH since `go get` will use
# that existing location as the destination.
$(DEPENDENCIES_DIR):
	@echo Downloading $(@:$(VENDOR_DIR)%=%)...
	@GOPATH=${PWD}/_vendor go get $(@:$(VENDOR_DIR)%=%)

# This defines the req's for each binary
# if you only have one then these can be placed under
# $(EXECUTABLES): Target
# you must define the main file for each binary that will
# be placed in bin.
# For Example:
# -----------
# binary #1
# bin/$(APP_NAME): main.go pkg/* $(DEPENDENCIES_DIR)
# binary #2
# bin/my_other_bin: other.go pkg/* $(DEPENDENCIES_DIR)
bin/$(APP_NAME): main.go $(DEPENDENCIES_DIR) 

$(EXECUTABLES): 
	$(GO) build -o $@ $<

run: bin/$(APP_NAME)
	bin/$(APP_NAME)

clean:
	@echo Cleaning Workspace...
	  rm -dRf _vendor
	  rm -dRf bin
	  rm -dRf dist

$(PACKAGE): all
	@echo Packaging Binaries...
	@mkdir -p tmp/$(APP_NAME)/bin
	@cp -R bin/ tmp/$(APP_NAME)/
	@mkdir -p $(DIR)/dist/
	tar -cf $@ -C tmp $(APP_NAME);
	@rm -rf tmp

dist: $(PACKAGE)
