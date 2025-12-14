package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestImportNameCodeFix_pathsWithExtension(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /tsconfig.json
{
  "compilerOptions": {
    "target": "ESNext",
    "module": "Node16",
    "moduleResolution": "Node16",
    "rootDir": "./src",
    "outDir": "./dist",
    "paths": {
      "#internals/*": ["./src/internals/*.ts"]
    }
  },
  "include": ["src"]
}
// @Filename: /src/internals/example.ts
export function helloWorld() {}
// @Filename: /src/index.ts
helloWorld/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixModuleSpecifiers(t, "", []string{"#internals/example"}, &lsutil.UserPreferences{ImportModuleSpecifierEnding: "js"})
}
