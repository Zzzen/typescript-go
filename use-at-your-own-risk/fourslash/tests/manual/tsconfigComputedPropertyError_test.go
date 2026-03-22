package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsconfigComputedPropertyError(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: tsconfig.json
{
    [|["oops!" + 42]|]: "true",
    "compilerOptions": { "lib": ["es5"] },
    "files": [
        "nonexistentfile.ts"
    ],
    "compileOnSave": true
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.MarkTestAsStradaServer()
	f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{
		{
			Message: "String literal with double quotes expected.",
			Code:    &lsproto.IntegerOrString{Integer: new(int32(1327))},
		},
	})
}
