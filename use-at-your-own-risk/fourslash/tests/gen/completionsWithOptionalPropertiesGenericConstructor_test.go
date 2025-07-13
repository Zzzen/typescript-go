package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionsWithOptionalPropertiesGenericConstructor(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
interface Options {
    someFunction?: () => string
    anotherFunction?: () => string
}

export class Clazz<T extends Options> {
    constructor(public a: T) {}
}

new Clazz({ /*1*/ })`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &defaultCommitCharacters,
			EditRange:        ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "someFunction?",
					InsertText: ptrTo("someFunction"),
					FilterText: ptrTo("someFunction"),
					SortText:   ptrTo(string(ls.SortTextOptionalMember)),
				},
				&lsproto.CompletionItem{
					Label:      "anotherFunction?",
					InsertText: ptrTo("anotherFunction"),
					FilterText: ptrTo("anotherFunction"),
					SortText:   ptrTo(string(ls.SortTextOptionalMember)),
				},
			},
		},
	})
}
