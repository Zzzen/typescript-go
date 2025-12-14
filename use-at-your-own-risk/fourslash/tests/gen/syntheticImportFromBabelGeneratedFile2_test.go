package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestSyntheticImportFromBabelGeneratedFile2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @allowSyntheticDefaultImports: true
// @Filename: /a.js
Object.defineProperty(exports, "__esModule", {
    value: true
});
exports.default = f;
/**
 * Run this function
 * @param {string} t
 */
function f(t) {}
// @Filename: /b.js
import f from "./a"
/**/f`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "", "(alias) function f(t: string): void\nimport f", "Run this function")
}
