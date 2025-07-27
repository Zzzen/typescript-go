package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionSourceUnit(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
 //MyFile Comments
 //more comments
 /// <reference path="so/*unknownFile*/mePath.ts" />
 /// <reference path="[|b/*knownFile*/.ts|]" />

 class clsInOverload {
     static fnOverload();
     static fnOverload(foo: string);
     static fnOverload(foo: any) { }
 }

// @Filename: b.ts
/*fileB*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "unknownFile", "knownFile")
}
