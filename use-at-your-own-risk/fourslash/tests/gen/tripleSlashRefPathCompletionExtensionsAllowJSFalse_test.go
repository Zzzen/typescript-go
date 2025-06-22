package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTripleSlashRefPathCompletionExtensionsAllowJSFalse(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: test0.ts
/// <reference path="/*0*/
/// <reference path=".//*1*/
/// <reference path="./f/*2*/
// @Filename: f1.ts

// @Filename: f1.js

// @Filename: f1.d.ts

// @Filename: f1.tsx

// @Filename: f1.js

// @Filename: f1.jsx

// @Filename: f1.cs
`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, f.Markers(), &fourslash.VerifyCompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &lsproto.CompletionItemDefaults{
			CommitCharacters: &[]string{},
		},
		Items: &fourslash.VerifyCompletionsExpectedItems{
			Exact: []fourslash.ExpectedCompletionItem{"f1.d.ts", "f1.ts", "f1.tsx"},
		},
	})
}
