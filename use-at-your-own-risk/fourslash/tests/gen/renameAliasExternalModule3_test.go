package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameAliasExternalModule3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
module SomeModule { [|export class [|{| "contextRangeIndex": 0 |}SomeClass|] { }|] }
export = SomeModule;
// @Filename: b.ts
import M = require("./a");
import C = M.[|SomeClass|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "SomeClass")
}
