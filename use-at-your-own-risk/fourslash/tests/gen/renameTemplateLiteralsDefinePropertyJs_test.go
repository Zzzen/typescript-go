package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameTemplateLiteralsDefinePropertyJs(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: a.js
let obj = {};

Object.defineProperty(obj, ` + "`" + `[|prop|]` + "`" + `, { value: 0 });

obj = {
    [|[` + "`" + `[|{| "contextRangeIndex": 1 |}prop|]` + "`" + `]: 1|]
};

obj.[|prop|];
obj['[|prop|]'];
obj["[|prop|]"];
obj[` + "`" + `[|prop|]` + "`" + `];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "prop")
}
