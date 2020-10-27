GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = hetzner-cloud-status

depend:
	env GOPATH=$(GOPATH) go get github.com/logrusorgru/aurora
	env GOPATH=$(GOPATH) go get github.com/olekukonko/tablewriter

build:
	env GOPATH=$(GOPATH) go install $(PROGRAMS)

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/bin

strip: build
	strip --strip-all $(BINDIR)/hetzner-cloud-status

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/hetzner-cloud-status $(DESTDIR)/usr/bin

clean:
	/bin/rm -f bin/hetzner-cloud-status

distclean: clean
	rm -rf src/github.com/
	rm -rf src/gopkg.in/
	rm -rf src/golang.org/
	test -d pkg && chmod -R u+w pkg/ && rm -rf pkg/

uninstall:
	/bin/rm -f $(DESTDIR)/usr/bin

all: depend build strip install

