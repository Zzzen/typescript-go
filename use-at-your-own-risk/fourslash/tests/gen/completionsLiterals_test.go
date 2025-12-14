package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionsLiterals(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const x: 0 | "one" = /**/;
const y: 0 | "one" | 1n = /*1*/;
const y2: 0 | "one" | 1n = 'one'/*2*/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "0",
					Kind:   PtrTo(lsproto.CompletionItemKindConstant),
					Detail: PtrTo("0"),
				},
				&lsproto.CompletionItem{
					Label:  "\"one\"",
					Kind:   PtrTo(lsproto.CompletionItemKindConstant),
					Detail: PtrTo("\"one\""),
				},
			},
		},
	})
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "0",
					Kind:   PtrTo(lsproto.CompletionItemKindConstant),
					Detail: PtrTo("0"),
				},
				&lsproto.CompletionItem{
					Label:  "\"one\"",
					Kind:   PtrTo(lsproto.CompletionItemKindConstant),
					Detail: PtrTo("\"one\""),
				},
				&lsproto.CompletionItem{
					Label:  "1n",
					Kind:   PtrTo(lsproto.CompletionItemKindConstant),
					Detail: PtrTo("1n"),
				},
			},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Excludes: []string{
				"\"one\"",
			},
		},
	})
}
