package findIdorUrl_test

import (
	"testing"
	"github.com/Daaaai0809/static-analysis/findIdorUrl"
	"github.com/Daaaai0809/static-analysis/findIdorUrl_testdata"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestStaticAnalysis(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, findIdorUrl.Analyzer, "findIdorUrl_testdata")
}