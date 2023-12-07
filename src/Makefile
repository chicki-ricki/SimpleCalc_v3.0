BINARY_NAME=smartCalc
VERSION=3.0
BUILDPATH=./build/bin/
# DEBPACKAGE=./build/packages/${BINARY_NAME}-${VERSION}/debian
DEBDIR=./build/packages/temp/DEB
#GO=$(shell which go)
#GOINSTALL=$(GO) install
#GOCLEAN=$(GO) clean
# USERDECKTOP=$${whoami}

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

# building:
# # GOARCH=arm GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/clevercalc/main.go
# # GOARCH=amd64 GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/clevercalc/main.go
# # GOARCH=amd64 GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/smartCalc/main.go
# 	GOARCH=amd64 GOOS=darwin go build -o ./build/bin/${BINARY_NAME} ./cmd/smartCalc/main.go
# # 	GOARCH=amd64 GOOS=linux go build -o ${BUILDPATH}${BINARY_NAME}-linux ./cmd/clevercalc/main.go
# # #	GOARCH=amd64 GOOS=windows go build -o ./build/bin/${BINARY_NAME}-windows ./cmd/clevercalc/main.go
# 	@echo "\033[0;32mbinary file builded\033[0m"

# name:
# 	@echo $(NAME)

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
	@rm -Rf ./build/packages/done/app/* ./build/packages/done/dmg/*
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
	hdiutil create "./build/packages/temp/smartCalcApp.dmg" -ov -volname "smartCalc" -fs HFS+ -srcfolder "./build/packages/done/smartCalc.app/"
	hdiutil attach "./build/packages/temp/smartCalcApp.dmg"
	cp -rf /Volumes/smartCalc/smartCalc.app /Applications
	ln -s "/Applications/smartCalc.app" "/Users/$$(whoami)/Desktop"
	hdiutil detach /dev/disk2s1
	make clean-temp

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

# test: testing

# bin:
# 	./build/bin/clevercalc-linux

# clean:
# 	rm -f ./build/bin/*
# 	@echo "\033[0;34m./build/bin folder was cleared\033[0m"

# cleanpackage:
# 	rm -Rf ./build/bin/*

# # testing:
# # 	go test ./... -v -cover

# debcreate: clean building todeb
# #	dpkg-deb --build ./build/packages/debian ./build/packages/done/clevercalc-0.0.2_amd64.deb
# 	dpkg-deb --build ${DEBPACKAGE} ./build/packages/done/${BINARY_NAME}-${VERSION}_amd64.deb
# 	@echo "\033[0;32m${BINARY_NAME}-${VERSION}_amd64.deb created in ./build/packages/done/\033[0m"

# tobin:
# 	 cp ./build/bin/clevercalc-linux /usr/bin/clevercalc
# 	 @echo "\033[0;32mbinary file copied to /usr/bin as clevercalc\033[0m"

# todeb:
# 	rm -f ${DEBPACKAGE}/usr/bin/clevercalc
# 	@echo "\033[0;34m${DEBPACKAGE}/usr/bin/clevercalc was removed\033[0m"
# 	cp ./build/bin/clevercalc-linux ${DEBPACKAGE}/usr/bin/clevercalc
# 	@echo "\033[0;32mbinary file copied to ${DEBPACKAGE}/usr/bin as clevercalc\033[0m"

# app:	build
# 	rm -R -f ./build/packages/done/smartCalc.app
# 	cp ./build/bin/smartCalc .
# 	cp ./build/packages/macOS/run.command .
# 	cp ./build/packages/macOS/appMaker .
# 	./appMaker run.command smartCalc
# 	rm -f run.command smartCalc appMaker

# installMacOS:	app
# # installMacOS:
# # 	hdiutil create "./build/packages/smartCalcApp.dmg" -ov -volname "smartCalc" -fs HFS+ -srcfolder "./build/packages/done/smartCalc.app/"
# # #	hdiutil create -srcfolder "./build/packages/done/smartCalc.app/" -volname "smartCalc" -fs HFS+  -fsargs "-c c=64,a=16,e=16" -format UDRW -size ${25000}k pack.temp.dmg 
# # 	hdiutil attach "./build/packages/smartCalcApp.dmg"
# # 	mkdir "/Volumes/smartCalc/.background"
# # 	sudo chmod -Rf go-w /Volumes/smartCalc/.background
# # 	cp "./build/packages/macOS/.background/backgroundMathw4.jpg" "/Volumes/smartCalc/.background/"

# # 	cp -rf /Volumes/smartCalc/smartCalc.app /Applications
# # 	ln -s "/Applications/smartCalc.app" "/Users/$$(whoami)/Desktop"
# # 	hdiutil detach /dev/disk2s1
# 	./build/packages/macOS/dmgMaker  -c 350:160:200:160 "./build/packages/done/"

	