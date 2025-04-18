package ls

import (
	"slices"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
)

func (l *LanguageService) GetDocumentDiagnostics(fileName string) []*ast.Diagnostic {
	program, file := l.getProgramAndFile(fileName)
	syntaxDiagnostics := program.GetSyntacticDiagnostics(file)
	semanticDiagnostics := program.GetSemanticDiagnostics(file)
	return slices.Concat(syntaxDiagnostics, semanticDiagnostics)
}
