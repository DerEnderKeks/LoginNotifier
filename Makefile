make: get build
get:
	go mod download
	go mod verify

build:
	go build -o bin/loginnotifier -linkshared -ldflags="-s -w"

upx:
	upx bin/*

install:
	cp bin/loginnotifier ${PREFIX}/usr/bin/
	cp init/loginnotifier.service ${PREFIX}/etc/systemd/system/

uninstall:
	rm -rf ${PREFIX}/usr/bin/loginnotifier
	rm -rf ${PREFIX}/etc/systemd/system/loginnotifier.service

clean:
	go clean
	rm -rf bin