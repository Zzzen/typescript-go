package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsForDefaultExport(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
export default function /*def*/f() {}
// @Filename: b.ts
import /*deg*/g from "./a";
[|/*ref*/g|]();
// @Filename: c.ts
import { f } from "./a";`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "def", "deg")
	f.VerifyBaselineGoToDefinition(t, true, "ref")
}
