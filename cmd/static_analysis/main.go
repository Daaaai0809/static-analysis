package main

import (
	"github.com/Daaaai0809/static-analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(findIdorUrl.Analyzer)
}