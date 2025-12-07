package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestImportNameCodeFix_pathsWithoutBaseUrl1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: tsconfig.json
{
  "compilerOptions": {
    "module": "commonjs",
    "paths": {
      "@app/*": ["./lib/*"]
    }
  }
}
// @Filename: index.ts
utils/**/
// @Filename: lib/utils.ts
export const utils = {};`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.VerifyImportFixAtPosition(t, []string{
		`import { utils } from "@app/utils";

utils`,
	}, nil /*preferences*/)
}
