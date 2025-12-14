package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestImportNameCodeFixNewImportNodeModules0(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|f1/*0*/();|]
// @Filename: ../package.json
{ "dependencies": { "fake-module": "latest" } }
// @Filename: ../node_modules/fake-module/index.ts
export var v1 = 5;
export function f1();
// @Filename: ../node_modules/fake-module/package.json
{}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixAtPosition(t, []string{
		`import { f1 } from "fake-module";

f1();`,
	}, nil /*preferences*/)
}
