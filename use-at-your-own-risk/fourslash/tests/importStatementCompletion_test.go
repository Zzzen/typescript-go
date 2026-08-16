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

func TestImportStatementCompletionUsesNamedImport(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
export interface I {}
// @Filename: 1.ts
import * as u from "./a";
[|import I/*a*/|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	preferences := lsutil.NewDefaultUserPreferences()
	preferences.IncludeCompletionsForModuleExports = core.TSFalse
	preferences.IncludeCompletionsForImportStatements = core.TSTrue
	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "I",
					InsertText: new(`import { I } from "./a";`),
					Data: &lsproto.CompletionItemData{
						AutoImport: &lsproto.AutoImportFix{ModuleSpecifier: "./a"},
					},
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: `import { I } from "./a";`,
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
		UserPreferences: &preferences,
	})
}
