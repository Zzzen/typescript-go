package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestMemberCompletionOnTypeParameters2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class A {
    foo(): string { return ''; }
}

class B extends A {
    bar(): string {
        return '';
    }
}

class C<U extends A, T extends A> {
    x: U;
    y = this.x./**/ // completion list here
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.VerifyCompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &lsproto.CompletionItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
		},
		Items: &fourslash.VerifyCompletionsExpectedItems{
			Exact: []fourslash.ExpectedCompletionItem{"foo"},
		},
	})
}
