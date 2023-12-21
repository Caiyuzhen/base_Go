## 编译文件
- go build hello.go

## 执行文件
- go run hello.go


## 注意
- 定义的变量或者导入的包没有使用就不能编译通过


## 要导出 utils 包的, 需要在 testPkg 内运行:


```bash
go mod init testPkg


