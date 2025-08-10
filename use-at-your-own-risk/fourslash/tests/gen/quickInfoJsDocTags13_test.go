package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoJsDocTags13(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @checkJs: true
// @filename: ./a.js
/**
 * First overload
 * @overload
 * @param {number} a
 * @returns {void}
 */

/**
 * Second overload
 * @overload
 * @param {string} a
 * @returns {void}
 */

/**
 * @param {string | number} a
 * @returns {void}
 */
function f(a) {}

f(/*a*/1);
f(/*b*/"");`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineSignatureHelp(t)
}
