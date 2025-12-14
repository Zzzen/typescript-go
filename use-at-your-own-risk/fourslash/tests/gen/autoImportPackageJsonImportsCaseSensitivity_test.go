package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestAutoImportPackageJsonImportsCaseSensitivity(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: node18
// @allowImportingTsExtensions: true
// @Filename: /package.json
{
  "type": "module",
  "imports": {
    "#src/*": "./SRC/*"
  }
}
// @Filename: /src/add.ts
export function add(a: number, b: number) {}
// @Filename: /src/index.ts
add/*imports*/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixModuleSpecifiers(t, "imports", []string{"#src/add.ts"}, &lsutil.UserPreferences{ImportModuleSpecifierPreference: "non-relative"})
}
