package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQualifiedName_import_declaration_with_variable_entity_names(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module Alpha {
    export var [|{| "name" : "def" |}x|] = 100;
}

module Beta {
    import p = Alpha.[|{| "name" : "import" |}x|];
}

var x = Alpha.[|{| "name" : "mem" |}x|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "import")
	f.VerifyCompletions(t, nil, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "x",
					Detail: PtrTo("var Alpha.x: number"),
				},
			},
		},
	})
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, "import")
	f.VerifyBaselineGoToDefinition(t, "import")
}
