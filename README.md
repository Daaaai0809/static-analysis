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

## ユースケース

```go
package main

type User struct {
    Id int // Fail
    Name string
}

type Path struct {
    Url string // Fail
}

func main() {
    // ...
}
```