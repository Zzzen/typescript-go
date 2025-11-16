package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestImportNameCodeFix_preferBaseUrl(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /tsconfig.json
{ "compilerOptions": { "baseUrl": "./src" } }
// @Filename: /src/d0/d1/d2/file.ts
foo/**/;
// @Filename: /src/d0/a.ts
export const foo = 0;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/src/d0/d1/d2/file.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import { foo } from "d0/a";

foo;`,
	}, nil /*preferences*/)
}
