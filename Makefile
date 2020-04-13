ifndef VERSION
override VERSION=dev
endif

.PHONY: all clean front config pkg-image pkd-debug pkg pkg-init

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

pkg-image:
	docker build ./build/ -t acp_package_eng

pkg-debug: pkg-image
	docker run -v $(CURDIR):/data -v $(CURDIR)/pkg:/pkg --rm -it acp_package_eng bash

pkg-init: pkg-image
	docker run -e VERSION=$(VERSION) -v $(CURDIR):/data -v $(CURDIR)/pkg:/pkg --rm acp_package_eng bash -c 'bundle install'

pkg: pkg-init
	docker run -e VERSION=$(VERSION) -v $(CURDIR):/data -v $(CURDIR)/pkg:/pkg --rm acp_package_eng bash -c 'bundle exec rake build'
