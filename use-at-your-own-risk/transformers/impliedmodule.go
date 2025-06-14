package transformers

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/binder"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/printer"
)

type ImpliedModuleTransformer struct {
	Transformer
	compilerOptions           *core.CompilerOptions
	resolver                  binder.ReferenceResolver
	getEmitModuleFormatOfFile func(file ast.HasFileName) core.ModuleKind
	cjsTransformer            *Transformer
	esmTransformer            *Transformer
}

func NewImpliedModuleTransformer(emitContext *printer.EmitContext, compilerOptions *core.CompilerOptions, resolver binder.ReferenceResolver, getEmitModuleFormatOfFile func(file ast.HasFileName) core.ModuleKind) *Transformer {
	if resolver == nil {
		resolver = binder.NewReferenceResolver(compilerOptions, binder.ReferenceResolverHooks{})
	}
	tx := &ImpliedModuleTransformer{compilerOptions: compilerOptions, resolver: resolver, getEmitModuleFormatOfFile: getEmitModuleFormatOfFile}
	return tx.NewTransformer(tx.visit, emitContext)
}

func (tx *ImpliedModuleTransformer) visit(node *ast.Node) *ast.Node {
	switch node.Kind {
	case ast.KindSourceFile:
		node = tx.visitSourceFile(node.AsSourceFile())
	}
	return node
}

func (tx *ImpliedModuleTransformer) visitSourceFile(node *ast.SourceFile) *ast.Node {
	if node.IsDeclarationFile {
		return node.AsNode()
	}

	format := tx.getEmitModuleFormatOfFile(node)

	var transformer *Transformer
	if format >= core.ModuleKindES2015 {
		if tx.esmTransformer == nil {
			tx.esmTransformer = NewESModuleTransformer(tx.emitContext, tx.compilerOptions, tx.resolver, tx.getEmitModuleFormatOfFile)
		}
		transformer = tx.esmTransformer
	} else {
		if tx.cjsTransformer == nil {
			tx.cjsTransformer = NewCommonJSModuleTransformer(tx.emitContext, tx.compilerOptions, tx.resolver, tx.getEmitModuleFormatOfFile)
		}
		transformer = tx.cjsTransformer
	}

	return transformer.TransformSourceFile(node).AsNode()
}
