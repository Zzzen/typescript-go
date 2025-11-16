package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetJavaScriptSyntacticDiagnostics21(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @experimentalDecorators: true
// @Filename: a.js
@internal class C {}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyNonSuggestionDiagnostics(t, nil)
}
