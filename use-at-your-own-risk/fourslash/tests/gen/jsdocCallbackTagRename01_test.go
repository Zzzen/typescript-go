package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsdocCallbackTagRename01(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowNonTsExtensions: true
// @Filename: jsDocCallback.js

/**
 * [|@callback [|{| "contextRangeIndex": 0 |}FooCallback|]
 * @param {string} eventName - Rename should work
 |]*/

/** @type {/*1*/[|FooCallback|]} */
var t;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1])
}
