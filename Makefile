default:
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -v -o bin/ipcalc_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -v -o bin/ipcalc_linux_amd64
