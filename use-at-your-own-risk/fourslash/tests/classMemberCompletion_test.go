package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestClassMemberCompletionKeepsNameFallback(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class B {
	constructor(public value: string) {}
}
class C extends B {
	/*a*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "value",
					InsertText: new("value"),
				},
			},
		},
		UserPreferences: &lsutil.UserPreferences{IncludeCompletionsWithClassMemberSnippets: core.TSTrue},
	})
}

func TestImplementClassFixDoesNotAddInvalidOverride(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noImplicitOverride: true
class B {
    method() {}
}
class C implements B {[| |]}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCodeFix(t, fourslash.VerifyCodeFixOptions{
		Description: "Implement interface 'B'",
		NewFileContent: `class B {
    method() {}
}
class C implements B {
    method(): void {
        throw new Error("Method not implemented.");
    }
}`,
		Index: 0,
	})
}
