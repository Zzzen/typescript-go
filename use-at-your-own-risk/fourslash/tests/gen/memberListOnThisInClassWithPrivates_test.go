package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestMemberListOnThisInClassWithPrivates(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C1 {
   public pubMeth() {this./**/} // test on 'this.'
   private privMeth() {}
   public pubProp = 0;
   private privProp = 0;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "privMeth",
					Detail: PtrTo("(method) C1.privMeth(): void"),
				},
				&lsproto.CompletionItem{
					Label:  "privProp",
					Detail: PtrTo("(property) C1.privProp: number"),
				},
				&lsproto.CompletionItem{
					Label:  "pubMeth",
					Detail: PtrTo("(method) C1.pubMeth(): void"),
				},
				&lsproto.CompletionItem{
					Label:  "pubProp",
					Detail: PtrTo("(property) C1.pubProp: number"),
				},
			},
		},
	})
}
