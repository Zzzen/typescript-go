package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	util "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestAutoImportCompletionsForArbitraryNonIdentifierExports(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
// @module: esnext
// @Filename: /a.ts
const foo = 0;
export { foo as "foo-bar" };
export const fooBar = 1;

// @Filename: /b.ts
foo/**/
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		Items: &fourslash.CompletionsExpectedItems{
			Excludes: []string{"foo-bar"},
			Includes: []fourslash.CompletionsExpectedItem{"fooBar"},
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &util.DefaultCommitCharacters,
			EditRange:        util.Ignored,
		},
	})
}
