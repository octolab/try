package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"honnef.co/go/tools/simple"
)

func NewGosimple() *goanalysis.Linter {
	analyzers := analyzersMapToSlice(simple.Analyzers)
	setAnalyzersGoVersion(analyzers)

	return goanalysis.NewLinter(
		"gosimple",
		"Linter for Go source code that specializes in simplifying a code",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
