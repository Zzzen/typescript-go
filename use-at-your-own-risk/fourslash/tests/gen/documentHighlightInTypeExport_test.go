package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestDocumentHighlightInTypeExport(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /1.ts
type [|A|] = 1;
export { [|A|] as [|B|] };
// @Filename: /2.ts
type [|A|] = 1;
let [|A|]: [|A|] = 1;
export { [|A|] as [|B|] };
// @Filename: /3.ts
type [|A|] = 1;
let [|A|]: [|A|] = 1;
export type { [|A|] as [|B|] };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
