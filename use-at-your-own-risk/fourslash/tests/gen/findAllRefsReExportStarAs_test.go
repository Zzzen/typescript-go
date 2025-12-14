package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsReExportStarAs(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /leafModule.ts
export const /*helloDef*/hello = () => 'Hello';
// @Filename: /exporting.ts
export * as /*leafDef*/Leaf from './leafModule';
// @Filename: /importing.ts
 import { /*leafImportDef*/Leaf } from './exporting';
 /*leafUse*/[|Leaf|]./*helloUse*/[|hello|]()`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyNoErrors(t)
	f.VerifyBaselineFindAllReferences(t, "helloDef", "helloUse", "leafDef", "leafImportDef", "leafUse")
}
