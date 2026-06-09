PLUGIN_NAME ?= goserver04rel64

# VC:MP server release tags: rel32, rel64, etc.
# Build on the same OS/arch as your server (usually Linux x86_64 or Windows).
.PHONY: build build-linux build-windows clean

build:
	go build -buildmode=c-shared -o $(PLUGIN_NAME).so .

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o $(PLUGIN_NAME).so .

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o $(PLUGIN_NAME).dll .

clean:
	rm -f $(PLUGIN_NAME).so $(PLUGIN_NAME).dll $(PLUGIN_NAME).h
