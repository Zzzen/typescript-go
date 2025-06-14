package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestBasicMultifileCompletions(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export const foo = { bar: 'baz' };

// @Filename: /b.ts
import { foo } from './a';
const test = foo./*1*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.VerifyCompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &lsproto.CompletionItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
		},
		Items: &fourslash.VerifyCompletionsExpectedItems{
			Includes: []fourslash.ExpectedCompletionItem{
				&lsproto.CompletionItem{
					Label:      "bar",
					Kind:       ptrTo(lsproto.CompletionItemKindField),
					SortText:   ptrTo(string(ls.SortTextLocationPriority)),
					FilterText: ptrTo(".bar"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						InsertReplaceEdit: &lsproto.InsertReplaceEdit{
							NewText: "bar",
							Insert: lsproto.Range{
								Start: lsproto.Position{Line: 1, Character: 17},
								End:   lsproto.Position{Line: 1, Character: 17},
							},
							Replace: lsproto.Range{
								Start: lsproto.Position{Line: 1, Character: 17},
								End:   lsproto.Position{Line: 1, Character: 17},
							},
						},
					},
				},
			},
		},
	})
}
