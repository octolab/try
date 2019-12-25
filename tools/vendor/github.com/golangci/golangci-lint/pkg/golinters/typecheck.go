package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"golang.org/x/tools/go/analysis"
)

func NewTypecheck() *goanalysis.Linter {
	const linterName = "typecheck"
	analyzer := &analysis.Analyzer{
		Name: linterName,
		Doc:  goanalysis.TheOnlyanalyzerDoc,
		Run: func(pass *analysis.Pass) (interface{}, error) {
			return nil, nil
		},
	}
	linter := goanalysis.NewLinter(
		linterName,
		"Like the front-end of a Go compiler, parses and type-checks Go code",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
	linter.SetTypecheckMode()
	return linter
}
