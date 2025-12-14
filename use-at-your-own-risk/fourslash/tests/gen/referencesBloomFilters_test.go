package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestReferencesBloomFilters(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: declaration.ts
var container = { /*1*/searchProp : 1 };
// @Filename: expression.ts
function blah() { return (1 + 2 + container.searchProp()) === 2;  };
// @Filename: stringIndexer.ts
function blah2() { container["searchProp"] };
// @Filename: redeclaration.ts
container = { "searchProp" : 18 };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t, "1")
}
