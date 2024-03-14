GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=v1.0.0'" -o linux/go-meat ../cmd/go-meat-factory/main.go
GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.Version=v1.0.0'" -o windows/go-meat.exe ../cmd/go-meat-factory/main.go
GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'main.Version=v1.0.0'" -o darwin/go-meat ../cmd/go-meat-factory/main.go