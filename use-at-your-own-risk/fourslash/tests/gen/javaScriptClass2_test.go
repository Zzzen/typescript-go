package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJavaScriptClass2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowNonTsExtensions: true
// @Filename: Foo.js
class Foo {
   constructor() {
       [|this.[|{| "contextRangeIndex": 0 |}union|] = 'foo';|]
       [|this.[|{| "contextRangeIndex": 2 |}union|] = 100;|]
   }
   method() { return this.[|union|]; }
}
var x = new Foo();
x.[|union|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "union")
}
