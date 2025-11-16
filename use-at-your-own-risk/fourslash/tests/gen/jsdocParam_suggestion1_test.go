package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsdocParam_suggestion1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
/**
 * @param options - whatever
 * @param options.zone - equally bad
 */
declare function bad(options: any): void

/**
 * @param {number} obtuse
 */
function worse(): void {
    arguments
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "a.ts")
	f.VerifySuggestionDiagnostics(t, nil)
}
