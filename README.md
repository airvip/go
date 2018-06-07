# go
初次相遇

## 下载第三方包，由于网络原因使用 gopm

```
$ go get -v github.com/gpmgo/gopm
$ gopm get -g -v golang.org/x/tools/cmd/goimports
```


## 将下载的包安装到 GOPATH/bin 目录中

```
$ go install golang.org/x/tools/cmd/goimports
```

GOPATH下目录结构

* go build 编译
* go install 产生pkg文件和可执行文件
* go run 直接编译执行
