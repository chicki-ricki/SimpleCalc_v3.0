BINARY_NAME=smartCalc
VERSION=3.0
DATE := $(shell date +_%d-%m-%Y_%H:%M)

BUILDPATH=./build/bin/
DEB_TEMPDIR=./build/packages/temp/DEB

DEB_PACKAGEDIR = ./build/packages/done/deb/
APP_PACKAGEDIR = ./build/packages/done/app/
DMG_PACKAGEDIR = ./build/packages/done/dmg/

# RUN COMMANDS
name:
	@echo ${BINARY_NAME}-${VERSION}

run:
	go run ./cmd/smartCalc/main.go 

test: testing

testing:
	go test ./... -v -cover

# BUILDING COMMANDS

build: building

# Uncomment required string for building on you machine, if necessary
building:
#	GOARCH=arm GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/smartCalc/main.go
#	GOARCH=amd64 GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/smartCalc/main.go
#	GOARCH=amd64 GOOS=linux go build -o ${BUILDPATH}${BINARY_NAME}-linux ./cmd/smartCalc/main.go
#	GOARCH=amd64 GOOS=windows go build -o ./build/bin/${BINARY_NAME}-windows ./cmd/smartCalc/main.go
	$(eval start := $(shell date +%s))
	go build -o ${BUILDPATH}${BINARY_NAME} ./cmd/smartCalc/main.go
	@echo "\033[0;32mbinary file builded in ${BUILDPATH}\033[0m"
	$(eval end := $(shell date +%s))
	@echo "" | awk '{ print "building time:",$(start) - $(end), "seconds"}'
#	@echo "\033[0m"

build-linux: building-linux

building-linux:
	@rm -f ./build/bin/${BINARY_NAME}-linux
	@echo "\033[0;34m./build/bin/${BINARY_NAME}-linux was removed\033[0m"
	@GOARCH=amd64 GOOS=linux go build -o ${BUILDPATH}${BINARY_NAME}-linux ./cmd/smartCalc/main.go
	@echo "\033[0;36mbinary file for Linux builded in ${BUILDPATH}\033[0m"

build-darwin: building-darwin

building-darwin:
	@rm -f ./build/bin/${BINARY_NAME}-darwin
	@echo "\033[0;34m./build/bin/${BINARY_NAME}-darwin was removed\033[0m"
	GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/smartCalc/main.go
	@echo "\033[0;32mbinary file for macOS builded in ${BUILDPATH}\033[0m"

# CLEANING COMMANDS

# clean all created files
clean-all: clean clean-temp clean-package

