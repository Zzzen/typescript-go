package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoForJSDocWithHttpLinks(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @checkJs: true
// @filename: quickInfoForJSDocWithHttpLinks.js
/** @typedef {number} /*1*/https://wat */

/**
* @typedef {Object} Oops
* @property {number} /*2*/https://wass
*/


/** @callback /*3*/http://vad */

/** @see https://hvad */
var /*4*/see1 = true

/** @see {@link https://hva} */
var /*5*/see2 = true

/** {@link https://hvaD} */
var /*6*/see3 = true`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
