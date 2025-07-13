package estransforms

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/printer"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/transformers"
)

type nullishCoalescingTransformer struct {
	transformers.Transformer
}

func (ch *nullishCoalescingTransformer) visit(node *ast.Node) *ast.Node {
	if node.SubtreeFacts()&ast.SubtreeContainsNullishCoalescing == 0 {
		return node
	}
	switch node.Kind {
	case ast.KindBinaryExpression:
		return ch.visitBinaryExpression(node.AsBinaryExpression())
	default:
		return ch.Visitor().VisitEachChild(node)
	}
}

func (ch *nullishCoalescingTransformer) visitBinaryExpression(node *ast.BinaryExpression) *ast.Node {
	switch node.OperatorToken.Kind {
	case ast.KindQuestionQuestionToken:
		left := ch.Visitor().VisitNode(node.Left)
		right := left
		if !transformers.IsSimpleCopiableExpression(left) {
			right = ch.Factory().NewTempVariable()
			ch.EmitContext().AddVariableDeclaration(right)
			left = ch.Factory().NewAssignmentExpression(right, left)
		}
		return ch.Factory().NewConditionalExpression(
			createNotNullCondition(ch.EmitContext(), left, right, false),
			ch.Factory().NewToken(ast.KindQuestionToken),
			right,
			ch.Factory().NewToken(ast.KindColonToken),
			ch.Visitor().VisitNode(node.Right),
		)
	default:
		return ch.Visitor().VisitEachChild(node.AsNode())
	}
}

func newNullishCoalescingTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &nullishCoalescingTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
