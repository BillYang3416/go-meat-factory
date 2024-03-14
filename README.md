# Go Meat Factory

使用 Golang 和 Concurrency 的概念撰寫該程式。

## 如何啟動程式

- clone repository
  
  ```shell
  git clone https://github.com/BillYang3416/go-meat-factory.git
  ```

- 開啟 terminal 並 cd 到該目錄

  ```shell
  cd go-meat-factory
  ```

- 執行程式

  - 確認是否有安裝 [Go](https://go.dev/doc/install)，如有則可執行 go run
    
    ```shell
    go version
    ```

    ```shell
    go run cmd/go-meat-factory/main.go
    ```

  - 直接執行
    
    - linux

      ```shell
       ./bin/linux/go-meat 
      ```

    - windows

      ```shell
       .\bin\windows\go-meat.exe
      ```

    - macOS
  
      ```shell
       ./bin/darwin/go-meat 
      ```


## 如何建置程式

### Linux

```shell
  GOOS=linux GOARCH=amd64 go build -o bin/linux/go-meat cmd/go-meat-factory/main.go
```

### Windows

```shell
  GOOS=windows GOARCH=amd64 go build -o bin/windows/go-meat.exe cmd/go-meat-factory/main.go
```

### MacOS

```shell
  GOOS=darwin GOARCH=amd64 go build -o bin/darwin/go-meat cmd/go-meat-factory/main.go
```