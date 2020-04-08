.PHONY: all clean front config pkg-image pkd-debug

all: server front config

server:
	go build .

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

pkg-image:
	docker build ./build/ -t acp_package_eng

pkg-debug: pkg-image
	docker run -v $(CURDIR):/data -v $(CURDIR)/pkg:/pkg --rm -it acp_package_eng bash
