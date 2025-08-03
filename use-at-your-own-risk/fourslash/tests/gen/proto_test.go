package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestProto(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module M {
    export interface /*1*/__proto__ {}
}
var /*2*/__proto__: M.__proto__;
/*3*/
var /*4*/fun: (__proto__: any) => boolean;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "interface M.__proto__", "")
	f.VerifyQuickInfoAt(t, "2", "var __proto__: M.__proto__", "")
	f.VerifyCompletions(t, "3", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "__proto__",
					Detail: PtrTo("var __proto__: M.__proto__"),
				},
			},
		},
	})
	f.Insert(t, "__proto__")
	f.VerifyBaselineGoToDefinition(t)
	f.VerifyQuickInfoAt(t, "4", "var fun: (__proto__: any) => boolean", "")
}
