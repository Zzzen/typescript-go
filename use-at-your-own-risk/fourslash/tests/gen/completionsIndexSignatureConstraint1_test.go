package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionsIndexSignatureConstraint1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true

repro #9900

interface Test {
  a?: number;
  b?: string;
}

interface TestIndex {
  [key: string]: Test;
}

declare function testFunc<T extends TestIndex>(t: T): void;

testFunc({
  test: {
    /**/
  },
});`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.VerifyCompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &lsproto.CompletionItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
		},
		Items: &fourslash.VerifyCompletionsExpectedItems{
			Exact: []fourslash.ExpectedCompletionItem{&lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "a"}, &lsproto.CompletionItem{SortText: ptrTo(string(ls.SortTextOptionalMember)), Label: "b"}},
		},
	})
}
