package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameInConfiguredProject(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: referencesForGlobals_1.ts
[|var [|{| "contextRangeIndex": 0 |}globalName|] = 0;|]
// @Filename: referencesForGlobals_2.ts
var y = [|globalName|];
// @Filename: tsconfig.json
{ "files": ["referencesForGlobals_1.ts", "referencesForGlobals_2.ts"] }`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, ToAny(f.Ranges()[1:])...)
}
