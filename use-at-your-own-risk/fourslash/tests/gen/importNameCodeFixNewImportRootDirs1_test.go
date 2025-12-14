package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestImportNameCodeFixNewImportRootDirs1(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a/f1.ts
[|foo/*0*/();|]
// @Filename: a/b/index.ts
export function foo() {};
// @Filename: tsconfig.json
{
    "compilerOptions": {
        "rootDirs": [
            "a"
        ]
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixAtPosition(t, []string{
		`import { foo } from "./b";

foo();`,
	}, nil /*preferences*/)
}
