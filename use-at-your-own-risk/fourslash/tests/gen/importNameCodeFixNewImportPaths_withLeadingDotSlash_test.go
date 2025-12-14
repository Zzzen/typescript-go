package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestImportNameCodeFixNewImportPaths_withLeadingDotSlash(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
[|foo|]
// @Filename: /thisHasPathMapping.ts
export function foo() {};
// @Filename: /tsconfig.json
{
    "compilerOptions": {
        "baseUrl": ".",
        "paths": {
            "foo": ["././thisHasPathMapping"]
        }
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixAtPosition(t, []string{
		`import { foo } from "foo";

foo`,
	}, nil /*preferences*/)
}
