package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionsImport_preferUpdatingExistingImport(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @Filename: /deep/module/why/you/want/this/path.ts
export const x = 0;
export const y = 1;
// @Filename: /nice/reexport.ts
import { x, y } from "../deep/module/why/you/want/this/path";
export { x, y };
// @Filename: /index.ts
import { x } from "./deep/module/why/you/want/this/path";

y/**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: CompletionGlobalsPlus(
				[]fourslash.CompletionsExpectedItem{
					"x",
					&lsproto.CompletionItem{
						Label: "y",
						Data: PtrTo(any(&ls.CompletionItemData{
							AutoImport: &ls.AutoImportData{
								ModuleSpecifier: "./deep/module/why/you/want/this/path",
							},
						})),
						SortText:            PtrTo(string(ls.SortTextAutoImportSuggestions)),
						AdditionalTextEdits: fourslash.AnyTextEdits,
					},
				}, false),
		},
	})
}
