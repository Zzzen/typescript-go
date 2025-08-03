package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestMemberListOnExplicitThis(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Restricted {
   n: number;
}
class C1 implements Restricted {
   n: number;
   m: number;
   f(this: this) {this./*1*/} // test on 'this.'
   g(this: Restricted) {this./*2*/}
}
function f(this: void) {this./*3*/}
function g(this: Restricted) {this./*4*/}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "f",
					Detail: PtrTo("(method) C1.f(this: this): void"),
				},
				&lsproto.CompletionItem{
					Label:  "g",
					Detail: PtrTo("(method) C1.g(this: Restricted): void"),
				},
				&lsproto.CompletionItem{
					Label:  "m",
					Detail: PtrTo("(property) C1.m: number"),
				},
				&lsproto.CompletionItem{
					Label:  "n",
					Detail: PtrTo("(property) C1.n: number"),
				},
			},
		},
	})
	f.VerifyCompletions(t, []string{"2", "4"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "n",
					Detail: PtrTo("(property) Restricted.n: number"),
				},
			},
		},
	})
	f.VerifyCompletions(t, "3", nil)
}
