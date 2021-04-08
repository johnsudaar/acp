ifndef VERSION
override VERSION=dev
endif

.PHONY: all clean front config pkg-windows

all: server front config

server:
	go build -ldflags="-X main.Version=$(VERSION)" .

front:
	cd front; yarn install; yarn build

config:
	go run build/generate_config.go

clean:
	rm -rf acp acp.yml front/dist/

install:
	mkdir -p $(DESTDIR)/usr/bin/
	mkdir -p $(DESTDIR)/var/lib/acp
	mkdir -p $(DESTDIR)/etc/acp
	mkdir -p $(DESTDIR)/etc/init.d
	cp acp $(DESTDIR)/usr/bin/acp
	cp -r front/dist $(DESTDIR)/var/lib/acp/front
	cp acp.yml $(DESTDIR)/etc/acp/acp.yml
	cp build/acp.init.d $(DESTDIR)/etc/init.d/acp

pkg-windows:
	mkdir dist/windows/
	cp -R front/dist/ dist/windows/front
	cp ./acp.exe dist/windows
	cp ./acp.yml dist/windows/acp.yml