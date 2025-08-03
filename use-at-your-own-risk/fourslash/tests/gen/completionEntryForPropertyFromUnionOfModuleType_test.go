package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionEntryForPropertyFromUnionOfModuleType(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module E {
    export var n = 1;
    export var x = 0;
}
module F {
    export var n = 1;
    export var y = 0;
}
var q: typeof E | typeof F;
var j = q./*1*/`
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
					Label:  "n",
					Detail: PtrTo("(property) n: number"),
				},
			},
		},
	})
}
