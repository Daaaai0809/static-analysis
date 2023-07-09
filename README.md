# static-analysis

このツールではGoプログラムに対して静的解析を行い、Id, Urlといった固有名詞に大文字小文字が混在していないかチェックします。

## インストール

```shell
$ go get -u github.com/Daaaai0809/static-analysis
```

## go vet での使い方

```shell
$ go vet -vettool=$(which findIdorUrl) <FileName>
```

## 使用例

```go
package main

type User struct {
    Id int // Fail: IDに修正するとok
    Name string
}

type Path struct {
    Url string // Fail: URLに修正するとok
}

func main() {
    // ...
}
```