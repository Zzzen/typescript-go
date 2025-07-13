package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllReferencesOfJsonModule(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @resolveJsonModule: true
// @module: commonjs
// @esModuleInterop: true
// @Filename: /foo.ts
/*1*/import /*2*/settings from "./settings.json";
/*3*/settings;
// @Filename: /settings.json
 {}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3")
}
