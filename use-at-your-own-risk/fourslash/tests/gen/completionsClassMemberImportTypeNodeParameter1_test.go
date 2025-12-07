package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionsClassMemberImportTypeNodeParameter1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: node18
// @Filename: /generation.d.ts
export type GenerationConfigType = { max_length?: number };
// @FileName: /index.d.ts
export declare class PreTrainedModel {
  _get_generation_config(
    param: import("./generation.js").GenerationConfigType,
  ): import("./generation.js").GenerationConfigType;
}

export declare class BlenderbotSmallPreTrainedModel extends PreTrainedModel {
  /*1*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:               "_get_generation_config",
					InsertText:          PtrTo("_get_generation_config(param: import(\"./generation.js\").GenerationConfigType): import(\"./generation.js\").GenerationConfigType;"),
					FilterText:          PtrTo("_get_generation_config"),
					AdditionalTextEdits: fourslash.AnyTextEdits,
				},
			},
		},
	})
}
