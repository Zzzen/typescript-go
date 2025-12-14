package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestAutoImportPaths(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /package1/jsconfig.json
{
  "compilerOptions": {
    checkJs: true,
    "paths": {
      "package1/*": ["./*"],
      "package2/*": ["../package2/*"]
    },
    "baseUrl": "."
  },
  "include": [
    ".",
    "../package2"
  ]
}
// @Filename: /package1/file1.js
bar/**/
// @Filename: /package2/file1.js
export const bar = 0;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixModuleSpecifiers(t, "", []string{"package2/file1"}, &lsutil.UserPreferences{ImportModuleSpecifierPreference: "shortest"})
}
