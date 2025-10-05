package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestReferencesIsAvailableThroughGlobalNoCrash(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /packages/playwright-core/bundles/utils/node_modules/@types/debug/index.d.ts
declare var debug: debug.Debug & { debug: debug.Debug; default: debug.Debug };
export = debug;
export as namespace debug;
declare namespace debug {
    interface Debug {
       coerce: (val: any) => any;
    }
}
// @Filename: /packages/playwright-core/bundles/utils/node_modules/@types/debug/package.json
{ "types": "index.d.ts" }
// @Filename: /packages/playwright-core/src/index.ts
export const debug: typeof import('../bundles/utils/node_modules//*1*/@types/debug') = require('./utilsBundleImpl').debug;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1")
}
