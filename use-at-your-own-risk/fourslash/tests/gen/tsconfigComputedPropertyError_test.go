package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsconfigComputedPropertyError(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: tsconfig.json
{
    ["oops!" + 42]: "true",
    "files": [
        "nonexistentfile.ts"
    ],
    "compileOnSave": true
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.MarkTestAsStradaServer()
	f.VerifyNonSuggestionDiagnostics(t, nil)
}
