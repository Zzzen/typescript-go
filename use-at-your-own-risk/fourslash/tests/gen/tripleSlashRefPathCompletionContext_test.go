package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTripleSlashRefPathCompletionContext(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: f.ts
/*f*/
// @Filename: test.ts
/// <reference path/*0*/=/*1*/"/*8*/
/// <reference path/*2*/=/*3*/"/*9*/"/*4*/ /*5*///*6*/>/*7*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"0", "1", "2", "3", "4", "5", "6", "7"}, nil)
	f.VerifyCompletions(t, []string{"8", "9"}, &fourslash.VerifyCompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &lsproto.CompletionItemDefaults{
			CommitCharacters: &[]string{},
		},
		Items: &fourslash.VerifyCompletionsExpectedItems{
			Exact: []fourslash.ExpectedCompletionItem{"f.ts"},
		},
	})
}
