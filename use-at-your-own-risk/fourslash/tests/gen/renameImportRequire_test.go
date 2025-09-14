package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameImportRequire(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
[|import [|{| "contextRangeIndex": 0 |}e|] = require("mod4");|]
[|e|];
a = { [|e|] };
[|export { [|{| "contextRangeIndex": 4 |}e|] };|]
// @Filename: /b.ts
[|import { [|{| "contextRangeIndex": 6 |}e|] } from "./a";|]
[|export { [|{| "contextRangeIndex": 8 |}e|] };|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1], f.Ranges()[2], f.Ranges()[3], f.Ranges()[5], f.Ranges()[7], f.Ranges()[9])
}
