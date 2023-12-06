BINARY_NAME=smartCalc
VERSION=0.0.2
BUILDPATH=./build/bin/
DEBPACKAGE=./build/packages/${BINARY_NAME}-${VERSION}/debian
#GO=$(shell which go)
#GOINSTALL=$(GO) install
#GOCLEAN=$(GO) clean
USERDECKTOP=$${whoami}


run:
	go run ./cmd/clevercalc/main.go 

build: building

building:
# GOARCH=arm GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/clevercalc/main.go
# GOARCH=amd64 GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/clevercalc/main.go
# GOARCH=amd64 GOOS=darwin go build -o ./build/bin/${BINARY_NAME}-darwin ./cmd/smartCalc/main.go
	GOARCH=amd64 GOOS=darwin go build -o ./build/bin/${BINARY_NAME} ./cmd/smartCalc/main.go
# 	GOARCH=amd64 GOOS=linux go build -o ${BUILDPATH}${BINARY_NAME}-linux ./cmd/clevercalc/main.go
# #	GOARCH=amd64 GOOS=windows go build -o ./build/bin/${BINARY_NAME}-windows ./cmd/clevercalc/main.go
	@echo "\033[0;32mbinary file builded\033[0m"

name:
	@echo $(NAME)

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

test: testing

bin:
	./build/bin/clevercalc-linux

clean:
	rm -f ./build/bin/*
	@echo "\033[0;34m./build/bin folder was cleared\033[0m"

cleanpackage:
	rm -Rf ./build/bin/*

testing:
	go test ./... -v -cover

debcreate: clean building todeb
#	dpkg-deb --build ./build/packages/debian ./build/packages/done/clevercalc-0.0.2_amd64.deb
	dpkg-deb --build ${DEBPACKAGE} ./build/packages/done/${BINARY_NAME}-${VERSION}_amd64.deb
	@echo "\033[0;32m${BINARY_NAME}-${VERSION}_amd64.deb created in ./build/packages/done/\033[0m"

tobin:
	 cp ./build/bin/clevercalc-linux /usr/bin/clevercalc
	 @echo "\033[0;32mbinary file copied to /usr/bin as clevercalc\033[0m"

todeb:
	rm -f ${DEBPACKAGE}/usr/bin/clevercalc
	@echo "\033[0;34m${DEBPACKAGE}/usr/bin/clevercalc was removed\033[0m"
	cp ./build/bin/clevercalc-linux ${DEBPACKAGE}/usr/bin/clevercalc
	@echo "\033[0;32mbinary file copied to ${DEBPACKAGE}/usr/bin as clevercalc\033[0m"

app:	build
	rm -R -f ./build/packages/done/smartCalc.app
	cp ./build/bin/smartCalc .
	cp ./build/packages/macOS/run.command .
	cp ./build/packages/macOS/appMaker .
	./appMaker run.command smartCalc
	rm -f run.command smartCalc appMaker

installMacOS:	app
# # cp -r ./build/packages/done/smartCalc.app /Applications
# 	sudo hdiutil attach smartCalcApp.dmg
# #	cd /Volumes/smartCalc
# 	sudo cp -rf /Volumes/smartCalc/smartCalc.app /Applications
# #	popd
# 	hdiutil detach /dev/disk2s1
	hdiutil create "./build/packages/smartCalcApp.dmg" -ov -volname "smartCalc" -fs HFS+ -srcfolder "./build/packages/done/smartCalc.app/"
	hdiutil attach "./build/packages/smartCalcApp.dmg"
	cp -rf /Volumes/smartCalc/smartCalc.app /Applications
	ln -s "/Applications/smartCalc.app" "/Users/$$(whoami)/Desktop"
#	ln -s "$PWD/app.jar" "$HOME/Рабочий стол/app"
	hdiutil detach /dev/disk2s1

	