# clean binaryes
clean:
	@rm -f ./build/bin/*
	@echo "\033[0;34m./build/bin folder was cleared\033[0m"

# clean finished packages
clean-package:
	@rm -Rf ./build/packages/done/app/* ./build/packages/done/dmg/* ./build/packages/done/deb/*
	@echo "\033[0;34mfinished packages was deleted\033[0m"

# clean temp files after packages building
clean-temp:
	@rm -Rf ./build/packages/temp/*
	@echo "\033[0;34mtempfolder was cleared\033[0m"

clean-zip:
	@rm -Rf ./zip/*.tar.gz
	@echo "\033[0;34marchive folder was cleared\033[0m"

# PACKAGE CREATE COMMANDS

deb: building-linux
# Remove old and create new deb-package folder
	@rm -Rf ${DEBDIR} ./build/done/{BINARY_NAME}-${VERSION}_amd64.deb
	@echo "\033[0;34m${DEBDIR} was removed\033[0m"
	@echo "\033[0;34m./build/done/${BINARY_NAME}-${VERSION}_amd64.deb was removed\033[0m"
	@mkdir ${DEBDIR}
# Copy data from .deb template to deb-package folder
	@cp -R ./build/packages/templates/template_deb/* ${DEBDIR}
	@cp ${BUILDPATH}${BINARY_NAME}-linux ${DEBDIR}/usr/bin/${BINARY_NAME}
	@echo "\033[0;32mtemp directory for debian package created in ${DEBDIR} \033[0m"
# Run command to compose .deb packages
	@dpkg-deb --build ${DEBDIR} ./build/packages/done/${BINARY_NAME}-${VERSION}_amd64.deb
	@echo "\033[0;36m${BINARY_NAME}-${VERSION}_amd64.deb created in ./build/packages/done/\033[0m"
# Remove temp files
	@rm -Rf ${DEBDIR}
	@echo "\033[0;34mTempfile was removed\033[0m"

app:	build
	rm -R -f ./build/packages/done/app/smartCalc.app
	cp ./build/bin/smartCalc .
	cp ./build/packages/templates/template_app/run.command .
	cp ./build/packages/templates/template_app/appMaker .
	./appMaker run.command smartCalc
	rm -f run.command smartCalc appMaker

dmgMacOS:	app
	./build/packages/templates/template_app/dmgMaker  -c 350:160:200:160 "./build/packages/done/app/"
	mv smartCalcApp.dmg ./build/packages/done/dmg/

installMacOS: app
	make clean-temp
	./build/packages/templates/template_app/installer
	make clean-temp

uninstallMacOs:
	rm -rf "/Users/$$(whoami)/Desktop/smartCalc.app" "/Applications/smartCalc.app"

# OTHER

help:
	@echo "\033[0;32m\nRUN COMMANDS\033[0m"
	@echo "\033[0;36mname\t\t-\tshow name with version"
	@echo "run\t\t-\trunning programm with go run"
	@echo "test\t\t-\ttesting test with go test ./... -v -cover\033[0m"

	@echo "\033[0;32m\nBUILD COMMANDS\033[0m"
	@echo "\033[0;36mbuild\t\t-\tbuilding binary file in ${BUILDPATH} in accordance with your machine"
	@echo "build-linux\t-\tbuilding binary file for Linux in ${BUILDPATH}"
	@echo "build-darwin\t-\tbuilding binary file for macOS in ${BUILDPATH}\033[0m"

	@echo "\033[0;32m\nPACKAGE CREATE COMMANDS\033[0m"
	@echo "\033[0;36mdeb\t\t-\tcreate .deb package in ./build/packages/done/deb (required linux)"
	@echo "dmg\t\t-\tcreate .dmg package in ./build/packages/done/dmg (required macOS)"
	@echo "app\t\t-\tcreate .app package in ./build/packages/done/app\033[0m"

	@echo "\033[0;32m\nCLEAN COMMANDS\033[0m"
	@echo "\033[0;36mclean\t\t-\tcleaning all binary files in ${BUILDPATH}"
	@echo "clean-temp\t-\tcleaning temp files in ./build/packages/temp"
	@echo "clean-package\t-\tcleaning all binary files in ./build/packages/done/*"
	@echo "clean-all\t-\trun clean clean-temp clean-package"
	@echo "clean-zip\t-\tclean archive from ./zip\033[0m"

	@echo "\033[0;32m\nUN/INSTALL COMMANDS\033[0m"
	@echo "\033[0;36minstallMacOS\t-\tinstall to MacOS"
	@echo "uninstallMacOs\t-\tuninstall from MacOS\033[0m"

	@echo "\033[0;32m\nOTHER COMMANDS\033[0m"
	@echo "\033[0;36mzip\t\t-\tcreate this folder archive in ./zip with datastamp"
	@echo "tobin\t\t-\tcopying binary file to /usr/bin"
	@echo "todeb\t\t-\tcopying binary file to ${DEBPACKAGE}/usr/bin"
	@echo "help\t\t-\tshow this help\n\033[0m"

zip: clean-all
	@tar --exclude 'zip' -cvaf ./zip/${BINARY_NAME}_srs_${DATE}.tar.gz *
	@echo "\033[0;32marchive file created to ../ as ${BINARY_NAME}_srs_${DATE}.tar.gz\033[0m"