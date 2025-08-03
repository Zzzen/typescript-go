package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoOnObjectLiteralWithOnlyGetter(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function /*1*/makePoint(x: number) {
    return {
        get x() { return x; },
    };
};
var /*4*/point = makePoint(2);
var /*2*/x = point./*3*/x;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "function makePoint(x: number): {\n    readonly x: number;\n}", "")
	f.VerifyQuickInfoAt(t, "2", "var x: number", "")
	f.VerifyQuickInfoAt(t, "4", "var point: {\n    readonly x: number;\n}", "")
	f.VerifyCompletions(t, "3", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "x",
					Detail: PtrTo("(property) x: number"),
				},
			},
		},
	})
}
