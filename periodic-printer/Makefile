# shell variables
DATE := $(shell date -u +"%Y-%m-%d")
PWD := $(shell pwd)
DISTRO := $(shell lsb_release -cs)

# Get the build/release version
ifdef PKG_VERSION
PKG_VERSION := $(PKG_VERSION)
else
PKG_VERSION := 0.0.1-0
endif

# program name
PERIODIC_PRINTER = periodicprinter

# directory containing the final debian file (this can be relative or absolute, default is '..')
DEB_OUT_DIR = dist

# Package Info
PKG_USER = pprintuser
PKG_NAME = periodicprinterpkg

# Install locations
PREF_PKG_LOG = /var/log/periodicprinter

# exports
export PERIODIC_PRINTER
export PREF_PKG_LOG
export PKG_USER
export DEB_OUT_DIR

# Makefile
.PHONY: all clean format lint ${PERIODIC_PRINTER} install uninstall deb deb-clean deb-changelog-create deb-build deb-clean-changelog
all: clean ${PERIODIC_PRINTER}

clean: deb-clean
	-rm -f build/${PERIODIC_PRINTER}


${PERIODIC_PRINTER}:
	go build -ldflags "-X main.date='${DATE}'" -o build/${PERIODIC_PRINTER} ./cmd/pprinter/pprinter.go

lint:
	golint ./...

format:
	goimports -l -w .

# install is called by the package builder
install: ${PERIODIC_PRINTER}
	install -v -d $(DESTDIR)$(prefix)
	install -v -d $(DESTDIR)$(prefix)/bin
	install -v -m 750 build/${PERIODIC_PRINTER} $(DESTDIR)$(prefix)/bin
# logs directory
	install -v -d $(DESTDIR)${PREF_PKG_LOG}

uninstall:
	-rm -f $(DESTDIR)$(prefix)/bin/${PERIODIC_PRINTER}
	-rm -f $(DESTDIR)$(PREF_PKG_LOG)

# debian building
# note: we create a deb changelog here because it's an example. Ideally, the changelog file should be only incremented with `dch -i`
deb: deb-clean deb-changelog-create deb-build

deb-clean:
	-rm -f dist/*.deb
	-rm -rf debian/${PKG_NAME}
	-rm -rf debian/.debhelper/

deb-clean-changelog:
	-rm -f debian/changelog

deb-changelog-create: deb-clean-changelog
	dch --create --package ${PKG_NAME} -D ${DISTRO} -v ${PKG_VERSION} "Release: ${PKG_VERSION}"

deb-build:
# check build deps
	dpkg-checkbuilddeps
# build
# --buildinfo-option is passed to dpkg-genbuildinfo
# --changes-option is passed to dpkg-genchanges
#	dpkg-buildpackage --unsigned-source --unsigned-changes --buildinfo-option="-u${DEB_OUT_DIR}" --changes-option="-u${DEB_OUT_DIR}"
	dpkg-buildpackage --unsigned-source --unsigned-changes

buildcontainer:
	docker build -t ppbuilder .
	docker run --rm --volume "${PWD}:/pkgbuilddir" --workdir="/pkgbuilddir" -it ppbuilder
