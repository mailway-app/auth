VERSION = 1.0.0
DIST = $(PWD)/dist
FPM_ARGS =

.PHONY: clean
clean:
	rm -rf $(DIST) *.deb

$(DIST)/auth: server.go
	mkdir -p $(DIST)
	go build -o $(DIST)/usr/local/sbin/auth

.PHONY: deb
deb: $(DIST)/auth
	fpm -n auth -s dir -t deb --chdir=$(DIST) --version=$(VERSION) $(FPM_ARGS)
