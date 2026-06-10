PLUGIN_DIR   ?= plugins
PLUGIN_NAME  ?= goserver04rel64
HEADER       := include/plugin.h
PLUGIN_SRC   := plugin

.PHONY: all deps build build-linux build-windows clean tidy

all: build

deps:
	cd scripts && go run fetch_plugin.go

$(HEADER):
	$(MAKE) deps

tidy:
	cd $(PLUGIN_SRC) && go mod tidy

build: $(HEADER) tidy
	@mkdir -p $(PLUGIN_DIR)
	cd $(PLUGIN_SRC) && CGO_ENABLED=1 go build -buildmode=c-shared -o ../../$(PLUGIN_DIR)/$(PLUGIN_NAME).so .

build-linux: $(HEADER) tidy
	@mkdir -p $(PLUGIN_DIR)
	cd $(PLUGIN_SRC) && GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o ../../$(PLUGIN_DIR)/$(PLUGIN_NAME).so .

build-windows: $(HEADER) tidy
	@mkdir -p $(PLUGIN_DIR)
	cd $(PLUGIN_SRC) && GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o ../../$(PLUGIN_DIR)/$(PLUGIN_NAME).dll .

clean:
	rm -f $(PLUGIN_DIR)/goserver04rel64.so $(PLUGIN_DIR)/goserver04rel64.dll
	rm -f goserver04rel64.h
