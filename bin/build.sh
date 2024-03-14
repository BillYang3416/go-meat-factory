GOOS=linux GOARCH=amd64 go build -o linux/go-meat ../cmd/go-meat-factory/main.go
GOOS=windows GOARCH=amd64 go build -o windows/go-meat.exe ../cmd/go-meat-factory/main.go
GOOS=darwin GOARCH=amd64 go build -o darwin/go-meat ../cmd/go-meat-factory/main.go