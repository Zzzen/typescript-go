package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoJsDocTagsCallback(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noEmit: true
// @allowJs: true
// @Filename: quickInfoJsDocTagsCallback.js
/**
 * @callback cb/*1*/
 * @param {string} x - x comment
 */

/**
 * @param {/*2*/cb} bar -callback comment
 */
function foo(bar) {
    bar(bar);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
