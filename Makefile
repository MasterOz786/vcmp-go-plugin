PLUGIN_DIR ?= plugins
EXAMPLE    ?= blank
HEADER     := include/plugin.h

PLUGIN_NAME_blank  ?= goplugin04rel64
PLUGIN_NAME_safari ?= goserver04rel64
PLUGIN_NAME        ?= $(PLUGIN_NAME_$(EXAMPLE))

.PHONY: all deps build build-blank build-safari build-linux build-linux-blank build-linux-safari build-windows build-windows-blank build-windows-safari build-all example clean tidy-safari

all: build

deps:
	cd scripts && go run fetch_plugin.go

$(HEADER):
	$(MAKE) deps

build: $(HEADER)
	@mkdir -p $(PLUGIN_DIR)
	cd examples/$(EXAMPLE) && CGO_ENABLED=1 go build -buildmode=c-shared -o ../../$(PLUGIN_DIR)/$(PLUGIN_NAME).so .

build-blank:
	$(MAKE) build EXAMPLE=blank

build-safari: tidy-safari
	$(MAKE) build EXAMPLE=safari

build-linux: $(HEADER)
	@mkdir -p $(PLUGIN_DIR)
	cd examples/$(EXAMPLE) && GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o ../../$(PLUGIN_DIR)/$(PLUGIN_NAME).so .

build-linux-blank:
	$(MAKE) build-linux EXAMPLE=blank

build-linux-safari: tidy-safari
	$(MAKE) build-linux EXAMPLE=safari

build-windows: $(HEADER)
	@mkdir -p $(PLUGIN_DIR)
	cd examples/$(EXAMPLE) && GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o ../../$(PLUGIN_DIR)/$(PLUGIN_NAME).dll .

build-windows-blank:
	$(MAKE) build-windows EXAMPLE=blank

build-windows-safari: tidy-safari
	$(MAKE) build-windows EXAMPLE=safari

build-all: build-blank build-safari

example: build-blank

tidy-safari:
	cd examples/safari && go mod tidy

clean:
	rm -f $(PLUGIN_DIR)/goplugin04rel64.so $(PLUGIN_DIR)/goplugin04rel64.dll
	rm -f $(PLUGIN_DIR)/goserver04rel64.so $(PLUGIN_DIR)/goserver04rel64.dll
	rm -f goplugin04rel64.h goserver04rel64.h
