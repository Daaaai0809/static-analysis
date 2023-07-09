package findIdorUrl_test

import (
	"testing"
	"github.com/Daaaai0809/static-analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestStaticAnalysis(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, findIdorUrl.Analyzer, "a")
}