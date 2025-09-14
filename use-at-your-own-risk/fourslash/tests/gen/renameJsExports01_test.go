package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameJsExports01(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: a.js
[|exports.[|{| "contextRangeIndex": 0 |}area|] = function (r) { return r * r; }|]
// @Filename: b.js
var mod = require('./a');
var t = mod./*1*/[|area|](10);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1")
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "area")
}
