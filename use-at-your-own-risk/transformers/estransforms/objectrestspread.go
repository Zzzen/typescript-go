package estransforms

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/printer"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/transformers"
)

type objectRestSpreadTransformer struct {
	transformers.Transformer
}

func (ch *objectRestSpreadTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newObjectRestSpreadTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &objectRestSpreadTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
