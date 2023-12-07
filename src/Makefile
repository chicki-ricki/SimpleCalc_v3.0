BINARY_NAME=smartCalc
VERSION=3.0
BUILDPATH=./build/bin/
DEBDIR=./build/packages/temp/DEB

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
	go build -o ${BUILDPATH}${BINARY_NAME} ./cmd/smartCalc/main.go
	@echo "\033[0;32mbinary file builded in ${BUILDPATH}\033[0m"

build-linux: building-linux

building-linux:
	@rm -f ./build/bin/${BINARY_NAME}-linux
	@echo "\033[0;34m./build/bin/${BINARY_NAME}-linux was removed\033[0m"
	@GOARCH=amd64 GOOS=linux go build -o ${BUILDPATH}${BINARY_NAME}-linux ./cmd/clevercalc/main.go
	@echo "\033[0;36mbinary file for Linux builded in ${BUILDPATH}\033[0m"

build-darwin: building-darwin

building-darwin:
	@rm -f ./build/bin/${BINARY_NAME}-darwin
	@echo "\033[0;34m./build/bin/${BINARY_NAME}-darwin was removed\033[0m"
	GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/clevercalc/main.go
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

uninstallMacOs:
	rm -rf "/Users/$$(whoami)/Desktop/smartCalc.app" "/Applications/smartCalc.app"

# OTHER

help:
	@echo "\033[0;36m\nCommand list for this Makefile:\n"
	@echo "build\t-\tbuilding binary file in ${BUILDPATH}"
	@echo "run\t-\trunning programm with go run"
	@echo "test\t-\ttesting test with go test ./... -v -cover"
	@echo "bin\t-\trunning binary file from ${BUILDPATH}"
	@echo "clean\t-\tcleaning all binary files in ${BUILDPATH}"
	@echo "tobin\t-\tcopying binary file to /usr/bin"
	@echo "todeb\t-\tcopying binary file to ${DEBPACKAGE}/usr/bin"
	@echo "debcreate\t-\tcreate .deb package from ${DEBPACKAGE}\n\033[0m"