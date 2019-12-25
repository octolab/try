package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"honnef.co/go/tools/stylecheck"
)

func NewStylecheck() *goanalysis.Linter {
	analyzers := analyzersMapToSlice(stylecheck.Analyzers)
	setAnalyzersGoVersion(analyzers)

	return goanalysis.NewLinter(
		"stylecheck",
		"Stylecheck is a replacement for golint",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
