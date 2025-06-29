package estransforms

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/printer"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/transformers"
)

type logicalAssignmentTransformer struct {
	transformers.Transformer
}

func (ch *logicalAssignmentTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newLogicalAssignmentTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &logicalAssignmentTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
