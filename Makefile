all:
	@rm -rf bin/*
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/ifcli cmd/cli/main.go
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/ifcli.exe cmd/cli/main.go
	@GOOS=darwin GOARCH=amd64 go build -o bin/mac/ifcli cmd/cli/main.go
	@tree -Csh bin
