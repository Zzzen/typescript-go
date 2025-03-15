package transformers

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/printer"
)

type Transformer struct {
	emitContext *printer.EmitContext
	factory     *ast.NodeFactory
	visitor     *ast.NodeVisitor
}

func (tx *Transformer) newTransformer(visit func(node *ast.Node) *ast.Node, emitContext *printer.EmitContext) *Transformer {
	if tx.emitContext != nil {
		panic("Transformer already initialized")
	}
	if emitContext == nil {
		emitContext = printer.NewEmitContext()
	}
	tx.emitContext = emitContext
	tx.factory = emitContext.Factory
	tx.visitor = emitContext.NewNodeVisitor(visit)
	return tx
}

func (tx *Transformer) TransformSourceFile(file *ast.SourceFile) *ast.SourceFile {
	return tx.visitor.VisitSourceFile(file)
}
