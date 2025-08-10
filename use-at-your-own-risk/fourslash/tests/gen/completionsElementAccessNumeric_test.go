package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionsElementAccessNumeric(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @target: esnext
type Tup = [
    /**
     * The first label
     */
    lbl1: number,
    /**
     * The second label
     */
    lbl2: number
];
declare var x: Tup;
x[|./**/|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "0",
					InsertText: PtrTo("[0]"),
					Documentation: &lsproto.StringOrMarkupContent{
						MarkupContent: &lsproto.MarkupContent{
							Kind:  lsproto.MarkupKindMarkdown,
							Value: "The first label",
						},
					},
					Detail: PtrTo("(property) 0: number (lbl1)"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "0",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
				&lsproto.CompletionItem{
					Label:      "1",
					InsertText: PtrTo("[1]"),
					Documentation: &lsproto.StringOrMarkupContent{
						MarkupContent: &lsproto.MarkupContent{
							Kind:  lsproto.MarkupKindMarkdown,
							Value: "The second label",
						},
					},
					Detail: PtrTo("(property) 1: number (lbl2)"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "1",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
}
