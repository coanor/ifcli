all:
	@rm -rf bin/*
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/ifcli -ldflags "-w -s" cmd/cli/main.go
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/ifcli.exe -ldflags "-w -s" cmd/cli/main.go
	@GOOS=darwin GOARCH=amd64 go build -o bin/mac/ifcli -ldflags "-w -s" cmd/cli/main.go
	@tree -Csh bin
